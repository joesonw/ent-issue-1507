package player_pb

import (
	"database/sql/driver"
	"fmt"
)

type Info struct {
}

type Status string

const (
	Status_OK     Status = "OK"
	Status_NOT_OK Status = "NOT_OK"
)

func (s Status) Values() []string {
	return []string{"OK", "NOT_OK"}
}

func (s *Status) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*s = Status(string(v))
	case string:
		*s = Status(v)
	default:
		return fmt.Errorf("%v is not ok", src)
	}
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return string(s), nil
}
