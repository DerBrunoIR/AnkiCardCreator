package HTMLTrees

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestMatchingNodes(t *testing.T) {
	root, err := html.Parse(strings.NewReader(htmlSrc))
	if err != nil {
		t.Fatal(err)
	}
	pattern, err := regexp.Compile("^World$")
	if err != nil {
		t.Fatal(err)
	}
	nodes := MatchingNodes(root, pattern)
	if len(nodes) != 1 {
		t.Fatal("expected 1 node")
	}
	node := nodes[0]
	if node.Data != "World" {
		t.Fatalf("'World' != '%s'\n", node.Data)
	}
	fmt.Printf("%+v\n", node)
}

