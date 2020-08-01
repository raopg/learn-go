package dictionary

import "errors"

//Dictionary is a  type that correspondings to a map of strings pointing to strings
type Dictionary map[string]string

//ErrNotFound is an error thrown when a search term is not found in the dictionary
var ErrNotFound = errors.New("Could not locate the word in dictionary")

//Search takes a dictionary and a search word, and returns the corrsponding value
func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
