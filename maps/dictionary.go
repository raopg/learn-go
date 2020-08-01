package dictionary

//Dictionary is a  type that correspondings to a map of strings pointing to strings
type Dictionary map[string]string

//Search takes a dictionary and a search word, and returns the corrsponding value
func Search(d Dictionary, word string) string {
	return d[word]
}
