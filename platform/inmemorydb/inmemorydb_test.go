package inmemorydb

import (
	"testing"
)

func TestAdd(t *testing.T) {
	memory := New()
	memory.Add(Item{})
	if len(memory.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	memory := New()
	memory.Add(Item{})
	results := memory.GetAll()
	println(results)
	if len(results) != 1 {
		t.Errorf("can not get All Items")
	}
}
