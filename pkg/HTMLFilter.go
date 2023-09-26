package HTMLTrees

import (
	"log"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)


func HTMLString(node *html.Node) string {
	var sb strings.Builder
	if err := html.Render(&sb, node); err != nil {
		log.Fatal(err)
	}
	return sb.String()
}


// Flat copy of an *html.Node, all pointers are set to nil
func Copy(node *html.Node) *html.Node {
	var attr []html.Attribute 
	if len(node.Attr) > 0 {
		attr = make([]html.Attribute, len(node.Attr))
	} else {
		attr = nil 
	}
	copy(attr, node.Attr)
	return &html.Node{
		DataAtom: node.DataAtom,
		Data: node.Data,
		Attr: attr,
		Parent: nil,
		NextSibling: nil,
		PrevSibling: nil,
		FirstChild: nil,
		LastChild: nil,
		Type: node.Type,
	}
}

// Perform a recursive copy of root, copied nodes are added to cpy. 
func rec(root *html.Node, cpy *html.Node, sel func (*html.Node) bool) {
	var prev *html.Node = nil
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if sel(c) || c.Type == html.TextNode {
			cur := Copy(c)
			cur.Parent = cpy
			if cpy.FirstChild == nil {
				cpy.FirstChild = cur
			}
			if prev != nil {
				cur.PrevSibling = prev 
				prev.NextSibling = cur
			}
			rec(c, cur, sel)
			prev = cur
		} 
	}
	cpy.LastChild = prev
}

// the returned copy contains a copy of each children 
func DeepCopy(root *html.Node) *html.Node {
	newRoot := Copy(root)
	sel := func(node *html.Node) bool {
		return true 
	}
	rec(root, newRoot, sel)
	return newRoot
}

func DeepCopyFunc(root *html.Node, sel func(*html.Node) bool) *html.Node {
	newRoot := Copy(root)
	rec(root, newRoot, sel)
	return newRoot
}

func DeepCopySelector(root *html.Node, selector *css.Selector) (*html.Node) {
	nodes := selector.Select(root)
	cache := make(map[*html.Node]bool, len(nodes))
	for _, node := range nodes {
		cache[node] = true
	}
	var lookup func(root *html.Node) bool

	lookup = func(root *html.Node) bool {
		if root == nil {
			return false 
		}
		if res, ok := cache[root]; ok {
			return res
		}

		for c := root.FirstChild; c != nil; c = c.NextSibling {
			if lookup(c) {
				cache[root] = true 
				return true
			}
		}
		cache[root] = false 
		return false
	}

	return DeepCopyFunc(root, lookup)
}

func DeepCopySubtrees(root *html.Node, subtrees []*html.Node) (*html.Node) {
	var lookup func(root *html.Node) bool
	cache := make(map[*html.Node]bool, len(subtrees))

	stack := subtrees 
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		cache[node] = true
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
	}

	lookup = func(root *html.Node) bool {
		if root == nil {
			return false 
		}
		if res, ok := cache[root]; ok {
			return res
		}

		for c := root.FirstChild; c != nil; c = c.NextSibling {
			if lookup(c) {
				cache[root] = true 
				return true
			}
		}
		cache[root] = false 
		return false
	}

	return DeepCopyFunc(root, lookup)

}

func Modify(node *html.Node, f func(*html.Node) error) error {
	if node == nil {
		return nil
	}
	f(node)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		err := Modify(c, f)
		if err != nil {
			return err 
		}
	}
	return nil
}
