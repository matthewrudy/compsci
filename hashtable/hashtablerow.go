package hashtable

type hashTableRow struct {
	key   string
	value interface{}
	next  *hashTableRow
}

func (self *hashTableRow) find(key string) (row *hashTableRow) {
	if self.key == key {
		return self
	} else if self.next != nil {
		return self.next.find(key)
	} else {
		return nil
	}
}

func (self *hashTableRow) Get(key string) (value interface{}) {
	if self == nil {
		return nil
	}

	row := self.find(key)

	if row != nil {
		return row.value
	} else {
		return nil
	}
}

func (self *hashTableRow) Set(key string, value interface{}) (row *hashTableRow, added bool) {
	if self == nil {
		return &hashTableRow{
			key:   key,
			value: value,
		}, true
	}

	entry := self.find(key)

	if entry == nil {
		return &hashTableRow{
			key:   key,
			value: value,
			next:  self,
		}, true

	} else {
		entry.value = value

		return self, false
	}
}
