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

	var childNode *Trie = t

	for pos := 0; pos < len(s); {
		letter, size := utf8.DecodeRuneInString(s[pos:])
		pos += size

		if _, ok := childNode.letters[letter]; !ok {
			childNode.letters[letter] = New()
		}

		childNode, _ = childNode.letters[letter]
		// TODO(marek): Currently adding same word twice will count it twice.
		// Perhaps it's not what we really want.
		childNode.children++
	}

	childNode.end = true

	return true

}

// Get checks if the string is a word stored in the Trie datastructure.
func (t *Trie) Get(s string) *Trie {
	if t == nil {
		return nil
	}

	if len(s) == 0 {
		return t
	}

	var childNode *Trie = t
	var ok bool

	for pos := 0; pos < len(s); {
		letter, size := utf8.DecodeRuneInString(s[pos:])
		pos += size
		if childNode, ok = childNode.letters[letter]; !ok {
			return nil
		}
	}

	return childNode
}

// HasWord looks for a word and return True if the word is present, false otherwise.
func (t *Trie) Has(s string) bool {
	if t == nil {
		return false
	}

	node := t.Get(s)
	return node != nil && node.IsWord()

}

// GetwordsFromPrefix returns list of words starting with provided prefix
func (t *Trie) GetWordsFromPrefix(s string) []string {
	result := make([]string, 0, 1)
	if t == nil {
		return result
	}

	node := t.Get(s)
	if node == nil {
		return result
	}

	result = node.getWordsFromPrefix(s)
	return result
}

// getWordsFromPrefix is internally used by GetWordsFromPrefix method.
// It calls itself recursively and adds a word if the checked node is marked as end of the word
func (t *Trie) getWordsFromPrefix(prefix string) []string {
	result := make([]string, 0, 0)
	if t.IsWord() {
		result = append(result, prefix)
	}
	for k, v := range t.letters {
		subresult := v.getWordsFromPrefix(prefix + string(k))
		result = append(result, subresult...)
	}
	return result
}
