package main

// Entry represents a key-value pair in the hash table
type Entry struct {
	Key   string
	Value interface{}
}

// HashTable represents the hash table
type HashTable struct {
	buckets [][]Entry
	size    int
}

// NewHashTable creates a new hash table with a specified size
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([][]Entry, size),
		size:    size,
	}
}

// hashFunction calculates the hash value for a given key
func (ht *HashTable) hashFunction(key string) int {
	hash := 0
	primeBase := 31 // distribute keys evenly

	for _, char := range key {
		hash = (hash*primeBase + int(char)) % ht.size
	}

	return hash
}

// Insert adds a new key-value pair to the hash table
func (ht *HashTable) Insert(key string, value interface{}) {
	hash := ht.hashFunction(key)

	// update value if entry exists
	for i, entry := range ht.buckets[hash] {
		if key == entry.Key {
			ht.buckets[hash][i].Value = value
			return
		}
	}

	// add if entry does not exist
	ht.buckets[hash] = append(ht.buckets[hash], Entry{Key: key, Value: value})
}

// Search finds the value associated with a given key
func (ht *HashTable) Search(key string) (interface{}, bool) {
	hash := ht.hashFunction(key)

	for _, entry := range ht.buckets[hash] {
		if key == entry.Key {
			return entry.Value, true
		}
	}

	return nil, false
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) {
	hash := ht.hashFunction(key)

	for i, entry := range ht.buckets[hash] {
		if key == entry.Key {
			ht.buckets[hash] = append(ht.buckets[hash][:i], ht.buckets[hash][i+1:]...)
			return
		}
	}
}
