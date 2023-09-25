package SelectorMap

import (
	"log"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)


type RootedSelectorMap struct {
	root *css.Selector
	selectors map[string]*css.Selector
}

func NewRootedSelectorMap(root string, selector_table map[string]string) (*RootedSelectorMap, error) {
	root_sel, err := css.Parse(root)
	if err != nil {
		return nil, err
	}
	parsed_selector_table := make(map[string]*css.Selector, len(selector_table))
	for k, selector_string := range selector_table {
		parsed_selector, err := css.Parse(selector_string)
		if err != nil {
			return nil, err
		}
		parsed_selector_table[k] = parsed_selector
	}
	return &RootedSelectorMap{
		root: root_sel,
		selectors: parsed_selector_table,
	}, nil
}

func (rms *RootedSelectorMap) Select(node *html.Node) []map[string][]*html.Node {
	roots := rms.root.Select(node)
	res := make([]map[string][]*html.Node, 0, len(roots))
	for _, root := range roots {
		selmap := make(map[string][]*html.Node)
		for k, sel := range rms.selectors {
			selmap[k] = sel.Select(root)
		}
		res = append(res, selmap)
	}
	return res
}

func RenderConcatenated(nodes... *html.Node) string {
	var sb strings.Builder
	for _, node := range nodes {
		err := html.Render(&sb, node)
		if err != nil {
			log.Fatal(err)
		}
		sb.Write([]byte("\n"))
	}
	return sb.String()
}
