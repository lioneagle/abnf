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

	if config.BuildSimpleTree {
		for _, v := range keys.Data {
			tree.AddKey(config, v)
		}
	} else {
		for _, v := range keys.Data {
			tree.AddKey2(config, v)
		}
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
}

func (this *TrieTree) NonFinalBranchNum() int {
	num := len(this.Branches)
	if num == 0 {
		return 0
	}

	if this.Branches[0].Value[0] == 0 {
		num--
	}
	return num
}

func (this *TrieTree) FirstNonFinalBranch() *Branch {
	if len(this.Branches) == 0 {
		return nil
	}

	if this.Branches[0].Value[0] != 0 {
		return this.Branches[0]
	}

	if len(this.Branches) <= 1 {
		return nil
	}
	return this.Branches[1]
}

func (this *TrieTree) AddKey(config *key_gen.Config, key *keys.Key) {
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
			tree.addBranch(config, new_tree, []byte(ch[0:1]))
		} else {
			new_tree = branch.Next
		}

		tree = new_tree
	}

	new_tree := &TrieTree{}
	new_tree.SetFinal()
	new_tree.SetKey(key)
	tree.addBranch(config, new_tree, []byte{0})
}

func (this *TrieTree) AddKey2(config *key_gen.Config, key *keys.Key) {
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

			if tree.NonFinalBranchNum() > 0 {
				firstNonFinalBranch := tree.FirstNonFinalBranch()
				if len(firstNonFinalBranch.Value) > 1 {
					new_tree1 := &TrieTree{}
					new_tree1.addBranch(config, firstNonFinalBranch.Next, firstNonFinalBranch.Value[1:])
					firstNonFinalBranch.Next = new_tree1
					firstNonFinalBranch.Value = firstNonFinalBranch.Value[0:1]
				}

				if len(ch) > 1 {
					tree.addBranch(config, new_tree, []byte(ch[0:1]))
					ch = ch[1:]
					tree = new_tree
					new_tree = &TrieTree{}

				}
			}

			tree.addBranch(config, new_tree, []byte(ch))
			tree = new_tree
			break

		} else {
			if len(branch.Value) > 1 {
				new_tree1 := &TrieTree{}
				new_tree1.addBranch(config, branch.Next, branch.Value[1:])
				branch.Next = new_tree1
				branch.Value = branch.Value[0:1]
			}

			tree = branch.Next
		}
	}

	new_tree2 := &TrieTree{}
	new_tree2.SetFinal()
	new_tree2.SetKey(key)

	tree.addBranch(config, new_tree2, []byte{0})
}

func PrintIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
}

func (this *TrieTree) Print(depth int) {
	/*if depth == 0 {
		fmt.Println("---------------------------------------")
	}*/
	PrintIndent(depth * 4)
	fmt.Printf("depth = %d, this = %p\n", depth, this)
	if this.IsFinal() {
		PrintIndent(depth * 4)
		fmt.Println("Final: key =", this.Key.Name)
		return
	}

	for i, v := range this.Branches {
		PrintIndent(depth * 4)
		if v.Value[0] == 0 {
			fmt.Printf("[%d]: Value = 0x00, Next = %p\n", i, v.Next)
		} else {
			fmt.Printf("[%d]: Value = \"%s\", Next = %p\n", i, string(v.Value), v.Next)
		}
		v.Next.Print(depth + 1)
	}
	if depth == 0 {
		fmt.Println("---------------------------------------")
	}
}

func (this *TrieTree) PrintBranches(depth int) {
	PrintIndent(depth * 4)
	fmt.Printf("depth = %d, this = %p\n", depth, this)
	for i, v := range this.Branches {
		PrintIndent(depth * 4)
		if v.Value[0] == 0 {
			fmt.Printf("[%d]: Value = 0x00, Next = %p\n", i, v.Next)
		} else {
			fmt.Printf("[%d]: Value = \"%s\", Next = %p\n", i, string(v.Value), v.Next)
		}
	}
	fmt.Println("---------------------------------------")
}

type KeyCmpGenerator interface {
	GenerateFile(config *key_gen.Config, keys *keys.Keys, filename, path string)
	GenerateIndex(config *key_gen.Config, keys *keys.Keys, w io.Writer)
	GenerateActionDeclaration(config *key_gen.Config, w io.Writer)
	GenerateActionDefinition(config *key_gen.Config, tree *TrieTree, w io.Writer)
}
