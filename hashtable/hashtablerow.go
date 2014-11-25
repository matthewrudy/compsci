package hashtable

type hashTableRow struct {
	key   string
	value interface{}
}

func (self *hashTableRow) Get(key string) (value interface{}) {
	if self == nil {
		return nil
	} else {
		return self.value
	}
}

func (self *hashTableRow) Set(key string, value interface{}) (row *hashTableRow, added bool) {
	if self == nil {
		return &hashTableRow{
			key:   key,
			value: value,
		}, true

	} else {
		if self.key != key {
			added = true
		} else {
			added = false
		}

		self.value = value

		return self, added
	}
}
