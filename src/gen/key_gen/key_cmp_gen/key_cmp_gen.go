package key_cmp_gen

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/lioneagle/abnf/src/gen/key_gen"
	"github.com/lioneagle/abnf/src/keys"

	"github.com/lioneagle/goutil/src/chars"
)

type Branch struct {
	Value []byte
	Next  *TrieTree
}

func (this *Branch) Equal(ch byte) bool {
	return this.Value[0] == ch
}

func (this *Branch) EqualNoCase(ch byte) bool {
	if this.Value[0] == ch {
		return true
	}

	if !chars.IsAlpha(ch) || !chars.IsAlpha(this.Value[0]) {
		return false
	}

	return (this.Value[0] | 0x20) == (ch | 0x20)
}

const (
	TRIE_TREE_NODE_TYPE_NORNAL = iota
	TRIE_TREE_NODE_TYPE_FINAL
)

type TrieTree struct {
	NodeType int
	Key      *keys.Key
	Branches []*Branch
}

func BuildTrieTreeFromKeys(config *key_gen.Config, keys *keys.Keys) *TrieTree {
	tree := &TrieTree{}

	for _, v := range keys.Data {
		tree.AddKey(config, v)
	}
	return tree
}

func (this *TrieTree) SetFinal() {
	this.NodeType = TRIE_TREE_NODE_TYPE_FINAL
}

func (this *TrieTree) IsFinal() bool {
	return this.NodeType == TRIE_TREE_NODE_TYPE_FINAL
}

func (this *TrieTree) SetKey(key *keys.Key) {
	this.Key = key
}

func (this *TrieTree) Find(config *key_gen.Config, ch byte) *Branch {
	for _, v := range this.Branches {
		if config.CaseSensitive {
			if v.Equal(ch) {
				return v
			}
		} else {
			if v.EqualNoCase(ch) {
				return v
			}
		}
	}
	return nil
}

func (this *TrieTree) FindFinal(config *key_gen.Config) *Branch {
	return this.Find(config, 0)
}

func (this *TrieTree) HasConflict(config *key_gen.Config) bool {
	if !config.CaseSensitive {
		return false
	}

	for _, v := range this.Branches {
		if v.Value[0] == 0 {
			continue
		}

		if chars.IsAlpha(v.Value[0]) {
			continue
		}

		branch := this.Find(config, v.Value[0]&0xDF)
		if branch == nil || branch.Next != v.Next {
			return true
		}

	}
	return false
}

func (this *TrieTree) addBranch(config *key_gen.Config, dest *TrieTree, ch []byte) {
	branch := &Branch{Value: ch, Next: dest}
	this.Branches = append(this.Branches, branch)

	sort.Slice(this.Branches, func(i, j int) bool { return this.Branches[i].Value[0] < this.Branches[j].Value[0] })

	/*num := len(this.Branches)
	if this.Branches[0].Value[0] == 0 {
		num--
	}

	for i, v := range this.Branches {
		fmt.Printf("this.Branches[%d].Value = %s\n", i, string(v.Value))
	}

	if num > 1 {
		for _, v := range this.Branches {
			if len(v.Value) > 1 {
				new_tree := &TrieTree{}
				new_tree.addBranch(config, new_tree, v.Value[1:])
				v.Value = v.Value[0:1]
				v.Next = new_tree
			}
		}
	}*/

}

func (this *TrieTree) AddKey(config *key_gen.Config, key *keys.Key) {
	for i, v := range this.Branches {
		fmt.Printf("this.Branches[%d].Value = %s\n", i, string(v.Value))
	}

	tree := this
	name := key.Name
	if !config.CaseSensitive {
		name = strings.ToLower(key.Name)
	}
	for i := 0; i < len(name); i++ {
		ch := name[i:]

		var new_tree *TrieTree

		branch := tree.Find(config, ch[0])
		if branch == nil {
			new_tree = &TrieTree{}
			//tree.addBranch(config, new_tree, []byte(ch))
			tree.addBranch(config, new_tree, []byte(ch[0:1]))
			tree = new_tree
			//break
		} else {
			new_tree = branch.Next
			//new_tree.addBranch(config, new_tree, []byte(ch[0:1]))
			//new_tree.addBranch(config, new_tree, branch.Value[1:])
			//branch.Value = branch.Value[0:1]
		}

		tree = new_tree
	}

	new_tree := &TrieTree{}
	new_tree.SetFinal()
	new_tree.SetKey(key)
	tree.addBranch(config, new_tree, []byte{0})
}

type KeyCmpGenerator interface {
	GenerateFile(config *key_gen.Config, keys *keys.Keys, filename, path string)
	GenerateIndex(config *key_gen.Config, keys *keys.Keys, w io.Writer)
	GenerateActionDeclaration(config *key_gen.Config, w io.Writer)
	GenerateActionDefinition(config *key_gen.Config, tree *TrieTree, w io.Writer)
}
