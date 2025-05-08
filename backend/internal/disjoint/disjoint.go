/*
 * The disjoint package implenents the disjoint set data structure.
 * author: Nicholas Sutton
 */
package disjoint

/*
 * A Node maintains data, the rank of the set it belongs to, and its parent node.
 */
type Node struct {
	Data   any
	Rank   int
	Parent *Node
}

/*
 * The MakeNode function creates a new node.
 * The parent of a new node is set to itself
 * and its rank is set to 1.
 *
 * param: the data stored in the node
 * return: a pointer to the new node
 */
func MakeNode(element any) *Node {
	n := &Node{}
	n.Data = element
	n.Rank = 1
	n.Parent = n
	return n
}

/*
 * A DisjointForest is a group of disjoint sets.
 * The DisjointForest maintains of map of all the
 * sets in its forest.
 */
type DisjointForest struct {
	Forest map[any]*Node
}

/*
 * The MakeDisjointForest function creates a new DisjointForest.
 * The Forest map is set to a new empty map
 *
 * return: a pointer to the DisjointForest
 */
func MakeDisjointForest() *DisjointForest {
	f := &DisjointForest{}
	f.Forest = make(map[any]*Node)
	return f
}

/*
 * The MakeSet function creates a new node and adds it to the
 * forest map. The map uess the data stored in the node as a key
 * and the node as the value.
 *
 * param: The data stored in the node
 * return: A pointer to the newly created node
 */
func (f *DisjointForest) MakeSet(element any) *Node {
	n := MakeNode(element)
	f.Forest[element] = n
	return n
}

/*
 * The Find function locates a node set
 * based on its element
 *
 * param: the element to find
 * return: a pointer to the node storing that element
 */
func (f *DisjointForest) Find(element any) *Node {
	return FindHelper(f.Forest[element])
}

/*
 * The FindHelper function recursively traverses the nodes in
 * a set to locate the root of the tree
 *
 * param: the current node
 * return: the parent of the current node
 */
func FindHelper(current *Node) *Node {
	if current != current.Parent {
		current.Parent = FindHelper(current.Parent)
	}

	return current.Parent
}

/*
 * The Union function joins two sets into a single set.
 * The smaller set is added to the larger set.
 * If the sets have the same rank then the first set is
 * added to the second set.
 *
 * param: The two sets to union
 */
func (f *DisjointForest) Union(setOne *Node, setTwo *Node) {
	if setOne.Rank > setTwo.Rank {
		setOne.Rank = setOne.Rank + setTwo.Rank
		setTwo.Parent = setOne
	} else {
		setTwo.Rank = setOne.Rank + setTwo.Rank
		setOne.Parent = setTwo
	}
}
