package graph

import (
	"fmt"

	"github.com/mfmayer/algos"
	"github.com/mfmayer/algos/linkedlist"
	"github.com/mfmayer/algos/ringbuffer"
	"github.com/mfmayer/algos/stack"
)

// AnyLink interface for any link in a graph
type AnyLink interface {
	linkedlist.AnyElement
	To() AnyNode
}

// AnyNode interface for any node in a graph
type AnyNode interface {
	algos.AnyElement
	Links() AnyLink
	AddLinkTo(AnyNode, ...func(*linkOptions)) AnyNode
	AddLinksTo(nodes, ...func(*linkOptions)) AnyNode
}

// Link in a graph
type Link struct {
	*linkedlist.Element
	to AnyNode
}

// NewLink creates a new link to a node with any value
func NewLink(to AnyNode, value algos.Any) *Link {
	link := &Link{
		Element: linkedlist.NewElement(value).(*linkedlist.Element),
		to:      to,
	}
	return link
}

func (l *Link) String() string {
	var next AnyLink = l
	var s string
	for next != nil {
		linksTo := next.To()
		s = s + fmt.Sprintf("%v,", linksTo.Value())
		if next.Next() == nil {
			break
		}
		next = next.Next().(AnyLink)
	}
	return s
}

// To returns linked node
func (l *Link) To() AnyNode {
	return l.to
}

// Node in a graph
type Node struct {
	*algos.Element
	links AnyLink
}

// NewNode creates a new node for a graph
func NewNode(value algos.Any, linkedNodes ...AnyNode) *Node {
	node := &Node{
		Element: algos.NewElement(value),
	}
	node.AddLinksTo(Nodes(linkedNodes...))
	return node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v -> %v", n.Value(), n.links)
}

// Links returns head of node's link list
func (n *Node) Links() AnyLink {
	return n.links
}

type linkOptions struct {
	bidirectional bool
	linkValue     algos.Any
}

type nodes func() []AnyNode

// Nodes to define multiple nodes e.g. to which a link shall be created from another node (see Node's AddLinksTo)
func Nodes(nodes ...AnyNode) nodes {
	return func() []AnyNode {
		return nodes
	}
}

// WithLinkValue option to add any value to a link
func WithLinkValue(value algos.Any) func(*linkOptions) {
	return func(lo *linkOptions) {
		lo.linkValue = value
	}
}

// Bidirectional option to create a link bidirectionally
var Bidirectional func(*linkOptions) = func(lo *linkOptions) {
	lo.bidirectional = true
}

// AddLinkTo adds a link to another node
func (n *Node) AddLinkTo(other AnyNode, options ...func(*linkOptions)) AnyNode {
	lo := linkOptions{}
	for _, opt := range options {
		opt(&lo)
	}
	link := NewLink(other, lo.linkValue)
	if n.links == nil {
		n.links = link
	} else {
		n.links.Append(link)
	}
	if lo.bidirectional {
		other.AddLinkTo(n, WithLinkValue(lo.linkValue))
	}
	return n
}

// AddLinksTo adds a link to multiple other nodes
func (n *Node) AddLinksTo(others nodes, options ...func(*linkOptions)) AnyNode {
	otherNodes := others()
	for _, other := range otherNodes {
		n.AddLinkTo(other, options...)
	}
	return n
}

// DepthFirstSearch calls visitFunc in a DFS (Depth First Search) Manner
func (n *Node) DepthFirstSearch(visitFunc func(AnyNode)) {
	stack := stack.NewStack(n)
	visited := map[AnyNode]struct{}{}

	var next AnyNode
	for next, _ = stack.Pop().(AnyNode); next != nil; next, _ = stack.Pop().(AnyNode) {
		if _, alreadyVisited := visited[next]; alreadyVisited {
			continue
		}
		for link := next.Links(); link != nil; link, _ = link.Next().(AnyLink) {
			stack.Push(link.To())
		}
		visitFunc(next)
		visited[next] = struct{}{}
	}
}

// BreadthFirstSearch
func (n *Node) BreadthFirstSearch(visitFunc func(AnyNode)) {
	queue := ringbuffer.NewRingBuffer(ringbuffer.WithInitialValues(n))
	visited := map[AnyNode]struct{}{}

	var next AnyNode
	for next, _ = queue.Pop().(AnyNode); next != nil; next, _ = queue.Pop().(AnyNode) {
		if _, alreadyVisited := visited[next]; alreadyVisited {
			continue
		}
		for link := next.Links(); link != nil; link, _ = link.Next().(AnyLink) {
			queue.Push(link.To())
		}
		visitFunc(next)
		visited[next] = struct{}{}
	}
}
