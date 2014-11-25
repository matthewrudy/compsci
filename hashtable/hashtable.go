package hashtable

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

func NewHashTable(tableSize int) *HashTable {
	return &HashTable{
		rows: make([]*hashTableRow, tableSize),
		size: 0,
	}
}

func (self *HashTable) hashKey(key string) (n int) {
	return hashString(key) % len(self.rows)
}

func (self *HashTable) Set(key string, value interface{}) {
	n := self.hashKey(key)
	var added bool
	self.rows[n], added = self.rows[n].Set(key, value)
	if added {
		self.size = self.size + 1
	}
}

func (self *HashTable) Get(key string) interface{} {
	n := self.hashKey(key)
	return self.rows[n].Get(key)
}

func (self *HashTable) Size() int {
	return self.size
}
