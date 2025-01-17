// Code generated by ent, DO NOT EDIT.

package player

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kevindamm/wits-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldID, id))
}

// Gcid applies equality check predicate on the "gcid" field. It's identical to GcidEQ.
func Gcid(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldGcid, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldName, v))
}

// GcidEQ applies the EQ predicate on the "gcid" field.
func GcidEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldGcid, v))
}

// GcidNEQ applies the NEQ predicate on the "gcid" field.
func GcidNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldGcid, v))
}

// GcidIn applies the In predicate on the "gcid" field.
func GcidIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldGcid, vs...))
}

// GcidNotIn applies the NotIn predicate on the "gcid" field.
func GcidNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldGcid, vs...))
}

// GcidGT applies the GT predicate on the "gcid" field.
func GcidGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldGcid, v))
}

// GcidGTE applies the GTE predicate on the "gcid" field.
func GcidGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldGcid, v))
}

// GcidLT applies the LT predicate on the "gcid" field.
func GcidLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldGcid, v))
}

// GcidLTE applies the LTE predicate on the "gcid" field.
func GcidLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldGcid, v))
}

// GcidContains applies the Contains predicate on the "gcid" field.
func GcidContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldGcid, v))
}

// GcidHasPrefix applies the HasPrefix predicate on the "gcid" field.
func GcidHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldGcid, v))
}

// GcidHasSuffix applies the HasSuffix predicate on the "gcid" field.
func GcidHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldGcid, v))
}

// GcidEqualFold applies the EqualFold predicate on the "gcid" field.
func GcidEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldGcid, v))
}

// GcidContainsFold applies the ContainsFold predicate on the "gcid" field.
func GcidContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldGcid, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldName, v))
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.PlayerRole) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newRolesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Player) predicate.Player {
	return predicate.Player(sql.NotPredicates(p))
}
