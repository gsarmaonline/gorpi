package routing

import (
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
		Name:      rootRoute.RequestURI,
		prefixMap: make(map[string]*Node),
		Route:     rootRoute,
	}
	tr = &Trie{
		rootNode: rootNode,
	}
	return
}

func (tr *Trie) findLongestPrefix(spPath []string) (nearestNode *Node) {
	currNode := tr.rootNode
	for _, pathElem := range spPath {
		pathNode := currNode.findOrCreateNode(pathElem)
		currNode = pathNode
	}
	return
}

func (tr *Trie) createPath(path string) (node *Node, err error) {
	spPath := strings.Split(path, "/")
	node = tr.findLongestPrefix(spPath)
	return
}

func (tr *Trie) AddPath(route *Route) (node *Node, err error) {
	if node, err = tr.createPath(route.RequestURI); err != nil {
		return
	}
	node.Route = route
	return
}

func (node *Node) findOrCreateNode(pathElem string) (pathNode *Node) {
	isPresent := false
	if pathNode, isPresent = node.prefixMap[pathElem]; !isPresent {
		pathNode = node.CreateChildNode(pathElem)
	}
	return
}

func (node *Node) CreateChildNode(pathElem string) (pathNode *Node) {
	// TODO: pathElem should be checked for the presence of
	// identifiers
	pathNode = &Node{
		Name:       pathElem,
		prefixMap:  make(map[string]*Node),
		ParentNode: node,
	}
	node.prefixMap[pathElem] = pathNode
	return
}
