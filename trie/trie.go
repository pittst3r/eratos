package trie

import (
	"github.com/rdpitts/eratos/record"
)

type Branches map[uint64]*Trie

type Trie struct {
	Label   uint64
	Counter uint32
	State   uint8
	Branches
}

func (t *Trie) AddLeaf(label uint64) *Trie {
	if t.Branches == nil {
		t.Branches = make(map[uint64]*Trie)
	}
	t.Branches[label] = &Trie{Label: label}
	return t.Branches[label]
}

func (t *Trie) FindBranch(label uint64) *Trie {
	return t.Branches[label]
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
			if f := t.FindBranch(a.Label); f != nil {
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
		if f := t.FindBranch(a.Label); f != nil {
			t = f.Search(attrs[i+1:]...)
		}
	}
	return t
}
