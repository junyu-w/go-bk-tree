/*
Package go-bk-tree works like a charm
 */
package go_bk_tree

import (
	"testing"
	l "github.com/texttheater/golang-levenshtein/levenshtein"
	"fmt"
	"math/rand"
	"github.com/icrowley/fake"
)


type Word struct {
	word string
}

func (w Word) DistanceFrom(w2 MetricTensor) Distance {
	return Distance(l.DistanceForStrings([]rune(w.word), []rune(w2.(Word).word), l.DefaultOptions))
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

func TestBKTree_Add(t *testing.T) {
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


// Word is a custom struct the implements the MetricTensor interface,
// and it uses the Levenshtein distance as distance function
func ExampleBKTree_Search() {
	wordsList := []string{"some", "soft", "sorted", "same", "mole", "soda", "salmon"}
	tree := createNewTree(wordsList)

	// fuzzy match
	query := NewWord("sort")
	results := tree.Search(query, 2)
	fmt.Println(results)
	// exact match
	query2 := NewWord("mole")
	results2 := tree.Search(query2, 0)
	fmt.Println(results2)
	// Output:
	// [{soft} {sorted}]
	// [{mole}]
}

func makeRandomTree(size int) ([]string, *BKTree) {
	fakeStuff := make([]string, size, size)
	for i := 0; i < size; i++ {
		fakeStuff = append(fakeStuff, fake.DomainName())
		//fakeStuff = append(fakeStuff, strconv.FormatInt(rand.Int63(), 10))
	}
	return fakeStuff, createNewTree(fakeStuff)
}

func BenchmarkBKTree_Search_ExactMatch(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	for i := 0; i < b.N; i++ {
		randWord := fakeStuff[rand.Intn(fakeSize)]
		query := NewWord(randWord)
		benchmarkTree.Search(query, 0)
	}
}


func BenchmarkBKTree_Search_Radius1Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	for i := 0; i < b.N; i++ {
		randWord := fakeStuff[rand.Intn(fakeSize)]
		query := NewWord(randWord)
		benchmarkTree.Search(query, 1)
	}
}

func BenchmarkBKTree_Search_Radius2Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	for i := 0; i < b.N; i++ {
		randWord := fakeStuff[rand.Intn(fakeSize)]
		query := NewWord(randWord)
		benchmarkTree.Search(query, 2)
	}
}

func BenchmarkBKTree_Search_Radius4Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	for i := 0; i < b.N; i++ {
		randWord := fakeStuff[rand.Intn(fakeSize)]
		query := NewWord(randWord)
		benchmarkTree.Search(query, 4)
	}
}

func BenchmarkBKTree_Search_Radius32Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	for i := 0; i < b.N; i++ {
		randWord := fakeStuff[rand.Intn(fakeSize)]
		query := NewWord(randWord)
		benchmarkTree.Search(query, 32)
	}
}
