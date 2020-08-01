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

//ErrWordExists is an error thrown by Add function when the word to be added already
//exists within the dictionary
var ErrWordExists = errors.New("The word already exists in the dictionary")

//Add function adds a new word, meaning pair to the dictionary.
func (d Dictionary) Add(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = meaning
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil

}
