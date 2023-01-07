package tokens

import (
	"fmt"
	"strconv"
)

type IdentTable struct {
	table  map[string]string
	lastID int
}

func NewIdentTable() *IdentTable {
	return &IdentTable{
		table:  make(map[string]string),
		lastID: 0,
	}
}

func (it *IdentTable) GetID(ident string) string {
	id, ok := it.table[ident]
	if ok {
		return id
	}

	idStr := strconv.Itoa(it.lastID)
	it.table[ident] = idStr

	it.lastID++

	return idStr
}

func (it *IdentTable) Print() {
	for value, id := range it.table {
		fmt.Printf("id: %s, value: %s\n", id, value)
	}
}
