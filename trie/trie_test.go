package trie

import (
	"testing"

	"github.com/rdpitts/eratos/record"
)

func TestAddLeaf(t *testing.T) {
	trie := Trie{}

	trie.AddLeaf(123)
	if trie.Branches[123].Label != 123 {
		t.Error("Failed to add a branch to trie, or did it incorrectly")
	}
}

func TestFindBranch(t *testing.T) {
	label := uint64(123)
	trie := Trie{
		Branches: Branches{
			label: &Trie{Label: uint64(456)},
		},
	}

	if b := trie.FindBranch(label); b.Label != 456 {
		t.Error("Expected to find branch with label 456, found", b.Label)
	}
}

func TestIncrementNode(t *testing.T) {
	trie := Trie{}

	trie.IncrementNode()
	if trie.Counter != 1 {
		t.Error("Expected trie counter to be 1, was", trie.Counter)
	}
}

func TestIncrementTrie(t *testing.T) {
	trie := Trie{
		Branches: Branches{
			123: &Trie{},
			456: &Trie{},
		},
	}

	trie.IncrementTrie(record.Attribute{Label: 456})
	if trie.Counter != 1 {
		t.Error("Trie counter was not incremented to 1")
	}
	if trie.Branches[123].Counter == 1 {
		t.Error("Incorrect branch was incremented")
	}
	if trie.Branches[456].Counter != 1 {
		t.Error("Expected branch to be incremented, was not")
	}
}

func TestIsLeaf(t *testing.T) {
	var trie *Trie

	// Trie is leaf
	trie = &Trie{}
	if !trie.IsLeaf() {
		t.Error("Expected IsLeaf() to return true, was false")
	}

	// Trie is not leaf
	trie = &Trie{
		Branches: Branches{
			2:  &Trie{},
			67: &Trie{},
		},
	}
	if trie.IsLeaf() {
		t.Error("Expected IsLeaf() to return false, was true")
	}
}

func TestSearch(t *testing.T) {
	// A complete trie with four different labels: 1, 2, 3, 4
	trie := Trie{
		Branches: Branches{
			1: &Trie{
				Label: 1,
				Branches: Branches{
					2: &Trie{
						Label: 2,
						Branches: Branches{
							3: &Trie{
								Label: 3,
								Branches: Branches{
									4: &Trie{Label: 4},
								},
							},
							4: &Trie{Label: 4},
						},
					},
					3: &Trie{
						Label: 3,
						Branches: Branches{
							4: &Trie{Label: 4},
						},
					},
					4: &Trie{Label: 4},
				},
			},
			2: &Trie{
				Label: 2,
				Branches: Branches{
					3: &Trie{
						Label: 3,
						Branches: Branches{
							4: &Trie{Label: 4},
						},
					},
					4: &Trie{Label: 4},
				},
			},
			3: &Trie{
				Label: 3,
				Branches: Branches{
					4: &Trie{Label: 4},
				},
			},
			4: &Trie{Label: 4},
		},
	}
	attrs := []record.Attribute{
		record.Attribute{Label: 2},
		record.Attribute{Label: 4},
	}

	if s := trie.Search(attrs...); s.Label != 4 {
		t.Error("Failed to find trie in search")
	}
}
