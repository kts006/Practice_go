package mydict

import (
	"errors"
)

// Dictionary mpa의 alias
// := map[string]string{"Name": "ABCD", "Age": "11"}
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExist = errors.New("The word already exist")
var errCantUpdate = errors.New("Can not update non-existing word")
var errCantDelete = errors.New("Can not delete non-existing word")

//Search for a value
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add add word to dict
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}

// Update the word in the dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}
	return nil
}

//Delete the word in the dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, word)
	}
	return nil
}
