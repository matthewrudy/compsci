package hashtable

type row struct {
	key   string
	value interface{}
}

type HashTable struct {
	rows []*row
	size int
}

func hashString(key string) (hash int) {
	bytes := []byte(key)
	for i, c := range bytes {
		hash += (i + 1) * int(c)
	}
	return hash
}

func NewHashTable() *HashTable {
	return &HashTable{
		rows: make([]*row, 50),
		size: 0,
	}
}

func (self *HashTable) hashKey(key string) (n int) {
	return hashString(key) % 50
}

func (self *HashTable) Set(key string, value interface{}) {
	n := self.hashKey(key)
	row := self.rows[n]
	row.Set(value)
	self.size = self.size + 1
}

func (self *HashTable) Get(key string) interface{} {
	n := self.hashKey(key)
	return self.rows[n].Get(key)
}

func (self *HashTable) Size() int {
	return self.size
}

func (self *row) Get(key string) (value interface{}) {
	if self == nil {
		return nil
	} else {
		return self.value
	}
}

func (self *row) Set(value interface{}) *row {
	if self == nil {
		return &row{
			value: value,
		}
	} else {
		self.value = value
		return self
	}
}
