// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/general"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
		Type: "Detail",
		Fields: map[string]*sqlgraph.FieldSpec{
			detail.FieldCreatedAt:              {Type: field.TypeUint32, Column: detail.FieldCreatedAt},
			detail.FieldUpdatedAt:              {Type: field.TypeUint32, Column: detail.FieldUpdatedAt},
			detail.FieldDeletedAt:              {Type: field.TypeUint32, Column: detail.FieldDeletedAt},
			detail.FieldAppID:                  {Type: field.TypeUUID, Column: detail.FieldAppID},
			detail.FieldUserID:                 {Type: field.TypeUUID, Column: detail.FieldUserID},
			detail.FieldGoodID:                 {Type: field.TypeUUID, Column: detail.FieldGoodID},
			detail.FieldOrderID:                {Type: field.TypeUUID, Column: detail.FieldOrderID},
			detail.FieldPaymentID:              {Type: field.TypeUUID, Column: detail.FieldPaymentID},
			detail.FieldCoinTypeID:             {Type: field.TypeUUID, Column: detail.FieldCoinTypeID},
			detail.FieldPaymentCoinTypeID:      {Type: field.TypeUUID, Column: detail.FieldPaymentCoinTypeID},
			detail.FieldPaymentCoinUsdCurrency: {Type: field.TypeFloat64, Column: detail.FieldPaymentCoinUsdCurrency},
			detail.FieldUnits:                  {Type: field.TypeUint32, Column: detail.FieldUnits},
			detail.FieldAmount:                 {Type: field.TypeFloat64, Column: detail.FieldAmount},
			detail.FieldUsdAmount:              {Type: field.TypeFloat64, Column: detail.FieldUsdAmount},
			detail.FieldCommission:             {Type: field.TypeFloat64, Column: detail.FieldCommission},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   general.Table,
			Columns: general.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: general.FieldID,
			},
		},
		Type: "General",
		Fields: map[string]*sqlgraph.FieldSpec{
			general.FieldCreatedAt:       {Type: field.TypeUint32, Column: general.FieldCreatedAt},
			general.FieldUpdatedAt:       {Type: field.TypeUint32, Column: general.FieldUpdatedAt},
			general.FieldDeletedAt:       {Type: field.TypeUint32, Column: general.FieldDeletedAt},
			general.FieldAppID:           {Type: field.TypeUUID, Column: general.FieldAppID},
			general.FieldUserID:          {Type: field.TypeUUID, Column: general.FieldUserID},
			general.FieldGoodID:          {Type: field.TypeUUID, Column: general.FieldGoodID},
			general.FieldCoinTypeID:      {Type: field.TypeUUID, Column: general.FieldCoinTypeID},
			general.FieldTotalUnits:      {Type: field.TypeUint32, Column: general.FieldTotalUnits},
			general.FieldSelfUnits:       {Type: field.TypeUint32, Column: general.FieldSelfUnits},
			general.FieldTotalAmount:     {Type: field.TypeFloat64, Column: general.FieldTotalAmount},
			general.FieldSelfAmount:      {Type: field.TypeFloat64, Column: general.FieldSelfAmount},
			general.FieldTotalCommission: {Type: field.TypeFloat64, Column: general.FieldTotalCommission},
			general.FieldSelfCommission:  {Type: field.TypeFloat64, Column: general.FieldSelfCommission},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DetailQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DetailQuery builder.
func (dq *DetailQuery) Filter() *DetailFilter {
	return &DetailFilter{dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DetailMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DetailMutation builder.
func (m *DetailMutation) Filter() *DetailFilter {
	return &DetailFilter{m}
}

// DetailFilter provides a generic filtering capability at runtime for DetailQuery.
type DetailFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *DetailFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *DetailFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *DetailFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *DetailFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *DetailFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *DetailFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *DetailFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldUserID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *DetailFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldGoodID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *DetailFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldOrderID))
}

// WherePaymentID applies the entql [16]byte predicate on the payment_id field.
func (f *DetailFilter) WherePaymentID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldPaymentID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *DetailFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldCoinTypeID))
}

// WherePaymentCoinTypeID applies the entql [16]byte predicate on the payment_coin_type_id field.
func (f *DetailFilter) WherePaymentCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldPaymentCoinTypeID))
}

// WherePaymentCoinUsdCurrency applies the entql float64 predicate on the payment_coin_usd_currency field.
func (f *DetailFilter) WherePaymentCoinUsdCurrency(p entql.Float64P) {
	f.Where(p.Field(detail.FieldPaymentCoinUsdCurrency))
}

// WhereUnits applies the entql uint32 predicate on the units field.
func (f *DetailFilter) WhereUnits(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldUnits))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *DetailFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(detail.FieldAmount))
}

// WhereUsdAmount applies the entql float64 predicate on the usd_amount field.
func (f *DetailFilter) WhereUsdAmount(p entql.Float64P) {
	f.Where(p.Field(detail.FieldUsdAmount))
}

// WhereCommission applies the entql float64 predicate on the commission field.
func (f *DetailFilter) WhereCommission(p entql.Float64P) {
	f.Where(p.Field(detail.FieldCommission))
}

// addPredicate implements the predicateAdder interface.
func (gq *GeneralQuery) addPredicate(pred func(s *sql.Selector)) {
	gq.predicates = append(gq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GeneralQuery builder.
func (gq *GeneralQuery) Filter() *GeneralFilter {
	return &GeneralFilter{gq}
}

// addPredicate implements the predicateAdder interface.
func (m *GeneralMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GeneralMutation builder.
func (m *GeneralMutation) Filter() *GeneralFilter {
	return &GeneralFilter{m}
}

// GeneralFilter provides a generic filtering capability at runtime for GeneralQuery.
type GeneralFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *GeneralFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *GeneralFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(general.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *GeneralFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(general.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *GeneralFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(general.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *GeneralFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(general.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *GeneralFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(general.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *GeneralFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(general.FieldUserID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *GeneralFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(general.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *GeneralFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(general.FieldCoinTypeID))
}

// WhereTotalUnits applies the entql uint32 predicate on the total_units field.
func (f *GeneralFilter) WhereTotalUnits(p entql.Uint32P) {
	f.Where(p.Field(general.FieldTotalUnits))
}

// WhereSelfUnits applies the entql uint32 predicate on the self_units field.
func (f *GeneralFilter) WhereSelfUnits(p entql.Uint32P) {
	f.Where(p.Field(general.FieldSelfUnits))
}

// WhereTotalAmount applies the entql float64 predicate on the total_amount field.
func (f *GeneralFilter) WhereTotalAmount(p entql.Float64P) {
	f.Where(p.Field(general.FieldTotalAmount))
}

// WhereSelfAmount applies the entql float64 predicate on the self_amount field.
func (f *GeneralFilter) WhereSelfAmount(p entql.Float64P) {
	f.Where(p.Field(general.FieldSelfAmount))
}

// WhereTotalCommission applies the entql float64 predicate on the total_commission field.
func (f *GeneralFilter) WhereTotalCommission(p entql.Float64P) {
	f.Where(p.Field(general.FieldTotalCommission))
}

// WhereSelfCommission applies the entql float64 predicate on the self_commission field.
func (f *GeneralFilter) WhereSelfCommission(p entql.Float64P) {
	f.Where(p.Field(general.FieldSelfCommission))
}
