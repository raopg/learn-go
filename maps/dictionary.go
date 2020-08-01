package dictionary

//Dictionary is a  type that correspondings to a map of strings pointing to strings
type Dictionary map[string]string

//DictErr type defines the error message types for dictionary methods
type DictErr string

//Error function implements the error interface
func (e DictErr) Error() string {
	return string(e)
}

const (
	//ErrNotFound is an error thrown when a search term is not found in the dictionary
	ErrNotFound = DictErr("Could not locate the word in dictionary")
	//ErrWordExists is an error thrown by Add function when the word to be added already
	//exists within the dictionary
	ErrWordExists = DictErr("The word already exists in the dictionary")
	//ErrWordDoesNotExist is an error thrown by Update function when the caller tries to update
	// a word that does not exist within the dictionary.
	ErrWordDoesNotExist = DictErr("The word does not exist in the dictionary")
)

//Search takes a dictionary and a search word, and returns the corrsponding value
func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

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

//Update function takes a word and its new meaning and updates the dictionary with new meaning
func (d Dictionary) Update(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = meaning
		return nil
	default:
		return err
	}
}
