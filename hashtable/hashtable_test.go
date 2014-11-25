package hashtable

import "testing"

func TestMake(t *testing.T) {
	NewHashTable()
}

func TestPut(t *testing.T) {
	hash := NewHashTable()
	hash.Put("abc", 1337)
}

func TestGet(t *testing.T) {
	key := "abc"

	hash := NewHashTable()
	hash.Put(key, 1337)
	value := hash.Get(key)

	if value != 1337 {
		t.Errorf("Got %v, expected %v", value, 1337)
	}
}

func TestSize(t *testing.T) {
	key1 := "abc"
	key2 := "def"

	hash := NewHashTable()
	hash.Put(key1, 123)
	hash.Put(key2, 456)
	hash.Put(key1, 789)

	if hash.Size() != 2 {
		t.Errorf("Size %v, expected %v", hash.Size(), 2)
	}
}