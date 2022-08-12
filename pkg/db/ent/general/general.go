// Code generated by entc, DO NOT EDIT.

package general

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the general type in the database.
	Label = "general"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldTotalUnits holds the string denoting the total_units field in the database.
	FieldTotalUnits = "total_units"
	// FieldSelfUnits holds the string denoting the self_units field in the database.
	FieldSelfUnits = "self_units"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldSelfAmount holds the string denoting the self_amount field in the database.
	FieldSelfAmount = "self_amount"
	// FieldTotalCommission holds the string denoting the total_commission field in the database.
	FieldTotalCommission = "total_commission"
	// FieldSelfCommission holds the string denoting the self_commission field in the database.
	FieldSelfCommission = "self_commission"
	// FieldSuperiorCommission holds the string denoting the superior_commission field in the database.
	FieldSuperiorCommission = "superior_commission"
	// Table holds the table name of the general in the database.
	Table = "generals"
)

// Columns holds all SQL columns for general fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldCoinTypeID,
	FieldTotalUnits,
	FieldSelfUnits,
	FieldTotalAmount,
	FieldSelfAmount,
	FieldTotalCommission,
	FieldSelfCommission,
	FieldSuperiorCommission,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/archivement-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultTotalUnits holds the default value on creation for the "total_units" field.
	DefaultTotalUnits uint32
	// DefaultSelfUnits holds the default value on creation for the "self_units" field.
	DefaultSelfUnits uint32
	// DefaultTotalAmount holds the default value on creation for the "total_amount" field.
	DefaultTotalAmount decimal.Decimal
	// DefaultSelfAmount holds the default value on creation for the "self_amount" field.
	DefaultSelfAmount decimal.Decimal
	// DefaultTotalCommission holds the default value on creation for the "total_commission" field.
	DefaultTotalCommission decimal.Decimal
	// DefaultSelfCommission holds the default value on creation for the "self_commission" field.
	DefaultSelfCommission decimal.Decimal
	// DefaultSuperiorCommission holds the default value on creation for the "superior_commission" field.
	DefaultSuperiorCommission decimal.Decimal
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
