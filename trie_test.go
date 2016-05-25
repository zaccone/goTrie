package goTrie

import "testing"
import "sort"

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
	}

	for _, word := range words {
		wordCheck := trie.Has(word)
		if wordCheck == false {
			t.Errorf("Couldn't find word %s\n", word)
		}
	}

}

func TestTrieErrorHandling(t *testing.T) {
	trie := New()

	trie.Add("abcd")

	r := trie.Has("abc")
	if r != false {
		t.Errorf("Got %v instead of %t", r, !r)
	}

	r = trie.Has("yyz")
	if r != false {
		t.Errorf("Got %v instead of %t", r, !r)
	}

}

func TestOperationsOnNilTries(t *testing.T) {
	var trie *Trie
	var ok bool
	const prefix = "sometext"

	if ok = trie.Add(prefix); ok {
		t.Error("Trie.Add(): Operation on nil trie wasn't catched")
	}

	if ok = trie.Has(prefix); ok {
		t.Error("Trie.Has(): Operation on nil trie wasn't catched")
	}

	if ok = trie.IsWord(); ok {
		t.Error("Trie.IsWord(): Operation on nil trie wasn't catched")
	}

	if node := trie.Get(prefix); node != nil {
		t.Error("Trie.Get(): Operation on nil trie should return nil")
	}

	if val := trie.Children(); val != uint32(0) {
		t.Errorf("Trie.Children(): Nil trie should return uint32(0), got %d instead\n", val)
	}

	if words := trie.GetWordsFromPrefix(prefix); len(words) != 0 {
		t.Errorf("Trie.GetWordsFromPrefix(): Operation on nil trie should return empty []string slice, got %d instead\n", len(words))
	}
}

func prepareTrie(expected []string, extra []string) *Trie {
	trie := New()

	for _, word := range expected {
		trie.Add(word)
	}

	for _, word := range extra {
		trie.Add(word)
	}

	return trie
}

func TestFuzzyMatching(t *testing.T) {

	expected := []string{"abcdefgh", "abcdefg",
		"abcdef", "abcde"}
	noise := []string{"azazel", "amiko", "abolicja"}

	trie := prepareTrie(expected, noise)

	const prefix = "abc"
	result := trie.GetWordsFromPrefix(prefix)

	sort.Strings(result)
	sort.Strings(expected)

	if len(result) != len(expected) {
		t.Errorf("Lengths not equal, exp: %d, got: %d\n",
			len(expected), len(result))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Errorf("Words not equal, exp: %s, got %s",
				expected[i], result[i])
		}
	}
}

func TestFuzzyMatchingNotFound(t *testing.T) {
	trie := New()
	words := trie.GetWordsFromPrefix("abrakadabra")
	if len(words) != 0 {
		t.Errorf("Trie.GetWordsFromPrefix was expected to return empty slice, got %d elements instead\n", len(words))
	}
}

func TestAddOnEmptyString(t *testing.T) {
	const word = "" // empty string
	trie := New()
	addResult := trie.Add(word)

	if addResult != true {
		t.Errorf("Trie.Add(): Expected result to be true")
	}

	if trie.IsWord() != true {
		t.Errorf("Trie.Add(): For empty string expected set root node as end of a word")
	}

}

func TestGetOnEmptyString(t *testing.T) {
	trie := New()
	node := trie.Get("")
	if node != trie {
		t.Errorf("Expected return node to be Root node")
	}
}
func TestWordCounting(t *testing.T) {

	expected := []string{"abcdefgh", "abcdefg",
		"abcdef", "abcde"}
	noise := []string{"azazel", "amiko", "abolicja"}

	trie := prepareTrie(expected, noise)

	const prefix = "abc"
	node := trie.Get(prefix)
	if node == nil {
		t.Errorf("Trie node is nil")
	}

	if node.Children() != uint32(len(expected)) {
		t.Errorf(
			"Cardinality mismatch: expected cardinality: %d vs node.childen: %d\n",
			len(expected), node.Children())
	}
}
