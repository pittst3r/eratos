package trie

import (
	"github.com/rdpitts/eratos/record"
)

type Branches []*Trie

type Trie struct {
	Label   uint64
	Counter uint32
	State   uint8
	Branches
}

func (t *Trie) AddLeaf(label uint64) *Trie {
	n := &Trie{Label: label}
	if t.Branches == nil {
		t.Branches = make([]*Trie, 1, 4)
		t.Branches[0] = n
	} else {
		t.Branches = append(t.Branches, n)
	}
	return n
}

func (t *Trie) FetchBranch(label uint64) *Trie {
	var fetched *Trie
	for _, b := range t.Branches {
		if b.Label == label {
			fetched = b
			break
		}
	}
	return fetched
}

func (t *Trie) IncrementNode() {
	t.Counter += 1
}

func (t *Trie) IncrementTrie(attrs ...record.Attribute) {
	t.IncrementNode()

	// Unless this trie is a leaf
	if !t.IsLeaf() {
		// Iterate through attributes
		for i, a := range attrs {
			// Find the current attribute in the trie's branches
			if f := t.FetchBranch(a.Label); f != nil {
				// Increment the found trie with all but current attribute
				f.IncrementTrie(attrs[i+1:]...)
			}
		}
	}
}

func (t *Trie) IsLeaf() bool {
	return len(t.Branches) == 0
}

func (t *Trie) Search(attrs ...record.Attribute) *Trie {
	for i, a := range attrs {
		if f := t.FetchBranch(a.Label); f != nil {
			t = f.Search(attrs[i+1:]...)
		}
	}
	return t
}
