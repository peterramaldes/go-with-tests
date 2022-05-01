package maps

const (
	ErrNotFound                = DictionaryErr("could not find the word you were looking for")
	ErrUpdateWordDoesNotExists = DictionaryErr("cannot update word because it does not exists")
	ErrDeleteWordDoesNotExists = DictionaryErr("cannot delete word because it does not exists")
	ErrWordExists              = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if val, ok := d[word]; ok {
		return val, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word, text string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = text
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, text string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrUpdateWordDoesNotExists
	case nil:
		d[word] = text
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrDeleteWordDoesNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
