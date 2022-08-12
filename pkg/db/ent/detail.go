// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/detail"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Detail is the model entity for the Detail schema.
type Detail struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// DirectContributorID holds the value of the "direct_contributor_id" field.
	DirectContributorID uuid.UUID `json:"direct_contributor_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// SelfOrder holds the value of the "self_order" field.
	SelfOrder bool `json:"self_order,omitempty"`
	// PaymentID holds the value of the "payment_id" field.
	PaymentID uuid.UUID `json:"payment_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// PaymentCoinTypeID holds the value of the "payment_coin_type_id" field.
	PaymentCoinTypeID uuid.UUID `json:"payment_coin_type_id,omitempty"`
	// PaymentCoinUsdCurrency holds the value of the "payment_coin_usd_currency" field.
	PaymentCoinUsdCurrency decimal.Decimal `json:"payment_coin_usd_currency,omitempty"`
	// Units holds the value of the "units" field.
	Units uint32 `json:"units,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// UsdAmount holds the value of the "usd_amount" field.
	UsdAmount decimal.Decimal `json:"usd_amount,omitempty"`
	// Commission holds the value of the "commission" field.
	Commission decimal.Decimal `json:"commission,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Detail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case detail.FieldPaymentCoinUsdCurrency, detail.FieldAmount, detail.FieldUsdAmount, detail.FieldCommission:
			values[i] = new(decimal.Decimal)
		case detail.FieldSelfOrder:
			values[i] = new(sql.NullBool)
		case detail.FieldCreatedAt, detail.FieldUpdatedAt, detail.FieldDeletedAt, detail.FieldUnits:
			values[i] = new(sql.NullInt64)
		case detail.FieldID, detail.FieldAppID, detail.FieldUserID, detail.FieldDirectContributorID, detail.FieldGoodID, detail.FieldOrderID, detail.FieldPaymentID, detail.FieldCoinTypeID, detail.FieldPaymentCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Detail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Detail fields.
func (d *Detail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case detail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case detail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = uint32(value.Int64)
			}
		case detail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = uint32(value.Int64)
			}
		case detail.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = uint32(value.Int64)
			}
		case detail.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				d.AppID = *value
			}
		case detail.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				d.UserID = *value
			}
		case detail.FieldDirectContributorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field direct_contributor_id", values[i])
			} else if value != nil {
				d.DirectContributorID = *value
			}
		case detail.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				d.GoodID = *value
			}
		case detail.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				d.OrderID = *value
			}
		case detail.FieldSelfOrder:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field self_order", values[i])
			} else if value.Valid {
				d.SelfOrder = value.Bool
			}
		case detail.FieldPaymentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value != nil {
				d.PaymentID = *value
			}
		case detail.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				d.CoinTypeID = *value
			}
		case detail.FieldPaymentCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_type_id", values[i])
			} else if value != nil {
				d.PaymentCoinTypeID = *value
			}
		case detail.FieldPaymentCoinUsdCurrency:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_usd_currency", values[i])
			} else if value != nil {
				d.PaymentCoinUsdCurrency = *value
			}
		case detail.FieldUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value.Valid {
				d.Units = uint32(value.Int64)
			}
		case detail.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				d.Amount = *value
			}
		case detail.FieldUsdAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field usd_amount", values[i])
			} else if value != nil {
				d.UsdAmount = *value
			}
		case detail.FieldCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field commission", values[i])
			} else if value != nil {
				d.Commission = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Detail.
// Note that you need to call Detail.Unwrap() before calling this method if this Detail
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Detail) Update() *DetailUpdateOne {
	return (&DetailClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Detail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Detail) Unwrap() *Detail {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Detail is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Detail) String() string {
	var builder strings.Builder
	builder.WriteString("Detail(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", d.DeletedAt))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", d.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", d.UserID))
	builder.WriteString(", direct_contributor_id=")
	builder.WriteString(fmt.Sprintf("%v", d.DirectContributorID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", d.GoodID))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", d.OrderID))
	builder.WriteString(", self_order=")
	builder.WriteString(fmt.Sprintf("%v", d.SelfOrder))
	builder.WriteString(", payment_id=")
	builder.WriteString(fmt.Sprintf("%v", d.PaymentID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", d.CoinTypeID))
	builder.WriteString(", payment_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", d.PaymentCoinTypeID))
	builder.WriteString(", payment_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", d.PaymentCoinUsdCurrency))
	builder.WriteString(", units=")
	builder.WriteString(fmt.Sprintf("%v", d.Units))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", d.Amount))
	builder.WriteString(", usd_amount=")
	builder.WriteString(fmt.Sprintf("%v", d.UsdAmount))
	builder.WriteString(", commission=")
	builder.WriteString(fmt.Sprintf("%v", d.Commission))
	builder.WriteByte(')')
	return builder.String()
}

// Details is a parsable slice of Detail.
type Details []*Detail

func (d Details) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
