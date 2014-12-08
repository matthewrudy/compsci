package hashtable

type hashTableRow struct {
	key   string
	value interface{}
	next  *hashTableRow
}

func (row *hashTableRow) find(key string) *hashTableRow {
	if row.key == key {
		return row
	} else if row.next != nil {
		return row.next.find(key)
	} else {
		return nil
	}
}

func (row *hashTableRow) Get(key string) (value interface{}) {
	if row == nil {
		return nil
	}

	row = row.find(key)

	if row == nil {
		return nil
	}

	return row.value
}

func (row *hashTableRow) Set(key string, value interface{}) (newRow *hashTableRow, added bool) {
	if row == nil {
		return &hashTableRow{
			key:   key,
			value: value,
		}, true
	}

	entry := row.find(key)

	if entry == nil {
		return &hashTableRow{
			key:   key,
			value: value,
			next:  row,
		}, true
	}

	entry.value = value
	return row, false
}
