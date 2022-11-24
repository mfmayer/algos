package graph

type LinkMap[T any] map[*Node[T]][]Link[T]

func LinkNodes[T any](bidirectional bool, linkMap LinkMap[T]) {
	if bidirectional {
		backLinkMap := LinkMap[T]{}
		for node, links := range linkMap {
			for _, link := range links {
				backLinkMap[link.Node()] = append(backLinkMap[link.Node()], Link[T]{node, link.Costs()})
			}
		}
		LinkNodes(false, backLinkMap)
	}
	for node, links := range linkMap {
		node.AddLinks(links...)
	}
}

////////////////////////

// type RouteLinks struct {
// 	*linkedlist.Element[AnyLink]
// }

// func NewRouteLink(link AnyLink) *RouteLinks {
// 	return &RouteLinks{
// 		Element: linkedlist.NewElement(link),
// 	}
// }

// func (rl *RouteLinks) TotalCosts() (costs int) {
// 	// var link linkedlist.AnyElement[AnyLink]
// 	for link := rl; link != nil; link = rl.Next().(*RouteLinks) {
// 		costs = costs + link.Data().Costs()
// 	}
// 	return
// }

// func (n *Node[T]) _Dijkstra(destination AnyNode) (route Route) {
// 	heap := heap.NewMinHeap(func(a, b *RouteLinks) bool { return a.TotalCosts() < b.TotalCosts() }, NewRouteLink(n))
// 	visited := map[AnyNode]struct{}{}
// 	for route = heap.Pop(); route != nil; route = heap.Pop() {
// 		link := route.LastLink()
// 		node := link.Node()
// 		if _, alreadyVisited := visited[node]; alreadyVisited {
// 			continue
// 		}
// 		if node == destination {
// 			return
// 		}
// 		for _, link := range node.Links() {
// 			r := append(Route{}, route...)
// 			r = append(r, link)
// 			heap.Push(r)
// 		}
// 		visited[node] = struct{}{}
// 	}
// 	return
// }
