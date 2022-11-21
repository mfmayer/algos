package graph

import (
	"fmt"

	"github.com/mfmayer/algos"
	"github.com/mfmayer/algos/heap"
	"github.com/mfmayer/algos/ringbuffer"
	"github.com/mfmayer/algos/stack"
)

type AnyLink interface {
	fmt.Stringer
	Node() AnyNode
	Costs() int
}

func lessLinkCost(a, b *AnyLink) bool {
	return (*a).Costs() < (*b).Costs()
}

type NodeLink struct {
	node AnyNode
}

func (nl *NodeLink) Node() AnyNode {
	return nl.node
}

func (nl *NodeLink) Costs() int {
	return 0
}

func (nl *NodeLink) String() string {
	return fmt.Sprintf("%v(%v)", nl.node.GetData(), nl.Costs())
}

type NodeLinkWithCosts struct {
	NodeLink
	costs int
}

func (nl *NodeLinkWithCosts) Node() AnyNode {
	return nl.NodeLink.Node()
}

func (nl *NodeLinkWithCosts) Costs() int {
	return nl.costs
}

func (nl *NodeLinkWithCosts) String() string {
	return fmt.Sprintf("%v(%v)", nl.node.GetData(), nl.Costs())
}

type LinkToOption func(lto *linkToOption) *linkToOption

type linkToOption struct {
	nodes         []AnyNode
	bidirectional bool
	costs         int
}

func Nodes(nodes ...AnyNode) LinkToOption {
	return func(lto *linkToOption) *linkToOption {
		// Nodes() creates and returns new option
		ret := &linkToOption{
			nodes:         nodes,
			bidirectional: false,
			costs:         0,
		}
		return ret
	}
}

func Bidirectional() LinkToOption {
	return func(lto *linkToOption) *linkToOption {
		lto.bidirectional = true
		return lto
	}
}

func BidirectionalWithCosts(costs int) LinkToOption {
	return func(lto *linkToOption) *linkToOption {
		lto.bidirectional = true
		lto.costs = costs
		return lto
	}
}

func WithCosts(costs int) LinkToOption {
	return func(lto *linkToOption) *linkToOption {
		lto.costs = costs
		return lto
	}
}

type AnyNode interface {
	algos.AnyElement
	Links() []AnyLink
	LinkTo(...LinkToOption) error
	// AddLink(...AnyLink)
}

// Node in a graph
type Node[T any] struct {
	*algos.Element[T]
	links []AnyLink
}

// NewNode creates a new node for a graph
func NewNode[T any](data T, links ...AnyLink) *Node[T] {
	node := &Node[T]{
		Element: &algos.Element[T]{Data: data},
		links:   heap.HeapSort(links, lessLinkCost),
	}
	return node
}

func (n *Node[T]) Links() []AnyLink {
	return n.links
}

func (n *Node[T]) LinkTo(options ...LinkToOption) (err error) {
	ltos := map[*linkToOption]struct{}{}
	var lto *linkToOption
	for _, option := range options {
		lto = option(lto)
		ltos[lto] = struct{}{}
	}
	links := make([]AnyLink, 0, len(lto.nodes))
	for lto, _ := range ltos {
		for _, node := range lto.nodes {
			var link AnyLink
			switch lto.costs {
			case 0:
				link = &NodeLink{
					node: node,
				}
			default:
				link = &NodeLinkWithCosts{
					NodeLink: NodeLink{
						node: node,
					},
					costs: lto.costs,
				}
			}
			links = append(links, link)
			if lto.bidirectional {
				node.LinkTo(Nodes(n), WithCosts(lto.costs))
			}
		}
	}
	n.links = append(n.links, links...)
	n.links = heap.HeapSort(n.links, lessLinkCost)
	return
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v -> %v", n.GetData(), n.links)
}

// DepthFirstSearch calls visitFunc in a DFS (Depth First Search) Manner
func (n *Node[T]) DepthFirstSearch(visitFunc func(AnyNode)) {
	stack := stack.NewStack[AnyNode](n)
	visited := map[AnyNode]struct{}{}

	var next AnyNode
	for next = stack.Pop(); next != nil; next = stack.Pop() {
		if _, alreadyVisited := visited[next]; alreadyVisited {
			continue
		}
		// push linked nodes in reverse order to the stack (to have links with lowest costs lying on top)
		links := next.Links()
		for i := len(links) - 1; i >= 0; i-- {
			link := links[i]
			stack.Push(link.Node())
		}
		visitFunc(next)
		visited[next] = struct{}{}
	}
}

// BreadthFirstSearch
func (n *Node[T]) BreadthFirstSearch(visitFunc func(AnyNode)) {
	queue := ringbuffer.NewRingBuffer[AnyNode](n)
	visited := map[AnyNode]struct{}{}

	var next AnyNode
	for next = queue.Pop(); next != nil; next = queue.Pop() {
		if _, alreadyVisited := visited[next]; alreadyVisited {
			continue
		}
		for _, link := range next.Links() {
			queue.Push(link.Node())
		}
		visitFunc(next)
		visited[next] = struct{}{}
	}
}
