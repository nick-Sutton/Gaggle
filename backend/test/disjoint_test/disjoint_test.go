package disjointtest_test

import (
	"testing"

	"github.com/nick-Sutton/Gaggle/backend/internal/disjoint"
)

func TestMakeSet(t *testing.T) {
	df := disjoint.MakeDisjointForest()

	nodeOne := df.MakeSet(1)

	if nodeOne.Data != 1 {
		t.Errorf("MakeSet result is incorrect, got: %d, expected %d", nodeOne.Data, 1)
	}

	if nodeOne.Rank != 1 {
		t.Errorf("MakeSet result is incorrect, got: %d, expected %d", nodeOne.Rank, 1)
	}

	if nodeOne.Parent != nodeOne {
		t.Errorf("MakeSet result is incorrect, got: %p, expected %p", nodeOne.Parent, nodeOne)
	}
}

func TestMakeDisjointForest(t *testing.T) {
	df := disjoint.MakeDisjointForest()

	if len(df.Forest) != 0 {
		t.Errorf("DisjointForest result is incorrect, got: %d, expected %d", len(df.Forest), 0)
	}
}

func TestFind(t *testing.T) {
	df := disjoint.MakeDisjointForest()
	nodeOne := df.MakeSet(1)
	nodeTwo := df.MakeSet(2)
	nodeThree := df.MakeSet(3)
	nodeFour := df.MakeSet(4)
	nodeFive := df.MakeSet(5)

	if df.Find(1) != nodeOne {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(1), nodeOne)
	}

	if df.Find(2) != nodeTwo {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(2), nodeTwo)
	}

	if df.Find(3) != nodeThree {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(3), nodeThree)
	}

	if df.Find(4) != nodeFour {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(4), nodeFour)
	}

	if df.Find(5) != nodeFive {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(5), nodeFive)
	}
}

func TestUnion(t *testing.T) {
	df := disjoint.MakeDisjointForest()
	nodeOne := df.MakeSet(1)
	nodeTwo := df.MakeSet(2)
	nodeThree := df.MakeSet(3)
	nodeFour := df.MakeSet(4)
	nodeFive := df.MakeSet(5)

	df.Union(nodeOne, nodeTwo)

	if df.Find(1) != nodeTwo {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(1), nodeTwo)
	}

	if nodeOne.Parent != nodeTwo {
		t.Errorf("Find result is incorrect, got: %p, expected %p", nodeOne.Parent, nodeTwo)
	}

	if df.Find(2) != nodeTwo {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(1), nodeTwo)
	}

	if nodeTwo.Rank != 2 {
		t.Errorf("Find result is incorrect, got: %d, expected %d", nodeTwo.Rank, 2)
	}

	df.Union(nodeTwo, nodeThree)

	if df.Find(3) != nodeTwo {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(3), nodeTwo)
	}

	if nodeTwo.Rank != 3 {
		t.Errorf("Find result is incorrect, got: %d, expected %d", nodeTwo.Rank, 3)
	}

	df.Union(nodeFour, nodeFive)

	if df.Find(4) != nodeFive {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(1), nodeTwo)
	}

	if nodeFour.Parent != nodeFive {
		t.Errorf("Find result is incorrect, got: %p, expected %p", nodeFour.Parent, nodeFive)
	}

	if df.Find(5) != nodeFive {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(5), nodeFive)
	}

	if nodeFive.Rank != 2 {
		t.Errorf("Find result is incorrect, got: %d, expected %d", nodeFive.Rank, 2)
	}

	df.Union(nodeTwo, nodeFive)

	if df.Find(5) != nodeTwo {
		t.Errorf("Find result is incorrect, got: %p, expected %p", df.Find(5), nodeTwo)
	}

	if nodeTwo.Rank != 5 {
		t.Errorf("Find result is incorrect, got: %d, expected %d", nodeTwo.Rank, 5)
	}
}
