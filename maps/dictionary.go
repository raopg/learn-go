package dictionary

import "errors"

//Dictionary is a  type that correspondings to a map of strings pointing to strings
type Dictionary map[string]string

//Search takes a dictionary and a search word, and returns the corrsponding value
func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", errors.New("Could not locate the word in dictionary")
	}
	return definition, nil
}
