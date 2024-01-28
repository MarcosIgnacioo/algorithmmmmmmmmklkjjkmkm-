package tries

import "fmt"

type TriesNode struct {
	IsWord     bool
	Characters [26]*TriesNode
}

func NewTrie() *TriesNode {
	return &TriesNode{}
}

func (tn *TriesNode) Delete(word string) {
	var idx int
	curr := tn
	wChars := []rune(word)
	wLength := len(wChars)

	for i := 0; i < wLength; i++ {
		idx = GetCharCode(wChars[i])
		tmp := curr.Characters[idx]
		if tmp != nil {
			curr.Characters[idx] = nil
			curr = tmp
		}
	}
	curr.IsWord = false
}

func (tn *TriesNode) Insert(word string) {
	var idx int
	curr := tn
	wChars := []rune(word)
	wLength := len(wChars)

	for i := 0; i < wLength; i++ {
		idx = GetCharCode(wChars[i])
		if curr.Characters[idx] == nil {
			curr.Characters[idx] = NewTrie()
		}
		curr = curr.Characters[idx]
	}
	curr.IsWord = true
}

func (tn *TriesNode) Search(word string) bool {
	var idx int
	curr := tn
	wChars := []rune(word)
	wLength := len(wChars)

	for i := 0; i < wLength; i++ {
		idx = GetCharCode(wChars[i])
		if curr.Characters[idx] == nil {
			return false
		}
		curr = curr.Characters[idx]
	}
	return curr.IsWord
}

func GetCharCode(char rune) int {
	a := int([]rune("a")[0])
	return int(char) - a
}

func TestTrieOperations() {
	trie := NewTrie()

	// Test Insert
	wordsToInsert := []string{"apple", "app", "apricot", "banana", "bat", "batman"}
	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	// Test Search
	existingWords := []string{"apple", "app", "apricot", "banana", "bat", "batman"}
	nonExistingWords := []string{"orange", "grape", "ball"}

	for _, word := range existingWords {
		if !trie.Search(word) {
			fmt.Errorf("Search failed for existing word: %s", word)
		}
	}

	for _, word := range nonExistingWords {
		if trie.Search(word) {
			fmt.Errorf("Search passed for non-existing word: %s", word)
		}
	}

	// Test Delete
	wordsToDelete := []string{"apple", "apricot", "batman"}
	for _, word := range wordsToDelete {
		trie.Delete(word)
	}

	// Verify after deletion
	deletedWords := []string{"apple", "apricot", "batman"}
	remainingWords := []string{"app", "banana", "bat"}

	for _, word := range deletedWords {
		if trie.Search(word) {
			fmt.Errorf("Search passed for deleted word: %s", word)
		}
	}

	for _, word := range remainingWords {
		if !trie.Search(word) {
			fmt.Errorf("Search failed for remaining word: %s", word)
		}
	}
}

// La manera de crear caracteres va a ser con la formula del offset que enseño prime en el video
