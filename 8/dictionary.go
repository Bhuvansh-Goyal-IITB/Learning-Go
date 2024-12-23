package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrorNotFound         = DictionaryErr("could not find what you were looking for")
	ErrorWordExists       = DictionaryErr("the word you are trying to add already exists")
	ErrorWordDoesNotExist = DictionaryErr("the word you are trying to update does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	switch _, err := d.Search(word); err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	switch _, err := d.Search(word); err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	switch _, err := d.Search(word); err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
