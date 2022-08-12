// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DetailsColumns holds the columns for the "details" table.
	DetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "direct_contributor_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "self_order", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "payment_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "payment_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "usd_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// DetailsTable holds the schema information for the "details" table.
	DetailsTable = &schema.Table{
		Name:       "details",
		Columns:    DetailsColumns,
		PrimaryKey: []*schema.Column{DetailsColumns[0]},
	}
	// GeneralsColumns holds the columns for the "generals" table.
	GeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "total_units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "self_units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "total_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "total_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "superior_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// GeneralsTable holds the schema information for the "generals" table.
	GeneralsTable = &schema.Table{
		Name:       "generals",
		Columns:    GeneralsColumns,
		PrimaryKey: []*schema.Column{GeneralsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DetailsTable,
		GeneralsTable,
	}
)

func init() {
}
