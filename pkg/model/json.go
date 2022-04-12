package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONArrayString custom type for handling string array in postgreSQL
type JSONArrayString []string

func (j JSONArrayString) Value() (driver.Value, error) {
	return strings.ReplaceAll(fmt.Sprintf("%q", j), " ", ","), nil
}

func (j *JSONArrayString) Scan(value interface{}) error {
	if value == nil {
		j = nil
		return nil
	}

	switch t := value.(type) {
	case []uint8:
		return json.Unmarshal(value.([]uint8), j)
	default:
		return fmt.Errorf("Could not scan type %T into JSONArrayString", t)
	}
}

// JSON custom type for storing db & responsing API req purpose only
type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return "null", nil
	}
	return string(j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		j = nil
		return nil
	}

	switch t := value.(type) {
	case []uint8:
		jsonData := value.([]uint8)
		if string(jsonData) == "null" {
			return nil
		}
		*j = JSON(jsonData)
		return nil
	default:
		return fmt.Errorf("could not scan type %T into json", t)
	}
}

func (j JSON) MarshalJSON() ([]byte, error) {
	switch true {
	case j == nil, len(j) == 0:
		return []byte("null"), nil
	case len(j) <= 2:
		if (j[0]) == '[' {
			return []byte("[]"), nil
		}
		return []byte("null"), nil
	default:
		return []byte(j), nil
	}
}

type JSONNullString struct {
	sql.NullString
}

func (v JSONNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JSONNullString) UnmarshalJSON(raw []byte) error {
	err := json.Unmarshal(raw, &v.NullString.String)

	v.NullString.Valid = err == nil
	return err
}

type JSONNullInt64 struct {
	sql.NullInt64
}

func (v JSONNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JSONNullInt64) UnmarshalJSON(raw []byte) error {
	err := json.Unmarshal(raw, &v.NullInt64.Int64)

	v.NullInt64.Valid = err == nil
	return err
}
