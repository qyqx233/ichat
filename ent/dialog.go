// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/qyqx233/chat-go-api/ent/dialog"
)

// Dialog is the model entity for the Dialog schema.
type Dialog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid int `json:"sid,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// Assistant holds the value of the "assistant" field.
	Assistant string `json:"assistant,omitempty"`
	// Error holds the value of the "error" field.
	Error string `json:"error,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dialog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dialog.FieldID, dialog.FieldSid:
			values[i] = new(sql.NullInt64)
		case dialog.FieldUser, dialog.FieldAssistant, dialog.FieldError:
			values[i] = new(sql.NullString)
		case dialog.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Dialog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dialog fields.
func (d *Dialog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dialog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case dialog.FieldSid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value.Valid {
				d.Sid = int(value.Int64)
			}
		case dialog.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				d.User = value.String
			}
		case dialog.FieldAssistant:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field assistant", values[i])
			} else if value.Valid {
				d.Assistant = value.String
			}
		case dialog.FieldError:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error", values[i])
			} else if value.Valid {
				d.Error = value.String
			}
		case dialog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Dialog.
// Note that you need to call Dialog.Unwrap() before calling this method if this Dialog
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dialog) Update() *DialogUpdateOne {
	return NewDialogClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dialog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dialog) Unwrap() *Dialog {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dialog is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dialog) String() string {
	var builder strings.Builder
	builder.WriteString("Dialog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("sid=")
	builder.WriteString(fmt.Sprintf("%v", d.Sid))
	builder.WriteString(", ")
	builder.WriteString("user=")
	builder.WriteString(d.User)
	builder.WriteString(", ")
	builder.WriteString("assistant=")
	builder.WriteString(d.Assistant)
	builder.WriteString(", ")
	builder.WriteString("error=")
	builder.WriteString(d.Error)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Dialogs is a parsable slice of Dialog.
type Dialogs []*Dialog
