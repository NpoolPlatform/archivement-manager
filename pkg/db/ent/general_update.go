// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GeneralUpdate is the builder for updating General entities.
type GeneralUpdate struct {
	config
	hooks    []Hook
	mutation *GeneralMutation
}

// Where appends a list predicates to the GeneralUpdate builder.
func (gu *GeneralUpdate) Where(ps ...predicate.General) *GeneralUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GeneralUpdate) SetCreatedAt(u uint32) *GeneralUpdate {
	gu.mutation.ResetCreatedAt()
	gu.mutation.SetCreatedAt(u)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableCreatedAt(u *uint32) *GeneralUpdate {
	if u != nil {
		gu.SetCreatedAt(*u)
	}
	return gu
}

// AddCreatedAt adds u to the "created_at" field.
func (gu *GeneralUpdate) AddCreatedAt(u int32) *GeneralUpdate {
	gu.mutation.AddCreatedAt(u)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GeneralUpdate) SetUpdatedAt(u uint32) *GeneralUpdate {
	gu.mutation.ResetUpdatedAt()
	gu.mutation.SetUpdatedAt(u)
	return gu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gu *GeneralUpdate) AddUpdatedAt(u int32) *GeneralUpdate {
	gu.mutation.AddUpdatedAt(u)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GeneralUpdate) SetDeletedAt(u uint32) *GeneralUpdate {
	gu.mutation.ResetDeletedAt()
	gu.mutation.SetDeletedAt(u)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableDeletedAt(u *uint32) *GeneralUpdate {
	if u != nil {
		gu.SetDeletedAt(*u)
	}
	return gu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gu *GeneralUpdate) AddDeletedAt(u int32) *GeneralUpdate {
	gu.mutation.AddDeletedAt(u)
	return gu
}

// SetAppID sets the "app_id" field.
func (gu *GeneralUpdate) SetAppID(u uuid.UUID) *GeneralUpdate {
	gu.mutation.SetAppID(u)
	return gu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableAppID(u *uuid.UUID) *GeneralUpdate {
	if u != nil {
		gu.SetAppID(*u)
	}
	return gu
}

// ClearAppID clears the value of the "app_id" field.
func (gu *GeneralUpdate) ClearAppID() *GeneralUpdate {
	gu.mutation.ClearAppID()
	return gu
}

// SetUserID sets the "user_id" field.
func (gu *GeneralUpdate) SetUserID(u uuid.UUID) *GeneralUpdate {
	gu.mutation.SetUserID(u)
	return gu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableUserID(u *uuid.UUID) *GeneralUpdate {
	if u != nil {
		gu.SetUserID(*u)
	}
	return gu
}

// ClearUserID clears the value of the "user_id" field.
func (gu *GeneralUpdate) ClearUserID() *GeneralUpdate {
	gu.mutation.ClearUserID()
	return gu
}

// SetGoodID sets the "good_id" field.
func (gu *GeneralUpdate) SetGoodID(u uuid.UUID) *GeneralUpdate {
	gu.mutation.SetGoodID(u)
	return gu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableGoodID(u *uuid.UUID) *GeneralUpdate {
	if u != nil {
		gu.SetGoodID(*u)
	}
	return gu
}

// ClearGoodID clears the value of the "good_id" field.
func (gu *GeneralUpdate) ClearGoodID() *GeneralUpdate {
	gu.mutation.ClearGoodID()
	return gu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gu *GeneralUpdate) SetCoinTypeID(u uuid.UUID) *GeneralUpdate {
	gu.mutation.SetCoinTypeID(u)
	return gu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableCoinTypeID(u *uuid.UUID) *GeneralUpdate {
	if u != nil {
		gu.SetCoinTypeID(*u)
	}
	return gu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gu *GeneralUpdate) ClearCoinTypeID() *GeneralUpdate {
	gu.mutation.ClearCoinTypeID()
	return gu
}

// SetTotalUnits sets the "total_units" field.
func (gu *GeneralUpdate) SetTotalUnits(u uint32) *GeneralUpdate {
	gu.mutation.ResetTotalUnits()
	gu.mutation.SetTotalUnits(u)
	return gu
}

// SetNillableTotalUnits sets the "total_units" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableTotalUnits(u *uint32) *GeneralUpdate {
	if u != nil {
		gu.SetTotalUnits(*u)
	}
	return gu
}

// AddTotalUnits adds u to the "total_units" field.
func (gu *GeneralUpdate) AddTotalUnits(u int32) *GeneralUpdate {
	gu.mutation.AddTotalUnits(u)
	return gu
}

// ClearTotalUnits clears the value of the "total_units" field.
func (gu *GeneralUpdate) ClearTotalUnits() *GeneralUpdate {
	gu.mutation.ClearTotalUnits()
	return gu
}

// SetSelfUnits sets the "self_units" field.
func (gu *GeneralUpdate) SetSelfUnits(u uint32) *GeneralUpdate {
	gu.mutation.ResetSelfUnits()
	gu.mutation.SetSelfUnits(u)
	return gu
}

// SetNillableSelfUnits sets the "self_units" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableSelfUnits(u *uint32) *GeneralUpdate {
	if u != nil {
		gu.SetSelfUnits(*u)
	}
	return gu
}

// AddSelfUnits adds u to the "self_units" field.
func (gu *GeneralUpdate) AddSelfUnits(u int32) *GeneralUpdate {
	gu.mutation.AddSelfUnits(u)
	return gu
}

// ClearSelfUnits clears the value of the "self_units" field.
func (gu *GeneralUpdate) ClearSelfUnits() *GeneralUpdate {
	gu.mutation.ClearSelfUnits()
	return gu
}

// SetTotalAmount sets the "total_amount" field.
func (gu *GeneralUpdate) SetTotalAmount(d decimal.Decimal) *GeneralUpdate {
	gu.mutation.SetTotalAmount(d)
	return gu
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableTotalAmount(d *decimal.Decimal) *GeneralUpdate {
	if d != nil {
		gu.SetTotalAmount(*d)
	}
	return gu
}

// ClearTotalAmount clears the value of the "total_amount" field.
func (gu *GeneralUpdate) ClearTotalAmount() *GeneralUpdate {
	gu.mutation.ClearTotalAmount()
	return gu
}

// SetSelfAmount sets the "self_amount" field.
func (gu *GeneralUpdate) SetSelfAmount(d decimal.Decimal) *GeneralUpdate {
	gu.mutation.SetSelfAmount(d)
	return gu
}

// SetNillableSelfAmount sets the "self_amount" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableSelfAmount(d *decimal.Decimal) *GeneralUpdate {
	if d != nil {
		gu.SetSelfAmount(*d)
	}
	return gu
}

// ClearSelfAmount clears the value of the "self_amount" field.
func (gu *GeneralUpdate) ClearSelfAmount() *GeneralUpdate {
	gu.mutation.ClearSelfAmount()
	return gu
}

// SetTotalCommission sets the "total_commission" field.
func (gu *GeneralUpdate) SetTotalCommission(d decimal.Decimal) *GeneralUpdate {
	gu.mutation.SetTotalCommission(d)
	return gu
}

// SetNillableTotalCommission sets the "total_commission" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableTotalCommission(d *decimal.Decimal) *GeneralUpdate {
	if d != nil {
		gu.SetTotalCommission(*d)
	}
	return gu
}

// ClearTotalCommission clears the value of the "total_commission" field.
func (gu *GeneralUpdate) ClearTotalCommission() *GeneralUpdate {
	gu.mutation.ClearTotalCommission()
	return gu
}

// SetSelfCommission sets the "self_commission" field.
func (gu *GeneralUpdate) SetSelfCommission(d decimal.Decimal) *GeneralUpdate {
	gu.mutation.SetSelfCommission(d)
	return gu
}

// SetNillableSelfCommission sets the "self_commission" field if the given value is not nil.
func (gu *GeneralUpdate) SetNillableSelfCommission(d *decimal.Decimal) *GeneralUpdate {
	if d != nil {
		gu.SetSelfCommission(*d)
	}
	return gu
}

// ClearSelfCommission clears the value of the "self_commission" field.
func (gu *GeneralUpdate) ClearSelfCommission() *GeneralUpdate {
	gu.mutation.ClearSelfCommission()
	return gu
}

// Mutation returns the GeneralMutation object of the builder.
func (gu *GeneralUpdate) Mutation() *GeneralMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GeneralUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := gu.defaults(); err != nil {
		return 0, err
	}
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GeneralMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GeneralUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GeneralUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GeneralUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GeneralUpdate) defaults() error {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		if general.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized general.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := general.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (gu *GeneralUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   general.Table,
			Columns: general.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: general.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldCreatedAt,
		})
	}
	if value, ok := gu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldCreatedAt,
		})
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldDeletedAt,
		})
	}
	if value, ok := gu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldDeletedAt,
		})
	}
	if value, ok := gu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldAppID,
		})
	}
	if gu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldAppID,
		})
	}
	if value, ok := gu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldUserID,
		})
	}
	if gu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldUserID,
		})
	}
	if value, ok := gu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldGoodID,
		})
	}
	if gu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldGoodID,
		})
	}
	if value, ok := gu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldCoinTypeID,
		})
	}
	if gu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldCoinTypeID,
		})
	}
	if value, ok := gu.mutation.TotalUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldTotalUnits,
		})
	}
	if value, ok := gu.mutation.AddedTotalUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldTotalUnits,
		})
	}
	if gu.mutation.TotalUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: general.FieldTotalUnits,
		})
	}
	if value, ok := gu.mutation.SelfUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldSelfUnits,
		})
	}
	if value, ok := gu.mutation.AddedSelfUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldSelfUnits,
		})
	}
	if gu.mutation.SelfUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: general.FieldSelfUnits,
		})
	}
	if value, ok := gu.mutation.TotalAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldTotalAmount,
		})
	}
	if gu.mutation.TotalAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldTotalAmount,
		})
	}
	if value, ok := gu.mutation.SelfAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldSelfAmount,
		})
	}
	if gu.mutation.SelfAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldSelfAmount,
		})
	}
	if value, ok := gu.mutation.TotalCommission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldTotalCommission,
		})
	}
	if gu.mutation.TotalCommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldTotalCommission,
		})
	}
	if value, ok := gu.mutation.SelfCommission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldSelfCommission,
		})
	}
	if gu.mutation.SelfCommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldSelfCommission,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{general.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GeneralUpdateOne is the builder for updating a single General entity.
type GeneralUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GeneralMutation
}

