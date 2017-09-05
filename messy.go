package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (item *Item) UnmarshalJSON(b []byte) error {
	var tmp []interface{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	if got, expected := len(tmp), 2; got != expected {
		return fmt.Errorf("expected length %d, got %d", expected, got)
	}

	// now handle every field

	name, ok := tmp[0].(string)
	if !ok {
		return errors.New("Name is not a string")
	}
	item.Name = name

	quantity, ok := tmp[1].(float64)
	if !ok {
		return errors.New("Quantity is not a number")
	}
	if float64(int(quantity)) != quantity {
		return errors.New("Quantity is not an int")
	}
	item.Quantity = int(quantity)

	return nil
}
