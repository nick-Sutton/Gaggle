package disjoint

type Node struct {
	Data   any
	Rank   int
	parent *Node
}

func MakeNode(element any) *Node {
	n := &Node{}
	n.Data = element
	n.parent = n
	return n
}

type DisjointForest struct {
	Forst map[any]*Node
}

func MakeDisjointForest() *DisjointForest {
	f := &DisjointForest{}
	f.Forst = make(map[any]*Node)
	return f
}

func (f *DisjointForest) MakeSet(element any) *Node {
	n := MakeNode(element)
	f.Forst[element] = n
	return n
}

func (f *DisjointForest) Find(element any) *Node {
	return FindHelper(f.Forst[element])
}

func FindHelper(current *Node) *Node {
	if current != current.parent {
		current.parent = FindHelper(current.parent)
	}

	return current.parent
}

func (f *DisjointForest) Union(setOne *Node, setTwo *Node) {
	if setOne.Rank > setTwo.Rank {
		setOne.Rank = setOne.Rank + setTwo.Rank
		setTwo.parent = setOne
	} else {
		setTwo.Rank = setOne.Rank + setTwo.Rank
		setOne.parent = setTwo
	}
}
