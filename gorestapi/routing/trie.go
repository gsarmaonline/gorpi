package routing

import (
	"errors"
	"fmt"
	"strings"
)

const (
	IdentifierKeyword = "___identifier___"
)

type (
	Trie struct {
		rootNode *Node
	}

	Node struct {
		Name      string
		prefixMap map[string]*Node

		ParentNode *Node
		Route      *Route
	}
)

func NewTrie(rootRoute *Route) (tr *Trie) {
	rootNode := &Node{
		Name:      rootRoute.GetName(),
		prefixMap: make(map[string]*Node),
		Route:     rootRoute,
	}
	tr = &Trie{
		rootNode: rootNode,
	}
	return
}

func (tr *Trie) String() (content string) {
	content = "hello"
	tr.traverse(tr.rootNode, "", []string{})
	return
}

func (tr *Trie) traverse(node *Node, path string, routesInPath []string) (updatedRoutesInPath []string) {
	path += fmt.Sprintf("/%s", node.Name)
	if node.Route != nil {
		routesInPath = append(routesInPath, path)
	}
	for _, childNode := range node.prefixMap {
		routesInPath = tr.traverse(childNode, path, routesInPath)
	}
	updatedRoutesInPath = routesInPath
	return
}

func (tr *Trie) findLongestPrefix(spPath []string) (nearestNode *Node) {
	currNode := tr.rootNode
	for _, pathElem := range spPath {
		pathNode := currNode.findOrCreateNode(pathElem)
		currNode = pathNode
	}
	nearestNode = currNode
	return
}

func (tr *Trie) translatePathFromStr(path string) (spPath []string) {
	path = strings.Trim(path, "/")
	spPath = strings.Split(path, "/")
	return
}

func (tr *Trie) createPath(path string) (node *Node, err error) {
	spPath := tr.translatePathFromStr(path)
	node = tr.findLongestPrefix(spPath)
	return
}

func (tr *Trie) AddPath(route *Route) (node *Node, err error) {
	if node, err = tr.createPath(route.GetName()); err != nil {
		return
	}
	node.Route = route
	return
}

func (tr *Trie) GetNode(path string) (pathNode *Node, err error) {
	spPath := tr.translatePathFromStr(path)
	currNode := tr.rootNode
	for _, pathElem := range spPath {
		if pathNode, err = currNode.findNode(pathElem); err != nil {
			return
		}
		currNode = pathNode
	}
	return
}

/*
**
====================================================
Node methods
**
*/
func (node *Node) findOrCreateNode(pathElem string) (pathNode *Node) {
	var (
		err error
	)
	if pathNode, err = node.findNode(pathElem); err != nil {
		pathNode = node.CreateChildNode(pathElem)
	}
	return
}

func (node *Node) findNode(pathElem string) (pathNode *Node, err error) {
	isPresent := false
	if pathNode, isPresent = node.prefixMap[pathElem]; !isPresent {
		err = errors.New("Node not present at path: " + pathElem)
	}
	return
}

func (node *Node) CreateChildNode(pathElem string) (pathNode *Node) {
	pathNode = &Node{
		Name:       pathElem,
		prefixMap:  make(map[string]*Node),
		ParentNode: node,
	}
	node.prefixMap[pathElem] = pathNode
	return
}
