package dic

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errWordExists = errors.New("That word already exists")

// Search for word in dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}



// Add new word to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}