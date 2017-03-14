package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Metadata map[string]interface{}

func (m Metadata) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	return j, err
}

func (m *Metadata) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*m, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
