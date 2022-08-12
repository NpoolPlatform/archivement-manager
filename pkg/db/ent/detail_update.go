// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// DetailUpdate is the builder for updating Detail entities.
type DetailUpdate struct {
	config
	hooks    []Hook
	mutation *DetailMutation
}

// Where appends a list predicates to the DetailUpdate builder.
func (du *DetailUpdate) Where(ps ...predicate.Detail) *DetailUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DetailUpdate) SetCreatedAt(u uint32) *DetailUpdate {
	du.mutation.ResetCreatedAt()
	du.mutation.SetCreatedAt(u)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCreatedAt(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetCreatedAt(*u)
	}
	return du
}

// AddCreatedAt adds u to the "created_at" field.
func (du *DetailUpdate) AddCreatedAt(u int32) *DetailUpdate {
	du.mutation.AddCreatedAt(u)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DetailUpdate) SetUpdatedAt(u uint32) *DetailUpdate {
	du.mutation.ResetUpdatedAt()
	du.mutation.SetUpdatedAt(u)
	return du
}

// AddUpdatedAt adds u to the "updated_at" field.
func (du *DetailUpdate) AddUpdatedAt(u int32) *DetailUpdate {
	du.mutation.AddUpdatedAt(u)
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DetailUpdate) SetDeletedAt(u uint32) *DetailUpdate {
	du.mutation.ResetDeletedAt()
	du.mutation.SetDeletedAt(u)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DetailUpdate) SetNillableDeletedAt(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetDeletedAt(*u)
	}
	return du
}

// AddDeletedAt adds u to the "deleted_at" field.
func (du *DetailUpdate) AddDeletedAt(u int32) *DetailUpdate {
	du.mutation.AddDeletedAt(u)
	return du
}

// SetAppID sets the "app_id" field.
func (du *DetailUpdate) SetAppID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetAppID(u)
	return du
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableAppID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetAppID(*u)
	}
	return du
}

// ClearAppID clears the value of the "app_id" field.
func (du *DetailUpdate) ClearAppID() *DetailUpdate {
	du.mutation.ClearAppID()
	return du
}

// SetUserID sets the "user_id" field.
func (du *DetailUpdate) SetUserID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetUserID(u)
	return du
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableUserID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetUserID(*u)
	}
	return du
}

// ClearUserID clears the value of the "user_id" field.
func (du *DetailUpdate) ClearUserID() *DetailUpdate {
	du.mutation.ClearUserID()
	return du
}

// SetDirectContributorID sets the "direct_contributor_id" field.
func (du *DetailUpdate) SetDirectContributorID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetDirectContributorID(u)
	return du
}

// SetNillableDirectContributorID sets the "direct_contributor_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableDirectContributorID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetDirectContributorID(*u)
	}
	return du
}

// ClearDirectContributorID clears the value of the "direct_contributor_id" field.
func (du *DetailUpdate) ClearDirectContributorID() *DetailUpdate {
	du.mutation.ClearDirectContributorID()
	return du
}

// SetGoodID sets the "good_id" field.
func (du *DetailUpdate) SetGoodID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetGoodID(u)
	return du
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableGoodID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetGoodID(*u)
	}
	return du
}

// ClearGoodID clears the value of the "good_id" field.
func (du *DetailUpdate) ClearGoodID() *DetailUpdate {
	du.mutation.ClearGoodID()
	return du
}

// SetOrderID sets the "order_id" field.
func (du *DetailUpdate) SetOrderID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetOrderID(u)
	return du
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableOrderID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetOrderID(*u)
	}
	return du
}

// ClearOrderID clears the value of the "order_id" field.
func (du *DetailUpdate) ClearOrderID() *DetailUpdate {
	du.mutation.ClearOrderID()
	return du
}

// SetSelfOrder sets the "self_order" field.
func (du *DetailUpdate) SetSelfOrder(b bool) *DetailUpdate {
	du.mutation.SetSelfOrder(b)
	return du
}

// SetNillableSelfOrder sets the "self_order" field if the given value is not nil.
func (du *DetailUpdate) SetNillableSelfOrder(b *bool) *DetailUpdate {
	if b != nil {
		du.SetSelfOrder(*b)
	}
	return du
}

// ClearSelfOrder clears the value of the "self_order" field.
func (du *DetailUpdate) ClearSelfOrder() *DetailUpdate {
	du.mutation.ClearSelfOrder()
	return du
}

