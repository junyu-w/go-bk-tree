package go_bk_tree

type distance int

type MetricTensor interface {
	distanceFrom(other MetricTensor) distance
}

type BKTreeNode struct {
	MetricTensor
	Children map[distance]*BKTreeNode
}

func NewBKTreeNode(v MetricTensor) *BKTreeNode {
	return &BKTreeNode{
		MetricTensor: v,
		Children: make(map[distance]*BKTreeNode),
	}
}

type BKTree struct {
	root *BKTreeNode
}

func (tree *BKTree) Add(val MetricTensor) {
	node := NewBKTreeNode(val)
	if tree.root == nil {
		tree.root = node
		return
	}
	curNode := tree.root
	for {
		dist := curNode.distanceFrom(val)
		target := curNode.Children[dist]
		if target == nil {
			curNode.Children[dist] = node
			break
		}
		curNode = target
	}
}

func (tree *BKTree) Search(val MetricTensor, radius distance) []MetricTensor {
	candidates := make([]*BKTreeNode, 0, 10)
	candidates = append(candidates, tree.root)
	results := make([]MetricTensor, 0, 5)
	for {
		cand := candidates[0]
		candidates = candidates[1:]
		dist := cand.distanceFrom(val)
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



