CompSci
=======

A bunch of stuff (maybe)
playing with Go and some computer science concepts/

HashTable
---------

``` go
import "github.com/matthewrudy/compsci/hashtable"

// create a hash table with a table size of 256
hash := NewHashTable(256)

// set a value
hash.Set("key", "value")

// get a value
hash.Get("key")

// check the size
hash.Size()
```
