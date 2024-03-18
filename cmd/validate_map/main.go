// Copyright (c) 2024 Kevin Damm
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// github:kevindamm/wits-go/cmd/validate_map/main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevindamm/wits-go/schema"
	"github.com/kevindamm/wits-go/state"
	"github.com/kevindamm/wits-go/witsjson"
)

func main() {
	// Expects a single argument, the path to the map to open and analyze.
	debug := flag.Bool("debug", false,
		"set this flag to see more detail printed to the console.")
	flag.Parse()

	filename := flag.Args()[0]
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// IsDir is short for fileInfo.Mode().IsDir()
	if fileInfo.IsDir() {
		if *debug {
			fmt.Printf("parsing maps in directory %s\n", filename)
		}
		for filename := range mapFiles(filename) {
			if *debug {
				fmt.Printf("parsing map %s\n", filename)
			}
			if _, err := readAndValidateGameMap(filename); err != nil {
				fmt.Printf("error loading map %s\n", filename)
				fmt.Println(err)
			}
		}
	} else {
		if *debug {
			fmt.Printf("parsing map %s\n", filename)
		}
		if _, err := readAndValidateGameMap(filename); err != nil {
			fmt.Printf("error loading map %s\n", filename)
			fmt.Println(err)
		}
	}
}

func readAndValidateGameMap(filename string) (state.GameMap, error) {
	filedata, err := os.ReadFile(filename)
	if err != nil {
		return state.GameMap{}, err
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error when reading map file %s\n", filename)
			fmt.Println(err)
		}
	}()

	var gamemap witsjson.GameMapJSON
	if err = json.Unmarshal(filedata, &gamemap); err != nil {
		fmt.Println(err)
		return state.GameMap{}, err
	}

	// Ensure no two tiles are located at the same coordinate.
	if err := checkExclusivity(gamemap.Terrain()); err != nil {
		fmt.Println(err)
		return state.GameMap{}, err
	}

	return state.NewGameMap(gamemap), nil
}

func checkExclusivity(terrain schema.TerrainDefinition) error {
	size := len(terrain.Floor()) + len(terrain.Wall()) + len(terrain.Bonus()) + len(terrain.Spawn()) + len(terrain.Base())
	positions := make(map[int16]schema.TileDefinition, size)
	coord := func(hx schema.HexCoord) int16 {
		// This is a bit of a cheat but we know that no coordinate value exceeds 12.  It'll fit with plenty of room.
		// THIS WOULDN'T SCALE TO ARBITRARY USER COORDINATES AS THE REST OF THE SYSTEM EASILY WOULD.  do it better if it comes to that.
		return int16((hx.I() << 8) + hx.J())
	}

	for _, floor := range terrain.Floor() {
		hex := coord(floor.Position())
		if positions[hex] != nil {
			fmt.Printf("Repeated floor position: [%d, %d]\n", floor.Position().I(), floor.Position().J())
			return fmt.Errorf("coordinate repeat [%d, %d]", floor.Position().I(), floor.Position().J())
		}
		positions[coord(floor.Position())] = floor
	}
	for _, wall := range terrain.Wall() {
		hex := coord(wall.Position())
		if positions[hex] != nil {
			fmt.Printf("Coordinate [%d, %d] repeated betewen %s, %s",
				wall.Position().I(), wall.Position().J(), positions[hex].Typename(), wall.Typename())
			return fmt.Errorf("coordinate repeat [%d, %d]", wall.Position().I(), wall.Position().J())
		}
		positions[coord(wall.Position())] = wall
	}
	for _, bonus := range terrain.Bonus() {
		hex := coord(bonus.Position())
		if positions[hex] != nil {
			fmt.Printf("Coordinate [%d, %d] repeated betewen %s, %s",
				bonus.Position().I(), bonus.Position().J(), positions[hex].Typename(), bonus.Typename())
			return fmt.Errorf("coordinate repeat [%d, %d]", bonus.Position().I(), bonus.Position().J())
		}
		positions[coord(bonus.Position())] = bonus
	}
	for _, spawn := range terrain.Spawn() {
		hex := coord(spawn.Position())
		if positions[hex] != nil {
			fmt.Printf("Coordinate [%d, %d] repeated betewen %s, %s",
				spawn.Position().I(), spawn.Position().J(), positions[hex].Typename(), spawn.Typename())
			return fmt.Errorf("coordinate repeat [%d, %d]", spawn.Position().I(), spawn.Position().J())
		}
		positions[coord(spawn.Position())] = spawn
	}
	for _, base := range terrain.Base() {
		hex := coord(base.Position())
		if positions[hex] != nil {
			fmt.Printf("Coordinate [%d, %d] repeated betewen %s, %s",
				base.Position().I(), base.Position().J(), positions[hex].Typename(), base.Typename())
			return fmt.Errorf("coordinate repeat [%d, %d]", base.Position().I(), base.Position().J())
		}
		// Also check six coordinates surrounding base.
		for _, neighbor := range listSurroundingPositions(base.Position().I(), base.Position().J()) {
			hex := coord(neighbor)
			if positions[hex] != nil {
				fmt.Printf("Coordinate [%d, %d] repeated betewen %s, %s",
					neighbor.I(), neighbor.J(), positions[hex].Typename(), base.Typename())
				return fmt.Errorf("coordinate collides with base [%d, %d]", neighbor.I(), neighbor.J())
			}
		}
		positions[coord(base.Position())] = base
	}

	return nil
}

func mapFiles(dirName string) <-chan string {
	fpathchan := make(chan string)
	go func() {
		defer close(fpathchan)
		err := filepath.WalkDir(dirName,
			func(fpath string, d fs.DirEntry, err error) error {
				if err != nil {
					// Forward errors.
					return err
				}
				if d.IsDir() {
					// Skip subdirectories.
					return nil
				}
				if strings.HasSuffix(fpath, ".json") {
					fpathchan <- fpath
				}
				return nil
			})
		if err != nil {
			fmt.Printf("error walking the path %q\n  %v\n", dirName, err)
			return
		}
	}()
	return fpathchan
}

func listSurroundingPositions(i, j int) []schema.HexCoord {
	surrounding := make([]schema.HexCoord, 0)

	neighborTable := [][]int{
		{-1, -1 + (i % 2)},
		{-1, 0 + (i % 2)},
		{0, -1},
		{0, 1},
		{1, -1 + (i % 2)},
		{1, 0 + (i % 2)},
	}

	for k := 0; k < 6; k += 1 {
		surrnext := schema.NewHexCoord(
			i+neighborTable[k][0],
			j+neighborTable[k][1])
		// Only positive coordinates are valid in the legacy coordinate system.
		if surrnext.I() < 0 || surrnext.J() < 0 {
			continue
		}
		surrounding = append(surrounding, surrnext)
	}
	return surrounding
}
