// Package go_bk_tree is a tree data structure (implemented in Golang) specialized to index data in a metric space.
// The BK-tree data structure was proposed by Burkhard and Keller in 1973 as a solution to the problem of
// searching a set of keys to find a key which is closest to a given query key. (Doc reference: http://signal-to-noise.xyz/post/bk-tree/)
package go_bk_tree

type Distance int

// MetricTensor is an interface of data that needs to be indexed
//
// Example:
//  import l "github.com/texttheater/golang-levenshtein/levenshtein"
//
//  type Word struct {
//	  word string
//  }
//
//  func (w Word) DistanceFrom(w2 MetricTensor) Distance {
//	  return Distance(l.DistanceForStrings([]rune(w.word), []rune(w2.(Word).word), l.DefaultOptions))
//  }
type MetricTensor interface {
	DistanceFrom(other MetricTensor) Distance
}

type bkTreeNode struct {
	MetricTensor
	Children map[Distance]*bkTreeNode
}

func newbkTreeNode(v MetricTensor) *bkTreeNode {
	return &bkTreeNode{
		MetricTensor: v,
		Children: make(map[Distance]*bkTreeNode),
	}
}

type BKTree struct {
	root *bkTreeNode
}

// Add a node to BK-Tree, the location of the new node
// depends on how distance between different tensors are defined
func (tree *BKTree) Add(val MetricTensor) {
	node := newbkTreeNode(val)
	if tree.root == nil {
		tree.root = node
		return
	}
	curNode := tree.root
	for {
		dist := curNode.DistanceFrom(val)
		target := curNode.Children[dist]
		if target == nil {
			curNode.Children[dist] = node
			break
		}
		curNode = target
	}
}

func (tree *BKTree) Search(val MetricTensor, radius Distance) []MetricTensor {
	candidates := make([]*bkTreeNode, 0, 10)
	candidates = append(candidates, tree.root)
	results := make([]MetricTensor, 0, 5)
	for {
		cand := candidates[0]
		candidates = candidates[1:]
		dist := cand.DistanceFrom(val)
		if dist <= radius {
			results = append(results, cand.MetricTensor)
		}
		low, high := dist - radius, dist + radius
		for dist, child := range cand.Children {
			if dist >= low && dist <= high {
				candidates = append(candidates, child)
			}
		}
		if len(candidates) == 0 {
			break
		}
	}
	return results
}



