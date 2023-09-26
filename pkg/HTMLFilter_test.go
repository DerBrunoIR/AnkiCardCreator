package HTMLTrees

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

var (
	htmlSrc string = 
`<html>
	<head></head>
	<body>
		<div>
			<div class="eins">
				<p>Hello</p>
			</div>
		</div>
		<div>
			<div class="zwei">
				<p>World</p>
			</div>
		</div>
		<div>
			<div class="drei">
				<p>!</p>
			</div>
		</div>
	</body>
</html>`
)

func compareTrees(a, b *html.Node) error {
	if a == nil && b == nil {
		return nil
	} else if a == nil && b != nil {
		return fmt.Errorf("Asymetric tree: nil and %v\n", b)
	} else if a != nil && b == nil {
		return fmt.Errorf("Asymetric tree: %v and nil\n", a)
	}
	if a == b {
		return fmt.Errorf("%v == %v\n", a, b)
	}
	if a.Type != b.Type {
		return fmt.Errorf("Type unequal: %v != %v\n", a.Type, b.Type)
	}
	if a.Data != b.Data {
		return fmt.Errorf("Data uneqal: %#v != %#v\n", a.Data, b.Data)
	}
	if a.DataAtom != b.DataAtom {
		return fmt.Errorf("DataAtom unequal: %v != %v\n", a.DataAtom, b.DataAtom)
	}
	if !slices.Equal(a.Attr, b.Attr) {
		return fmt.Errorf("Attr unequal: %v != %v\n", a.Attr, b.Attr)
	}
	for ca, cb := a.FirstChild, b.FirstChild; ca != nil || cb != nil; ca, cb = ca.NextSibling, cb.NextSibling {
		err := compareTrees(ca, cb)
		if err != nil {
			return err 
		}
	}
	return nil
}

func TestDeepCopy(t *testing.T) {
	root, err := html.Parse(strings.NewReader(htmlSrc))
	if err != nil {
		t.Fatal(err)
	}
	rootCpy := DeepCopy(root)
	if err := compareTrees(root, rootCpy); err != nil {
		t.Fatal(err)
	}
}

var (
	expected_html_with_children string =
`<html>
	<head></head>
	<body>
		<div>
			<div class="eins">
				<p>Hello</p>
			</div>
		</div>
		<div>
			<div class="drei">
				<p>!</p>
			</div>
		</div>
	</body>
</html>`
	expected_html string =
`<html>
	<head></head>
	<body>
		<div>
			<div class="eins">
			</div>
		</div>
		<div>
			<div class="drei">
				<p>!</p>
			</div>
		</div>
	</body>
</html>`
)

func RemoveNewlinesAndTabs(s string) string {
	noNewlines := strings.Replace(s, "\n", "", -1)
	return strings.Replace(noNewlines, "\t", "", -1)
}

func TestDeepCopySelector(t *testing.T) {
	root, err := html.Parse(strings.NewReader(RemoveNewlinesAndTabs(htmlSrc))) 
	if err != nil { 
		t.Fatal(err)
	}
	selector, err := css.Parse(".eins, .drei *, head")
	if err != nil {
		t.Fatal(err)
	}
	rootCpy := DeepCopySelector(root, selector)
	expectedRoot, err := html.Parse(strings.NewReader(RemoveNewlinesAndTabs(expected_html)))
	if err != nil {
		t.Fatal(err)
	}
	/*
	fmt.Printf("Got: \n%v\n", HTMLString(rootCpy))
	fmt.Printf("Expected: \n%v\n", HTMLString(expectedRoot))
	*/
	if err := compareTrees(rootCpy, expectedRoot); err != nil {
		t.Fatal(err)
	}

	rootTest, err := html.Parse(strings.NewReader(RemoveNewlinesAndTabs(htmlSrc)))
	if err != nil {
		t.Fatal(err)
	}
	if err := compareTrees(root, rootTest); err != nil {
		t.Fatal(err)
	}
}

func TestDeepCopySubtree(t *testing.T) {
	root, err := html.Parse(strings.NewReader(RemoveNewlinesAndTabs(htmlSrc))) 
	if err != nil { 
		t.Fatal(err)
	}
	selector, err := css.Parse(".eins, .drei *, head")
	if err != nil {
		t.Fatal(err)
	}
	nodes := selector.Select(root)

	rootCpy := DeepCopySubtrees(root, nodes)
	
	expectedTree, err := html.Parse(strings.NewReader(RemoveNewlinesAndTabs(expected_html_with_children)))
	if err != nil {
		t.Fatal(err)
	}

	if err := compareTrees(rootCpy, expectedTree); err != nil {
		fmt.Println("Got:\n", HTMLString(rootCpy))
		fmt.Println("Expected:\n", HTMLString(expectedTree))
		t.Fatal(err)
	}

	rootTest, err := html.Parse(strings.NewReader(RemoveNewlinesAndTabs(htmlSrc)))
	if err != nil {
		t.Fatal(err)
	}
	if err := compareTrees(root, rootTest); err != nil {
		t.Fatal(err)
	}
}
