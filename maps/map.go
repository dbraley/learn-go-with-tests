package maps

// Search searches a map given key, returning the value
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
