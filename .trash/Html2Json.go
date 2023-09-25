package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

type Descriptor struct {
	name 	string
	root 	*css.Selector
	single 	map[string]*css.Selector
	many 	map[string]*css.Selector
}

func (d *Descriptor) Create(root*html.Node) *HtmlObject {
	o := &HtmlObject{
		name: d.name,
		single: make(map[string]*html.Node),
		many: make(map[string][]*html.Node),
	}
	if d.root != nil {
		roots := d.root.Select(root)
		if len(roots) == 0 {
			log.Fatal("Descriptor::Create::No root element available")
		}
		root= roots[0]
	}
	for selectorName, selector := range d.single {
		nodes := selector.Select(root)
		if len(nodes) == 0 {
			log.Fatal("Descriptor::Create::ElementNotFound")
		}
		o.single[selectorName] = nodes[0]
	}
	for selectorName, selector := range d.single {
		o.many[selectorName] = selector.Select(root)
	}
	return o
}

func UnmarshalJSONDescriptors(b []byte) ([]Descriptor, error) {
	m := make(map[string]map[string]string)
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatal("UnmarshalJSONDescriptors::",err)
	}
	res := make([]Descriptor, 0, len(m))
	for name, fields := range m {
		d := Descriptor{
			name: name,
			single: make(map[string]*css.Selector),
			many: make(map[string]*css.Selector),
		}
		for fieldName, fieldVal := range fields {
			if len(fieldName) == 0 {
				log.Fatal("UnmarshalJSONDescriptors::empty fieldName")
			}
			sel, err := css.Parse(fieldVal)
			if err != nil {
				log.Fatal("UnmarshalJSONDescriptors::", err)
			}
			if fieldName == "root" {
				d.root = sel
			} else {
				isMulti := fieldName[len(fieldName)-1] == '*'
				if isMulti {
					d.many[fieldName] = sel
				} else {
					d.single[fieldName] = sel
				}
			}
		}
		res = append(res, d)
	}
	return res, nil
}

type HtmlObject struct {
	name 	string
	single 	map[string]*html.Node 
	many 	map[string][]*html.Node
}

func MarshalJSONHtmlObjects(objs []HtmlObject) ([]byte, error) {
	res := make(map[string]any, len(objs))

	for _, o := range objs {
		d := make(map[string]any, 2 * len(o.single))
		var sb strings.Builder

		for key, node := range o.single {
			if err := html.Render(&sb, node); err != nil {
				log.Fatal("MarshalJSONHtmlObjects::", err)
			}
			d[key] = sb.String()
			sb.Reset()
		}

		for key, nodes := range o.many {
			l := make([]string, len(nodes))
			for _, node := range nodes {
				if err := html.Render(&sb, node); err != nil {
					log.Fatal("MarshalJSONHtmlObjects::", err)
				}
				l = append(l, sb.String())
				sb.Reset()
			}
			d[key] = l
		}
		res[o.name] = res
	}
	return json.Marshal(res)
}


func main() {
	jsonFilePath, htmlFilePath := os.Args[1], os.Args[2]
	buf, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal("main::JsonFileDescriptor", err)
	}
	descriptors, err := UnmarshalJSONDescriptors(buf)
	if err != nil {
		log.Fatal(err)
	}

	htmlFile, err := os.Open(htmlFilePath)
	if err != nil {
		log.Fatal("main::HtmlFileDescriptor", err)
	}
	htmlRoot, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatal("main::HtmlFileParsing", err)
	}

	htmlObjects := make([]HtmlObject, len(descriptors))
	for _, descriptor := range descriptors { 
		//fmt.Println(descriptor)
		htmlObj := descriptor.Create(htmlRoot)
		htmlObjects = append(htmlObjects, *htmlObj)
	}

	b, err := MarshalJSONHtmlObjects(htmlObjects)
	fmt.Println(b)

}
