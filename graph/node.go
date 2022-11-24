package graph

import (
	"fmt"

	"github.com/mfmayer/algos"
	"github.com/mfmayer/algos/heap"
	"github.com/mfmayer/algos/ringbuffer"
	"github.com/mfmayer/algos/stack"
)

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
		links := next.Links()
		for i := len(links) - 1; i >= 0; i-- {
			link := links[i]
			queue.Push(link.Node())
		}
		visitFunc(next)
		visited[next] = struct{}{}
	}
}

// Dijkstra algorithm to find shortest route from this node to destination
func (n *Node[T]) Dijkstra(destination *Node[T]) (route Route[T]) {
	return n.AStar(destination, nil)
}

// AStar (A*) algorithm to find shortest route from this node to destination using a heuristic function to search more targeted
func (n *Node[T]) AStar(destination *Node[T], heuristic func(node *Node[T]) int) (route Route[T]) {

	heap := heap.NewMinHeap(func(a, b Route[T]) bool {
		aCosts := a.TotalCosts()
		bCosts := b.TotalCosts()
		if heuristic != nil {
			aCosts = aCosts + heuristic(a.LastNode())
			bCosts = bCosts + heuristic(b.LastNode())
		}
		return aCosts < bCosts
	}, Route[T]{Link[T]{n, 0}})

	visited := map[*Node[T]]struct{}{}
	for route = heap.Pop(); route != nil; route = heap.Pop() {
		link := route.LastLink()
		node := link.Node()
		if _, alreadyVisited := visited[node]; alreadyVisited {
			continue
		}
		if node == destination {
			return
		}
		for _, link := range node.Links() {
			r := append(Route[T]{}, route...)
			r = append(r, link)
			heap.Push(r)
		}
		visited[node] = struct{}{}
	}
	return
}
