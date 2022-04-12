package model

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

type BigFloat big.Float

func (b *BigFloat) Value() (driver.Value, error) {
	if b != nil {
		return (*big.Float)(b).String(), nil
	}
	return nil, nil
}

func (b *BigFloat) Scan(value interface{}) error {
	if value == nil {
		b = nil
	}

	switch t := value.(type) {
	case float64:
		(*big.Float)(b).SetFloat64(value.(float64))
	case []uint8:
		_, ok := (*big.Float)(b).SetString(string(value.([]uint8)))
		if !ok {
			return fmt.Errorf("failed to load value to []uint8: %v", value)
		}
	case string:
		_, ok := (*big.Float)(b).SetString(value.(string))
		if !ok {
			return fmt.Errorf("failed to load value to string: %v", value)
		}
	default:
		return fmt.Errorf("Could not scan type %T into BigFloat", t)
	}

	return nil
}

func (b *BigFloat) MarshalJSON() ([]byte, error) {
	return []byte((*big.Float)(b).String()), nil
}

func (b *BigFloat) UnmarshalJSON(data []byte) error {
	_, ok := (*big.Float)(b).SetString(string(data))
	if !ok {
		return fmt.Errorf("failed to load value to string: %v", data)
	}

	return nil
}
