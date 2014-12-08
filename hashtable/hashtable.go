// Package hashtable implements a hash table
// with string keys, and int values
package hashtable

// HashTable is a hash table with string keys and int values
type HashTable struct {
	rows []*hashTableRow
	size int
}

// hash a given string
func hashString(key string) (hash int) {
	bytes := []byte(key)
	for i, c := range bytes {
		hash += (i + 1) * int(c)
	}
	return hash
}

// NewHashTable creates a new HashTable with a given table size
func NewHashTable(tableSize int) *HashTable {
	return &HashTable{
		rows: make([]*hashTableRow, tableSize),
		size: 0,
	}
}

func (table *HashTable) hashKey(key string) (n int) {
	return hashString(key) % table.tableSize()
}

func (table *HashTable) tableSize() (n int) {
	return len(table.rows)
}

// Set sets the value for the provided key
func (table *HashTable) Set(key string, value interface{}) {
	n := table.hashKey(key)
	var added bool
	table.rows[n], added = table.rows[n].Set(key, value)
	if added {
		table.size = table.size + 1
	}
}

// Get gets the value for the provided key
func (table *HashTable) Get(key string) interface{} {
	n := table.hashKey(key)
	return table.rows[n].Get(key)
}

// Size returns the number of keys set in the hash table
func (table *HashTable) Size() int {
	return table.size
}