// SetPaymentID sets the "payment_id" field.
func (du *DetailUpdate) SetPaymentID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetPaymentID(u)
	return du
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillablePaymentID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetPaymentID(*u)
	}
	return du
}

// ClearPaymentID clears the value of the "payment_id" field.
func (du *DetailUpdate) ClearPaymentID() *DetailUpdate {
	du.mutation.ClearPaymentID()
	return du
}

// SetCoinTypeID sets the "coin_type_id" field.
func (du *DetailUpdate) SetCoinTypeID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetCoinTypeID(u)
	return du
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCoinTypeID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetCoinTypeID(*u)
	}
	return du
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (du *DetailUpdate) ClearCoinTypeID() *DetailUpdate {
	du.mutation.ClearCoinTypeID()
	return du
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (du *DetailUpdate) SetPaymentCoinTypeID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetPaymentCoinTypeID(u)
	return du
}

// SetNillablePaymentCoinTypeID sets the "payment_coin_type_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillablePaymentCoinTypeID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetPaymentCoinTypeID(*u)
	}
	return du
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (du *DetailUpdate) ClearPaymentCoinTypeID() *DetailUpdate {
	du.mutation.ClearPaymentCoinTypeID()
	return du
}

// SetPaymentCoinUsdCurrency sets the "payment_coin_usd_currency" field.
func (du *DetailUpdate) SetPaymentCoinUsdCurrency(d decimal.Decimal) *DetailUpdate {
	du.mutation.SetPaymentCoinUsdCurrency(d)
	return du
}

// SetNillablePaymentCoinUsdCurrency sets the "payment_coin_usd_currency" field if the given value is not nil.
func (du *DetailUpdate) SetNillablePaymentCoinUsdCurrency(d *decimal.Decimal) *DetailUpdate {
	if d != nil {
		du.SetPaymentCoinUsdCurrency(*d)
	}
	return du
}

// ClearPaymentCoinUsdCurrency clears the value of the "payment_coin_usd_currency" field.
func (du *DetailUpdate) ClearPaymentCoinUsdCurrency() *DetailUpdate {
	du.mutation.ClearPaymentCoinUsdCurrency()
	return du
}

// SetUnits sets the "units" field.
func (du *DetailUpdate) SetUnits(u uint32) *DetailUpdate {
	du.mutation.ResetUnits()
	du.mutation.SetUnits(u)
	return du
}

// SetNillableUnits sets the "units" field if the given value is not nil.
func (du *DetailUpdate) SetNillableUnits(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetUnits(*u)
	}
	return du
}

// AddUnits adds u to the "units" field.
func (du *DetailUpdate) AddUnits(u int32) *DetailUpdate {
	du.mutation.AddUnits(u)
	return du
}

// ClearUnits clears the value of the "units" field.
func (du *DetailUpdate) ClearUnits() *DetailUpdate {
	du.mutation.ClearUnits()
	return du
}

// SetAmount sets the "amount" field.
func (du *DetailUpdate) SetAmount(d decimal.Decimal) *DetailUpdate {
	du.mutation.SetAmount(d)
	return du
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (du *DetailUpdate) SetNillableAmount(d *decimal.Decimal) *DetailUpdate {
	if d != nil {
		du.SetAmount(*d)
	}
	return du
}

// ClearAmount clears the value of the "amount" field.
func (du *DetailUpdate) ClearAmount() *DetailUpdate {
	du.mutation.ClearAmount()
	return du
}

// SetUsdAmount sets the "usd_amount" field.
func (du *DetailUpdate) SetUsdAmount(d decimal.Decimal) *DetailUpdate {
	du.mutation.SetUsdAmount(d)
	return du
}

// SetNillableUsdAmount sets the "usd_amount" field if the given value is not nil.
func (du *DetailUpdate) SetNillableUsdAmount(d *decimal.Decimal) *DetailUpdate {
	if d != nil {
		du.SetUsdAmount(*d)
	}
	return du
}

// ClearUsdAmount clears the value of the "usd_amount" field.
func (du *DetailUpdate) ClearUsdAmount() *DetailUpdate {
	du.mutation.ClearUsdAmount()
	return du
}

// SetCommission sets the "commission" field.
func (du *DetailUpdate) SetCommission(d decimal.Decimal) *DetailUpdate {
	du.mutation.SetCommission(d)
	return du
}

// SetNillableCommission sets the "commission" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCommission(d *decimal.Decimal) *DetailUpdate {
	if d != nil {
		du.SetCommission(*d)
	}
	return du
}

