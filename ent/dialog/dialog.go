// Code generated by ent, DO NOT EDIT.

package dialog

import (
	"time"
)

const (
	// Label holds the string label denoting the dialog type in the database.
	Label = "dialog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldAssistant holds the string denoting the assistant field in the database.
	FieldAssistant = "assistant"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the dialog in the database.
	Table = "dialogs"
)

// Columns holds all SQL columns for dialog fields.
var Columns = []string{
	FieldID,
	FieldSid,
	FieldUser,
	FieldAssistant,
	FieldError,
	FieldCreatedAt,
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

var (
	// DefaultUser holds the default value on creation for the "user" field.
	DefaultUser string
	// DefaultAssistant holds the default value on creation for the "assistant" field.
	DefaultAssistant string
	// DefaultError holds the default value on creation for the "error" field.
	DefaultError string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
