package go_bk_tree

import (
	"testing"
	l "github.com/texttheater/golang-levenshtein/levenshtein"
	"fmt"
)

type Word struct {
	word string
}

func (w Word) distanceFrom(w2 MetricTensor) distance {
	return distance(l.DistanceForStrings([]rune(w.word), []rune(w2.(Word).word), l.DefaultOptions))
}

func NewWord(w string) Word {
	return Word{
		word: w,
	}
}

func createNewTree(words []string) *BKTree {
	tree := new(BKTree)
	for w := range words {
		tree.Add(NewWord(words[w]))
	}
	return tree
}

func TestBKTreeAdd(t *testing.T) {
	wordsList := []string{"a", "ab", "abc", "d"}
	tree := createNewTree(wordsList)
	if rootVal := tree.root.MetricTensor.(Word).word; rootVal != "a" {
		t.Errorf("expected: %s, got: %s", "a", rootVal)
	}
	level1Children := tree.root.Children
	if len(level1Children) != 2 {
		t.Errorf("expected: %d, got: %d", 2, len(level1Children))
	}
	level2Children := tree.root.Children[2].Children // 'd' should be child of 'abc'
	if len(level2Children) != 1 {
		t.Errorf("expected: %d, got: %d", 1, len(level2Children))
	}
}

func ExampleBKTreeSearch() {
	wordsList := []string{"some", "soft", "sorted", "same", "mole", "soda", "salmon"}
	tree := createNewTree(wordsList)

	query := NewWord("sort")
	results := tree.Search(query, 2)
	fmt.Println(results)
	// Output: [{soft} {sorted}]
}