// ClearCommission clears the value of the "commission" field.
func (du *DetailUpdate) ClearCommission() *DetailUpdate {
	du.mutation.ClearCommission()
	return du
}

// Mutation returns the DetailMutation object of the builder.
func (du *DetailUpdate) Mutation() *DetailMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := du.defaults(); err != nil {
		return 0, err
	}
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DetailUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DetailUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DetailUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DetailUpdate) defaults() error {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		if detail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (du *DetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldAppID,
		})
	}
	if du.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldAppID,
		})
	}
	if value, ok := du.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldUserID,
		})
	}
	if du.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldUserID,
		})
	}
	if value, ok := du.mutation.DirectContributorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldDirectContributorID,
		})
	}
	if du.mutation.DirectContributorIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldDirectContributorID,
		})
	}
	if value, ok := du.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldGoodID,
		})
	}
	if du.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldGoodID,
		})
	}
	if value, ok := du.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldOrderID,
		})
	}
	if du.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldOrderID,
		})
	}
	if value, ok := du.mutation.SelfOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: detail.FieldSelfOrder,
		})
	}
	if du.mutation.SelfOrderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: detail.FieldSelfOrder,
		})
	}
	if value, ok := du.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldPaymentID,
		})
	}
	if du.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldPaymentID,
		})
	}
	if value, ok := du.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldCoinTypeID,
		})
	}
	if du.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldCoinTypeID,
		})
	}
	if value, ok := du.mutation.PaymentCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldPaymentCoinTypeID,
		})
	}
	if du.mutation.PaymentCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldPaymentCoinTypeID,
		})
	}
	if value, ok := du.mutation.PaymentCoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldPaymentCoinUsdCurrency,
		})
	}
	if du.mutation.PaymentCoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldPaymentCoinUsdCurrency,
		})
	}
	if value, ok := du.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUnits,
		})
	}
	if value, ok := du.mutation.AddedUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUnits,
		})
	}
	if du.mutation.UnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: detail.FieldUnits,
		})
	}
	if value, ok := du.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldAmount,
		})
	}
	if du.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldAmount,
		})
	}
	if value, ok := du.mutation.UsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldUsdAmount,
		})
	}
	if du.mutation.UsdAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldUsdAmount,
		})
	}
	if value, ok := du.mutation.Commission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldCommission,
		})
	}
	if du.mutation.CommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldCommission,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{detail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DetailUpdateOne is the builder for updating a single Detail entity.
type DetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DetailMutation
}

// SetCreatedAt sets the "created_at" field.
func (duo *DetailUpdateOne) SetCreatedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetCreatedAt()
	duo.mutation.SetCreatedAt(u)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCreatedAt(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetCreatedAt(*u)
	}
	return duo
}

// AddCreatedAt adds u to the "created_at" field.
func (duo *DetailUpdateOne) AddCreatedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddCreatedAt(u)
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DetailUpdateOne) SetUpdatedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetUpdatedAt()
	duo.mutation.SetUpdatedAt(u)
	return duo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (duo *DetailUpdateOne) AddUpdatedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddUpdatedAt(u)
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DetailUpdateOne) SetDeletedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetDeletedAt()
	duo.mutation.SetDeletedAt(u)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableDeletedAt(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetDeletedAt(*u)
	}
	return duo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (duo *DetailUpdateOne) AddDeletedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddDeletedAt(u)
	return duo
}

// SetAppID sets the "app_id" field.
func (duo *DetailUpdateOne) SetAppID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetAppID(u)
	return duo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableAppID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetAppID(*u)
	}
	return duo
}

// ClearAppID clears the value of the "app_id" field.
func (duo *DetailUpdateOne) ClearAppID() *DetailUpdateOne {
	duo.mutation.ClearAppID()
	return duo
}

// SetUserID sets the "user_id" field.
func (duo *DetailUpdateOne) SetUserID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetUserID(u)
	return duo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableUserID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetUserID(*u)
	}
	return duo
}

// ClearUserID clears the value of the "user_id" field.
func (duo *DetailUpdateOne) ClearUserID() *DetailUpdateOne {
	duo.mutation.ClearUserID()
	return duo
}

// SetDirectContributorID sets the "direct_contributor_id" field.
func (duo *DetailUpdateOne) SetDirectContributorID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetDirectContributorID(u)
	return duo
}

// SetNillableDirectContributorID sets the "direct_contributor_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableDirectContributorID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetDirectContributorID(*u)
	}
	return duo
}