// SetCreatedAt sets the "created_at" field.
func (guo *GeneralUpdateOne) SetCreatedAt(u uint32) *GeneralUpdateOne {
	guo.mutation.ResetCreatedAt()
	guo.mutation.SetCreatedAt(u)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableCreatedAt(u *uint32) *GeneralUpdateOne {
	if u != nil {
		guo.SetCreatedAt(*u)
	}
	return guo
}

// AddCreatedAt adds u to the "created_at" field.
func (guo *GeneralUpdateOne) AddCreatedAt(u int32) *GeneralUpdateOne {
	guo.mutation.AddCreatedAt(u)
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GeneralUpdateOne) SetUpdatedAt(u uint32) *GeneralUpdateOne {
	guo.mutation.ResetUpdatedAt()
	guo.mutation.SetUpdatedAt(u)
	return guo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (guo *GeneralUpdateOne) AddUpdatedAt(u int32) *GeneralUpdateOne {
	guo.mutation.AddUpdatedAt(u)
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GeneralUpdateOne) SetDeletedAt(u uint32) *GeneralUpdateOne {
	guo.mutation.ResetDeletedAt()
	guo.mutation.SetDeletedAt(u)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableDeletedAt(u *uint32) *GeneralUpdateOne {
	if u != nil {
		guo.SetDeletedAt(*u)
	}
	return guo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (guo *GeneralUpdateOne) AddDeletedAt(u int32) *GeneralUpdateOne {
	guo.mutation.AddDeletedAt(u)
	return guo
}

// SetAppID sets the "app_id" field.
func (guo *GeneralUpdateOne) SetAppID(u uuid.UUID) *GeneralUpdateOne {
	guo.mutation.SetAppID(u)
	return guo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableAppID(u *uuid.UUID) *GeneralUpdateOne {
	if u != nil {
		guo.SetAppID(*u)
	}
	return guo
}

// ClearAppID clears the value of the "app_id" field.
func (guo *GeneralUpdateOne) ClearAppID() *GeneralUpdateOne {
	guo.mutation.ClearAppID()
	return guo
}

// SetUserID sets the "user_id" field.
func (guo *GeneralUpdateOne) SetUserID(u uuid.UUID) *GeneralUpdateOne {
	guo.mutation.SetUserID(u)
	return guo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableUserID(u *uuid.UUID) *GeneralUpdateOne {
	if u != nil {
		guo.SetUserID(*u)
	}
	return guo
}

// ClearUserID clears the value of the "user_id" field.
func (guo *GeneralUpdateOne) ClearUserID() *GeneralUpdateOne {
	guo.mutation.ClearUserID()
	return guo
}

// SetGoodID sets the "good_id" field.
func (guo *GeneralUpdateOne) SetGoodID(u uuid.UUID) *GeneralUpdateOne {
	guo.mutation.SetGoodID(u)
	return guo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableGoodID(u *uuid.UUID) *GeneralUpdateOne {
	if u != nil {
		guo.SetGoodID(*u)
	}
	return guo
}

// ClearGoodID clears the value of the "good_id" field.
func (guo *GeneralUpdateOne) ClearGoodID() *GeneralUpdateOne {
	guo.mutation.ClearGoodID()
	return guo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (guo *GeneralUpdateOne) SetCoinTypeID(u uuid.UUID) *GeneralUpdateOne {
	guo.mutation.SetCoinTypeID(u)
	return guo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *GeneralUpdateOne {
	if u != nil {
		guo.SetCoinTypeID(*u)
	}
	return guo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (guo *GeneralUpdateOne) ClearCoinTypeID() *GeneralUpdateOne {
	guo.mutation.ClearCoinTypeID()
	return guo
}

// SetTotalUnits sets the "total_units" field.
func (guo *GeneralUpdateOne) SetTotalUnits(u uint32) *GeneralUpdateOne {
	guo.mutation.ResetTotalUnits()
	guo.mutation.SetTotalUnits(u)
	return guo
}

// SetNillableTotalUnits sets the "total_units" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableTotalUnits(u *uint32) *GeneralUpdateOne {
	if u != nil {
		guo.SetTotalUnits(*u)
	}
	return guo
}

// AddTotalUnits adds u to the "total_units" field.
func (guo *GeneralUpdateOne) AddTotalUnits(u int32) *GeneralUpdateOne {
	guo.mutation.AddTotalUnits(u)
	return guo
}

// ClearTotalUnits clears the value of the "total_units" field.
func (guo *GeneralUpdateOne) ClearTotalUnits() *GeneralUpdateOne {
	guo.mutation.ClearTotalUnits()
	return guo
}

// SetSelfUnits sets the "self_units" field.
func (guo *GeneralUpdateOne) SetSelfUnits(u uint32) *GeneralUpdateOne {
	guo.mutation.ResetSelfUnits()
	guo.mutation.SetSelfUnits(u)
	return guo
}

// SetNillableSelfUnits sets the "self_units" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableSelfUnits(u *uint32) *GeneralUpdateOne {
	if u != nil {
		guo.SetSelfUnits(*u)
	}
	return guo
}

// AddSelfUnits adds u to the "self_units" field.
func (guo *GeneralUpdateOne) AddSelfUnits(u int32) *GeneralUpdateOne {
	guo.mutation.AddSelfUnits(u)
	return guo
}

// ClearSelfUnits clears the value of the "self_units" field.
func (guo *GeneralUpdateOne) ClearSelfUnits() *GeneralUpdateOne {
	guo.mutation.ClearSelfUnits()
	return guo
}

// SetTotalAmount sets the "total_amount" field.
func (guo *GeneralUpdateOne) SetTotalAmount(d decimal.Decimal) *GeneralUpdateOne {
	guo.mutation.SetTotalAmount(d)
	return guo
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableTotalAmount(d *decimal.Decimal) *GeneralUpdateOne {
	if d != nil {
		guo.SetTotalAmount(*d)
	}
	return guo
}

// ClearTotalAmount clears the value of the "total_amount" field.
func (guo *GeneralUpdateOne) ClearTotalAmount() *GeneralUpdateOne {
	guo.mutation.ClearTotalAmount()
	return guo
}

// SetSelfAmount sets the "self_amount" field.
func (guo *GeneralUpdateOne) SetSelfAmount(d decimal.Decimal) *GeneralUpdateOne {
	guo.mutation.SetSelfAmount(d)
	return guo
}

// SetNillableSelfAmount sets the "self_amount" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableSelfAmount(d *decimal.Decimal) *GeneralUpdateOne {
	if d != nil {
		guo.SetSelfAmount(*d)
	}
	return guo
}

// ClearSelfAmount clears the value of the "self_amount" field.
func (guo *GeneralUpdateOne) ClearSelfAmount() *GeneralUpdateOne {
	guo.mutation.ClearSelfAmount()
	return guo
}

// SetTotalCommission sets the "total_commission" field.
func (guo *GeneralUpdateOne) SetTotalCommission(d decimal.Decimal) *GeneralUpdateOne {
	guo.mutation.SetTotalCommission(d)
	return guo
}

// SetNillableTotalCommission sets the "total_commission" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableTotalCommission(d *decimal.Decimal) *GeneralUpdateOne {
	if d != nil {
		guo.SetTotalCommission(*d)
	}
	return guo
}

// ClearTotalCommission clears the value of the "total_commission" field.
func (guo *GeneralUpdateOne) ClearTotalCommission() *GeneralUpdateOne {
	guo.mutation.ClearTotalCommission()
	return guo
}

// SetSelfCommission sets the "self_commission" field.
func (guo *GeneralUpdateOne) SetSelfCommission(d decimal.Decimal) *GeneralUpdateOne {
	guo.mutation.SetSelfCommission(d)
	return guo
}

// SetNillableSelfCommission sets the "self_commission" field if the given value is not nil.
func (guo *GeneralUpdateOne) SetNillableSelfCommission(d *decimal.Decimal) *GeneralUpdateOne {
	if d != nil {
		guo.SetSelfCommission(*d)
	}
	return guo
}

// ClearSelfCommission clears the value of the "self_commission" field.
func (guo *GeneralUpdateOne) ClearSelfCommission() *GeneralUpdateOne {
	guo.mutation.ClearSelfCommission()
	return guo
}

// Mutation returns the GeneralMutation object of the builder.
func (guo *GeneralUpdateOne) Mutation() *GeneralMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GeneralUpdateOne) Select(field string, fields ...string) *GeneralUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated General entity.
func (guo *GeneralUpdateOne) Save(ctx context.Context) (*General, error) {
	var (
		err  error
		node *General
	)
	if err := guo.defaults(); err != nil {
		return nil, err
	}
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GeneralMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GeneralUpdateOne) SaveX(ctx context.Context) *General {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GeneralUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GeneralUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GeneralUpdateOne) defaults() error {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		if general.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized general.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := general.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (guo *GeneralUpdateOne) sqlSave(ctx context.Context) (_node *General, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   general.Table,
			Columns: general.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: general.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "General.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, general.FieldID)
		for _, f := range fields {
			if !general.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != general.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldCreatedAt,
		})
	}
	if value, ok := guo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldCreatedAt,
		})
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldDeletedAt,
		})
	}
	if value, ok := guo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldDeletedAt,
		})
	}
	if value, ok := guo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldAppID,
		})
	}
	if guo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldAppID,
		})
	}
	if value, ok := guo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldUserID,
		})
	}
	if guo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldUserID,
		})
	}
	if value, ok := guo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldGoodID,
		})
	}
	if guo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldGoodID,
		})
	}
	if value, ok := guo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: general.FieldCoinTypeID,
		})
	}
	if guo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: general.FieldCoinTypeID,
		})
	}
	if value, ok := guo.mutation.TotalUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldTotalUnits,
		})
	}
	if value, ok := guo.mutation.AddedTotalUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldTotalUnits,
		})
	}
	if guo.mutation.TotalUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: general.FieldTotalUnits,
		})
	}
	if value, ok := guo.mutation.SelfUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldSelfUnits,
		})
	}
	if value, ok := guo.mutation.AddedSelfUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: general.FieldSelfUnits,
		})
	}
	if guo.mutation.SelfUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: general.FieldSelfUnits,
		})
	}
	if value, ok := guo.mutation.TotalAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldTotalAmount,
		})
	}
	if guo.mutation.TotalAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldTotalAmount,
		})
	}
	if value, ok := guo.mutation.SelfAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldSelfAmount,
		})
	}
	if guo.mutation.SelfAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldSelfAmount,
		})
	}
	if value, ok := guo.mutation.TotalCommission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldTotalCommission,
		})
	}
	if guo.mutation.TotalCommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldTotalCommission,
		})
	}
	if value, ok := guo.mutation.SelfCommission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: general.FieldSelfCommission,
		})
	}
	if guo.mutation.SelfCommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: general.FieldSelfCommission,
		})
	}
	_node = &General{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{general.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
