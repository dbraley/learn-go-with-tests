package maps

import "errors"

// Dictionary type maps strings to strings
type Dictionary map[string]string

// Search searches a map given key, returning the value
func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return def, nil
}

// Add puts a new key value pair into the dictionary
func (d Dictionary) Add(word, definition string) error {
	// Maps are a reference type, so they can be modified without being passed by reference
	d[word] = definition
	return nil
}
