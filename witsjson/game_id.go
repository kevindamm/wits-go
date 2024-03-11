package witsjson

import "encoding/json"

// Game IDs read in by JSON Unmarshaling will have already been shortened.
type OsnGameID string

func (gameID OsnGameID) ShortID() string {
	return string(gameID)
}

func (gameID OsnGameID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(gameID))
}
