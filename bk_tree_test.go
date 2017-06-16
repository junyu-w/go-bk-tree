/*
Package go-bk-tree works like a charm
*/
package go_bk_tree

import (
	"fmt"
	l "github.com/texttheater/golang-levenshtein/levenshtein"
	"math/rand"
	"testing"
)

type Word string

func (w Word) DistanceFrom(w2 MetricTensor) Distance {
	return Distance(l.DistanceForStrings([]rune(string(w)), []rune(string(w2.(Word))), l.DefaultOptions))
}

func createNewTreeFromWords(words []string) *BKTree {
	tree := new(BKTree)
	for w := range words {
		tree.Add(Word(words[w]))
	}
	return tree
}

func TestBKTree_Add(t *testing.T) {
	wordsList := []string{"a", "ab", "abc", "d"}
	tree := createNewTreeFromWords(wordsList)
	if rootVal := string(tree.root.MetricTensor.(Word)); rootVal != "a" {
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
	tree := createNewTreeFromWords(wordsList)

	// fuzzy match
	query := Word("sort")
	results := tree.Search(query, 2)
	fmt.Println(results)
	// exact match
	query2 := Word("mole")
	results2 := tree.Search(query2, 0)
	fmt.Println(results2)
	// Output:
	// [soft sorted]
	// [mole]
}

// ################### Benchmark tests ########################

type Number uint64

// Reference: https://github.com/agatan/bktree/blob/master/bktree_test.go
func hamming(a, b uint64) int {
	count := 0
	var k uint64 = 1
	for i := 0; i < 64; i++ {
		if a&k != b&k {
			count++
		}
		k <<= 1
	}
	return count
}

func (n Number) DistanceFrom(other MetricTensor) Distance {
	return Distance(hamming(uint64(n), uint64(other.(Number))))
}

func createNewTreeFromNumbers(nums []Number) *BKTree {
	tree := new(BKTree)
	for i := range nums {
		tree.Add(nums[i])
	}
	return tree
}

func makeRandomTree(size int) ([]Number, *BKTree) {
	fakeStuff := make([]Number, size, size)
	for i := 0; i < size; i++ {
		fakeStuff = append(fakeStuff, Number(rand.Int()))
	}
	return fakeStuff, createNewTreeFromNumbers(fakeStuff)
}

func BenchmarkBKTree_Search_ExactMatch(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randNum := fakeStuff[rand.Intn(fakeSize)]
		benchmarkTree.Search(randNum, 0)
	}
}

func BenchmarkBKTree_Search_Radius1Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randNum := fakeStuff[rand.Intn(fakeSize)]
		benchmarkTree.Search(randNum, 0)
	}
}

func BenchmarkBKTree_Search_Radius2Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randNum := fakeStuff[rand.Intn(fakeSize)]
		benchmarkTree.Search(randNum, 0)
	}
}

func BenchmarkBKTree_Search_Radius4Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randNum := fakeStuff[rand.Intn(fakeSize)]
		benchmarkTree.Search(randNum, 0)
	}
}

func BenchmarkBKTree_Search_Radius32Match(b *testing.B) {
	fakeSize := 1000000
	fakeStuff, benchmarkTree := makeRandomTree(fakeSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randNum := fakeStuff[rand.Intn(fakeSize)]
		benchmarkTree.Search(randNum, 0)
	}
}