// ClearDirectContributorID clears the value of the "direct_contributor_id" field.
func (duo *DetailUpdateOne) ClearDirectContributorID() *DetailUpdateOne {
	duo.mutation.ClearDirectContributorID()
	return duo
}

// SetGoodID sets the "good_id" field.
func (duo *DetailUpdateOne) SetGoodID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetGoodID(u)
	return duo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableGoodID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetGoodID(*u)
	}
	return duo
}

// ClearGoodID clears the value of the "good_id" field.
func (duo *DetailUpdateOne) ClearGoodID() *DetailUpdateOne {
	duo.mutation.ClearGoodID()
	return duo
}

// SetOrderID sets the "order_id" field.
func (duo *DetailUpdateOne) SetOrderID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetOrderID(u)
	return duo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableOrderID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetOrderID(*u)
	}
	return duo
}

// ClearOrderID clears the value of the "order_id" field.
func (duo *DetailUpdateOne) ClearOrderID() *DetailUpdateOne {
	duo.mutation.ClearOrderID()
	return duo
}

// SetSelfOrder sets the "self_order" field.
func (duo *DetailUpdateOne) SetSelfOrder(b bool) *DetailUpdateOne {
	duo.mutation.SetSelfOrder(b)
	return duo
}

// SetNillableSelfOrder sets the "self_order" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableSelfOrder(b *bool) *DetailUpdateOne {
	if b != nil {
		duo.SetSelfOrder(*b)
	}
	return duo
}

// ClearSelfOrder clears the value of the "self_order" field.
func (duo *DetailUpdateOne) ClearSelfOrder() *DetailUpdateOne {
	duo.mutation.ClearSelfOrder()
	return duo
}

// SetPaymentID sets the "payment_id" field.
func (duo *DetailUpdateOne) SetPaymentID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetPaymentID(u)
	return duo
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillablePaymentID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetPaymentID(*u)
	}
	return duo
}

// ClearPaymentID clears the value of the "payment_id" field.
func (duo *DetailUpdateOne) ClearPaymentID() *DetailUpdateOne {
	duo.mutation.ClearPaymentID()
	return duo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (duo *DetailUpdateOne) SetCoinTypeID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetCoinTypeID(u)
	return duo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetCoinTypeID(*u)
	}
	return duo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (duo *DetailUpdateOne) ClearCoinTypeID() *DetailUpdateOne {
	duo.mutation.ClearCoinTypeID()
	return duo
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (duo *DetailUpdateOne) SetPaymentCoinTypeID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetPaymentCoinTypeID(u)
	return duo
}

// SetNillablePaymentCoinTypeID sets the "payment_coin_type_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillablePaymentCoinTypeID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetPaymentCoinTypeID(*u)
	}
	return duo
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (duo *DetailUpdateOne) ClearPaymentCoinTypeID() *DetailUpdateOne {
	duo.mutation.ClearPaymentCoinTypeID()
	return duo
}

// SetPaymentCoinUsdCurrency sets the "payment_coin_usd_currency" field.
func (duo *DetailUpdateOne) SetPaymentCoinUsdCurrency(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.SetPaymentCoinUsdCurrency(d)
	return duo
}

// SetNillablePaymentCoinUsdCurrency sets the "payment_coin_usd_currency" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillablePaymentCoinUsdCurrency(d *decimal.Decimal) *DetailUpdateOne {
	if d != nil {
		duo.SetPaymentCoinUsdCurrency(*d)
	}
	return duo
}

// ClearPaymentCoinUsdCurrency clears the value of the "payment_coin_usd_currency" field.
func (duo *DetailUpdateOne) ClearPaymentCoinUsdCurrency() *DetailUpdateOne {
	duo.mutation.ClearPaymentCoinUsdCurrency()
	return duo
}

// SetUnits sets the "units" field.
func (duo *DetailUpdateOne) SetUnits(u uint32) *DetailUpdateOne {
	duo.mutation.ResetUnits()
	duo.mutation.SetUnits(u)
	return duo
}

// SetNillableUnits sets the "units" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableUnits(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetUnits(*u)
	}
	return duo
}

// AddUnits adds u to the "units" field.
func (duo *DetailUpdateOne) AddUnits(u int32) *DetailUpdateOne {
	duo.mutation.AddUnits(u)
	return duo
}

// ClearUnits clears the value of the "units" field.
func (duo *DetailUpdateOne) ClearUnits() *DetailUpdateOne {
	duo.mutation.ClearUnits()
	return duo
}

