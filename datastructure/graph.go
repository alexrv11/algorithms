package datastructure

type Graph struct {
	Nodes []*GraphNode
}

//GraphNode data
type GraphNode struct {
	Id int
	Edges map[int]int
}

//New creates a Graph
func New() *Graph {
	return &Graph{ Nodes: []*GraphNode{}}
} 

func (g *Graph) AddNode(id int){
	g.Nodes = append(g.Nodes, &GraphNode{
		Id: id, 
		Edges: make(map[int]int),
	})
	return
}

func (g *Graph) AddEdge(source int, destine int, weigth int) {
	g.Nodes[source].Edges[destine] = weigth
}