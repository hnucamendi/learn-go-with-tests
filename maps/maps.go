package maps

const (
	ErrNotFound         = DictionaryError("word was not found in dictionary")
	ErrWordExists       = DictionaryError("word already exists in the dictionary")
	ErrWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
)

type DictionaryError string

func (dE DictionaryError) Error() string {
	return string(dE)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