// SetAmount sets the "amount" field.
func (duo *DetailUpdateOne) SetAmount(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.SetAmount(d)
	return duo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableAmount(d *decimal.Decimal) *DetailUpdateOne {
	if d != nil {
		duo.SetAmount(*d)
	}
	return duo
}

// ClearAmount clears the value of the "amount" field.
func (duo *DetailUpdateOne) ClearAmount() *DetailUpdateOne {
	duo.mutation.ClearAmount()
	return duo
}

// SetUsdAmount sets the "usd_amount" field.
func (duo *DetailUpdateOne) SetUsdAmount(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.SetUsdAmount(d)
	return duo
}

// SetNillableUsdAmount sets the "usd_amount" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableUsdAmount(d *decimal.Decimal) *DetailUpdateOne {
	if d != nil {
		duo.SetUsdAmount(*d)
	}
	return duo
}

// ClearUsdAmount clears the value of the "usd_amount" field.
func (duo *DetailUpdateOne) ClearUsdAmount() *DetailUpdateOne {
	duo.mutation.ClearUsdAmount()
	return duo
}

// SetCommission sets the "commission" field.
func (duo *DetailUpdateOne) SetCommission(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.SetCommission(d)
	return duo
}

// SetNillableCommission sets the "commission" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCommission(d *decimal.Decimal) *DetailUpdateOne {
	if d != nil {
		duo.SetCommission(*d)
	}
	return duo
}

// ClearCommission clears the value of the "commission" field.
func (duo *DetailUpdateOne) ClearCommission() *DetailUpdateOne {
	duo.mutation.ClearCommission()
	return duo
}

// Mutation returns the DetailMutation object of the builder.
func (duo *DetailUpdateOne) Mutation() *DetailMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DetailUpdateOne) Select(field string, fields ...string) *DetailUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Detail entity.
func (duo *DetailUpdateOne) Save(ctx context.Context) (*Detail, error) {
	var (
		err  error
		node *Detail
	)
	if err := duo.defaults(); err != nil {
		return nil, err
	}
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DetailUpdateOne) SaveX(ctx context.Context) *Detail {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DetailUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DetailUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DetailUpdateOne) defaults() error {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		if detail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (duo *DetailUpdateOne) sqlSave(ctx context.Context) (_node *Detail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Detail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, detail.FieldID)
		for _, f := range fields {
			if !detail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != detail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldAppID,
		})
	}
	if duo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldAppID,
		})
	}
	if value, ok := duo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldUserID,
		})
	}
	if duo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldUserID,
		})
	}
	if value, ok := duo.mutation.DirectContributorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldDirectContributorID,
		})
	}
	if duo.mutation.DirectContributorIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldDirectContributorID,
		})
	}
	if value, ok := duo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldGoodID,
		})
	}
	if duo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldGoodID,
		})
	}
	if value, ok := duo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldOrderID,
		})
	}
	if duo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldOrderID,
		})
	}
	if value, ok := duo.mutation.SelfOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: detail.FieldSelfOrder,
		})
	}
	if duo.mutation.SelfOrderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: detail.FieldSelfOrder,
		})
	}
	if value, ok := duo.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldPaymentID,
		})
	}
	if duo.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldPaymentID,
		})
	}
	if value, ok := duo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldCoinTypeID,
		})
	}
	if duo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldCoinTypeID,
		})
	}
	if value, ok := duo.mutation.PaymentCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldPaymentCoinTypeID,
		})
	}
	if duo.mutation.PaymentCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldPaymentCoinTypeID,
		})
	}
	if value, ok := duo.mutation.PaymentCoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldPaymentCoinUsdCurrency,
		})
	}
	if duo.mutation.PaymentCoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldPaymentCoinUsdCurrency,
		})
	}
	if value, ok := duo.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUnits,
		})
	}
	if value, ok := duo.mutation.AddedUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUnits,
		})
	}
	if duo.mutation.UnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: detail.FieldUnits,
		})
	}
	if value, ok := duo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldAmount,
		})
	}
	if duo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldAmount,
		})
	}
	if value, ok := duo.mutation.UsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldUsdAmount,
		})
	}
	if duo.mutation.UsdAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldUsdAmount,
		})
	}
	if value, ok := duo.mutation.Commission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: detail.FieldCommission,
		})
	}
	if duo.mutation.CommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: detail.FieldCommission,
		})
	}
	_node = &Detail{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{detail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
