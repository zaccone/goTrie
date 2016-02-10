package goTrie

import "testing"

func TestgoTrie(t *testing.T) {
	trie := New()
	if trie == nil {
		t.Error("New failed")
	}
}

func TestWordAdding(t *testing.T) {
	trie := New()
	words := []string{
		"abrakadabra", "babajaga",
	}
	var result bool
	for _, word := range words {
		result = trie.Add(word)
		if result == false {
			t.Errorf("Adding word %s failed\n", word)
		}

		wordCheck := trie.Check(word)
		if wordCheck == false {
			t.Errorf("Couldn't find word %s\n", word)
		}

	}

}

func TestTrieErrorHandling(t *testing.T) {
	trie := New()

	trie.Add("abcd")

	r := trie.Check("abc")
	if r != false {
		t.Errorf("Got %v instead of %t", r, !r)
	}

	r = trie.Check("yyz")
	if r != false {
		t.Errorf("Got %v instead of %t", r, !r)
	}

}

func TestOperationsOnNilTries(t *testing.T) {
	var trie *Trie = nil
	var ok bool

	if ok = trie.Add("sometext"); ok {
		t.Error("Operation on nil trie wasn't catched")
	}

	if ok = trie.Check("sometext"); ok {
		t.Error("Operation on nil trie wasn't catched")
	}

	if ok = trie.IsWord(); ok {
		t.Error("Operation on nil trie wasn't catched")
	}
}
