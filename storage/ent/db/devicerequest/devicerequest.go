// Code generated by ent, DO NOT EDIT.

package devicerequest

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the devicerequest type in the database.
	Label = "device_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserCode holds the string denoting the user_code field in the database.
	FieldUserCode = "user_code"
	// FieldDeviceCode holds the string denoting the device_code field in the database.
	FieldDeviceCode = "device_code"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldClientSecret holds the string denoting the client_secret field in the database.
	FieldClientSecret = "client_secret"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// Table holds the table name of the devicerequest in the database.
	Table = "device_requests"
)

// Columns holds all SQL columns for devicerequest fields.
var Columns = []string{
	FieldID,
	FieldUserCode,
	FieldDeviceCode,
	FieldClientID,
	FieldClientSecret,
	FieldScopes,
	FieldExpiry,
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
	// UserCodeValidator is a validator for the "user_code" field. It is called by the builders before save.
	UserCodeValidator func(string) error
	// DeviceCodeValidator is a validator for the "device_code" field. It is called by the builders before save.
	DeviceCodeValidator func(string) error
	// ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	ClientIDValidator func(string) error
	// ClientSecretValidator is a validator for the "client_secret" field. It is called by the builders before save.
	ClientSecretValidator func(string) error
)

// OrderOption defines the ordering options for the DeviceRequest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserCode orders the results by the user_code field.
func ByUserCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserCode, opts...).ToFunc()
}

// ByDeviceCode orders the results by the device_code field.
func ByDeviceCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceCode, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByClientSecret orders the results by the client_secret field.
func ByClientSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientSecret, opts...).ToFunc()
}

// ByExpiry orders the results by the expiry field.
func ByExpiry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiry, opts...).ToFunc()
}
