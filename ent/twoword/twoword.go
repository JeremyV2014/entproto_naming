// Code generated by entc, DO NOT EDIT.

package twoword

import (
	"fmt"
)

const (
	// Label holds the string label denoting the twoword type in the database.
	Label = "two_word"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEnums holds the string denoting the enums field in the database.
	FieldEnums = "enums"
	// Table holds the table name of the twoword in the database.
	Table = "two_words"
)

// Columns holds all SQL columns for twoword fields.
var Columns = []string{
	FieldID,
	FieldEnums,
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

// Enums defines the type for the "enums" enum field.
type Enums string

// EnumsOption1 is the default value of the Enums enum.
const DefaultEnums = EnumsOption1

// Enums values.
const (
	EnumsOption1 Enums = "option1"
	EnumsOption2 Enums = "option2"
)

func (e Enums) String() string {
	return string(e)
}

// EnumsValidator is a validator for the "enums" field enum values. It is called by the builders before save.
func EnumsValidator(e Enums) error {
	switch e {
	case EnumsOption1, EnumsOption2:
		return nil
	default:
		return fmt.Errorf("twoword: invalid enum value for enums field: %q", e)
	}
}
