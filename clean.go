// +build clean

package main

import (
	"encoding/json"
	"fmt"
)

func (item *Item) UnmarshalJSON(b []byte) error {
	tmp := []interface{}{&item.Name, &item.Quantity}
	expected := len(tmp)

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	got := len(tmp)
	if got != expected {
		return fmt.Errorf("expected length %d, got %d", expected, got)
	}

	return nil
}
