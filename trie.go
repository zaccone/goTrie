package goTrie

import "unicode/utf8"

type Trie struct {
	end      bool
	letters  map[rune]*Trie
	children uint32
}

// New initializes new Trie object with attributes default values
func New() *Trie {
	return &Trie{
		end:      false,
		letters:  make(map[rune]*Trie),
		children: 0,
	}
}

// IsWord returns true if the node is also indicator for a whole word,
// false otherwise
func (t *Trie) IsWord() bool {
	if t == nil {
		return false
	}
	return t.end
}

// Add adds a string to the Trie datastructure
func (t *Trie) Add(s string) bool {
	if t == nil {
		return false
	}

	if len(s) == 0 {
		t.end = true
		return true
	}

	letter, size := utf8.DecodeRuneInString(s)

	if _, ok := t.letters[letter]; !ok {
		t.letters[letter] = New()
		t.children++
	}

	trieChild := t.letters[letter]

	return trieChild.Add(s[size:])

}

// Check checks if the string is a word stored in the Trie datastructure.
func (t *Trie) Check(s string) bool {
	if t == nil {
		return false
	}

	if len(s) == 0 {
		return t.IsWord()
	}

	letter, size := utf8.DecodeRuneInString(s)

	if childNode, ok := t.letters[letter]; !ok {
		return false
	} else {
		return childNode.Check(s[size:])
	}
}
