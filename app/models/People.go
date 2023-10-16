package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type People struct {
	Id        string    `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Nickname  string    `gorm:"column:nickname;not null;unique"`
	Name      string    `gorm:"column:name;not null"`
	Birthdate string    `gorm:"column:birthdate"`
	Stack     StackList `gorm:"column:stack"`
}

type StackList []string

// Implement the Scanner interface for the StackList type
func (sl *StackList) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("Invalid stack value")
	}
	var stack []string
	if err := json.Unmarshal([]byte(str), &stack); err != nil {
		return err
	}
	*sl = StackList(stack)
	return nil
}

// Implement the Valuer interface for the StackList type
func (sl StackList) Value() (driver.Value, error) {
	stackBytes, err := json.Marshal(sl)
	if err != nil {
		return nil, err
	}
	return string(stackBytes), nil
}
