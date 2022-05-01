package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "unknown"
		_, err := dictionary.Search(word)
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error, when trying to get " + word + " word, because doesn't have in our dictionary")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		word := "test"
		text := "this is just a test"

		_ = dic.Add(word, text)

		assertText(t, dic, word, text)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		text := "this is just a test"
		dic := Dictionary{word: text}

		err := dic.Add(word, text)

		assertError(t, err, ErrWordExists)
		assertText(t, dic, word, text)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		text := "this is just a test"
		newText := "new text"
		dic := Dictionary{word: text}

		err := dic.Update(word, newText)

		assertError(t, err, nil)
		assertText(t, dic, word, newText)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		text := "new definition"

		dic.Update(word, text)

		assertText(t, dic, word, text)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "foo"
		definition := "bar"
		dic := Dictionary{word: definition}
		text := "foobar"

		err := dic.Update("baz", text)

		assertError(t, err, ErrUpdateWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		dic := Dictionary{word: "test definition"}
		dic.Delete(word)

		_, err := dic.Search(word)
		if err != ErrNotFound {
			t.Errorf("expected %q to be deleted", word)
		}
	})

	t.Run("delete unknown word", func(t *testing.T) {
		word := "Test"
		dic := Dictionary{word: "test definition"}
		err := dic.Delete("foo")

		if err != ErrDeleteWordDoesNotExists {
			t.Errorf("got %q want %q", err, ErrDeleteWordDoesNotExists)
		}
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertText(t testing.TB, dic Dictionary, word string, text string) {
	t.Helper()
	got, err := dic.Search(word)

	if err != nil {
		t.Fatalf("should find added word: `err: %s`", err)
	}

	if got != text {
		t.Errorf("got %q want %q", got, text)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
