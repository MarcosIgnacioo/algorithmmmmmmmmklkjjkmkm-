package tries

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/arraylist"
)

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

func (tn *TriesNode) Autocompletion(word string) {
	var idx int
	curr := tn
	wChars := []rune(word)
	wLength := len(wChars)

	for i := 0; i < wLength; i++ {
		idx = GetCharCode(wChars[i])
		curr = curr.Characters[idx]
	}
	fmt.Println(word, " ... ")
	suggestions := arraylist.NewArrayList(5)
	Travel(curr, &suggestions, "")
	fmt.Println(suggestions.String())
}

func Travel(curr *TriesNode, suggestions *arraylist.ArrayList, word string) {
	alphabet := [26]string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	for j := 0; j < len(curr.Characters); j++ {
		if curr.Characters[j] != nil {
			word += alphabet[j]
			if curr.Characters[j].IsWord == true {
				suggestions.Enqueue(word)
				word = ""
			}
			Travel(curr.Characters[j], suggestions, word)
		}
	}
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
		} else {
			fmt.Println("Passed ✅")
		}
	}

	for _, word := range nonExistingWords {
		if trie.Search(word) {
			fmt.Errorf("Search passed for non-existing word: %s", word)
		} else {
			fmt.Println("Passed ✅")
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
		} else {
			fmt.Println("Passed ✅")
		}

	}

	for _, word := range remainingWords {
		if !trie.Search(word) {
			fmt.Errorf("Search failed for remaining word: %s", word)
		} else {
			fmt.Println("Passed ✅")
		}

	}
}

// La manera de crear caracteres va a ser con la formula del offset que enseño prime en el video
