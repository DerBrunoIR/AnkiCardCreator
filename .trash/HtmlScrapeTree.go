package gostdlibintoankicards

import (
	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)


type ShadowNode struct {
	name string
	selector *css.Selector
	node *html.Node
	nextSibling, prevSibling, firstChild, lastChild *ShadowNode 
}


func NewShadowTree(node *html.Node) *ShadowNode {

}
