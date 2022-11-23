package graph

import (
	"fmt"

	"github.com/mfmayer/algos"
	"github.com/mfmayer/algos/heap"
	"github.com/mfmayer/algos/ringbuffer"
	"github.com/mfmayer/algos/stack"
)

func lessLinkCost[T any](a, b Link[T]) bool {
	return a.Costs() < b.Costs()
}

type Link[T any] struct {
	node  *Node[T]
	costs int
}

func Links[T any](nodes ...*Node[T]) []Link[T] {
	links := make([]Link[T], len(nodes))
	for i, n := range nodes {
		links[i] = Link[T]{n, 0}
	}
	return links
}

func (l Link[T]) Node() *Node[T] {
	return l.node
}

func (l Link[T]) Costs() int {
	return l.costs
}

func (l Link[T]) String() string {
	return fmt.Sprintf("%v(%v)", l.node.Data, l.costs)
}

// Node in a graph
type Node[T any] struct {
	*algos.Element[T]
	links []Link[T]
}

// NewNode creates a new node for a graph
func NewNode[T any](data T, links ...Link[T]) *Node[T] {
	node := &Node[T]{
		Element: algos.NewElement(data),
		links:   heap.HeapSort(links, lessLinkCost[T]),
	}
	return node
}

func (n *Node[T]) Links() []Link[T] {
	return n.links
}

func (n *Node[T]) AddLinks(links ...Link[T]) {
	n.links = append(n.links, links...)
	n.links = heap.HeapSort(n.links, lessLinkCost[T])
}

func (n *Node[T]) AddLinksBidir(links ...Link[T]) {
	n.AddLinks(links...)
	for _, l := range links {
		l.node.AddLinks(Link[T]{n, l.costs})
	}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v -> %v", n.Data, n.links)
}

// DepthFirstSearch calls visitFunc in a DFS (Depth First Search) Manner
func (n *Node[T]) DepthFirstSearch(visitFunc func(*Node[T])) {
	stack := stack.NewStack(n)
	visited := map[*Node[T]]struct{}{}

	var next *Node[T]
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
func (n *Node[T]) BreadthFirstSearch(visitFunc func(*Node[T])) {
	queue := ringbuffer.NewRingBuffer(n)
	visited := map[*Node[T]]struct{}{}

	var next *Node[T]
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
