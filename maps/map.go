package maps

const (
	// ErrNotFound occurs when a word that should have been found was not
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists occurs when a word that should not exist is found
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist occurs when a word that should have been found was not
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr is an error that can occur wwhen performing operations on a Dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary type maps strings to strings
type Dictionary map[string]string

// Search searches a map given key, returning the value
func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

// Add puts a new key value pair into the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update a value in a dictionary, only if it exists
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete removes a word from a dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
