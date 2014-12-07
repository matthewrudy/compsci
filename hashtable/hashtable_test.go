package hashtable

import "testing"

func TestMake(t *testing.T) {
	NewHashTable(5)
}

func TestSet(t *testing.T) {
	hash := NewHashTable(5)
	hash.Set("abc", 1337)
}

func TestGet(t *testing.T) {
	key := "abc"

	hash := NewHashTable(5)
	hash.Set(key, 1337)
	value := hash.Get(key)

	if value != 1337 {
		t.Errorf("Got %v, expected %v", value, 1337)
	}
}

func TestGetEmpty(t *testing.T) {
	key := "abc"

	hash := NewHashTable(5)
	value := hash.Get(key)

	if value != nil {
		t.Errorf("Got %v, expected %v", value, nil)
	}
}

func TestSize(t *testing.T) {
	key1 := "abc"
	key2 := "def"

	hash := NewHashTable(5)
	hash.Set(key1, 123)
	hash.Set(key2, 456)
	hash.Set(key1, 789)

	if hash.Size() != 2 {
		t.Errorf("Size %v, expected %v", hash.Size(), 2)
	}
}

func TestMultipleValues(t *testing.T) {
	key1 := "abc"
	key2 := "def"

	// table size 1, all keys collide
	hash := NewHashTable(1)
	hash.Set(key1, 123)
	hash.Set(key2, 456)
	hash.Set(key1, 789)

	if hash.Get(key1) != 789 {
		t.Errorf("Key: %v is %v, expected %v", key1, hash.Get(key1), 789)
	}

	if hash.Get(key2) != 456 {
		t.Errorf("Key: %v is %v, expected %v", key2, hash.Get(key2), 456)
	}

	if hash.Size() != 2 {
		t.Errorf("Size %v, expected %v", hash.Size(), 2)
	}
}
