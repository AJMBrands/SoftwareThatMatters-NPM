package main

import (
	"fmt"
	g "github.com/AJMBrands/SoftwareThatMatters/graph"
)

func main() {
	// graph G {
	//G := g.NewGraphNode(0)
	//
	//// 	e
	//e := g.NewGraphNode(1)
	//
	//// 	subgraph clusterA {
	//// 		a -- b
	//a := g.NewGraphNode(3)
	//b := g.NewGraphNode(4)
	//a.AddNeighbor(b)
	//b.AddNeighbor(a)
	//
	//// 		subgraph clusterC {
	//clusterC := g.NewGraphNode(5)
	//// 			C -- D
	//C := g.NewGraphNode(6)
	//D := g.NewGraphNode(7)
	//C.AddNeighbor(D)
	//D.AddNeighbor(C)
	//// 	}
	//
	//// 	subgraph clusterB {
	//clusterB := g.NewGraphNode(8)
	//// 		d -- f
	//d := g.NewGraphNode(9)
	//f := g.NewGraphNode(10)
	//d.AddNeighbor(f)
	//f.AddNeighbor(d)
	//// 	}
	//
	//// 	d -- D
	//d.AddNeighbor(D)
	//D.AddNeighbor(d)
	//
	//// 	e -- clusterB
	//e.AddNeighbor(clusterB)
	//clusterB.AddNeighbor(e)
	//
	//// 	clusterC -- clusterB
	//clusterC.AddNeighbor(clusterB)
	//clusterB.AddNeighbor(clusterC)
	//clusterB.AddNeighbor(G)
	//// }

	//x := g.PackageInfo{Name: "junit:junit", Versions: map[string]g.DependenciesInfo{"2020-10-11T15:19:50", map[string]string{"org.hamcrest:hamcrest-core": "1.3",
	//	"org.hamcrest:hamcrest-library": "1.3"}}}
	x := g.PackageInfo{Name: "junit:junit", Versions: map[string]g.DependenciesInfo{"3.8.1": {"2020-10-11T15:19:50", map[string]string{"org.hamcrest:hamcrest-core": "1.3",
		"org.hamcrest:hamcrest-library": "1.3"}}}}
	y := g.PackageInfo{Name: "junit:junit", Versions: map[string]g.DependenciesInfo{"3.8.2": {"2021-10-11T15:19:50", map[string]string{"org.hamcrest:hamcrest-core": "1.3",
		"org.hamcrest:hamcrest-library": "1.3"}}}}
	z := g.PackageInfo{Name: "junit:junit", Versions: map[string]g.DependenciesInfo{"3.8.3": {"2021-11-11T15:19:50", map[string]string{"org.hamcrest:hamcrest-core": "1.3",
		"org.hamcrest:hamcrest-library": "1.4"}}}}
	as := g.PackageInfo{Name: "junit:junit", Versions: map[string]g.DependenciesInfo{"4.0.2": {"2022-10-11T15:19:50", map[string]string{"org.hamcrest:hamcrest-core": "1.3",
		"org.hamcrest:hamcrest-library": "2.0"}}}}
	var myarr = []g.PackageInfo{x, y, z}
	//newgraph := simple.NewDirectedGraph()
	//a1 := g.NewGraphNode(0)
	//a2 := g.NewGraphNode(1)
	//newgraph.SetEdge(newgraph.NewEdge(a1, a2))
	//a1.AddNeighbor(a2)
	//fmt.Println(newgraph.HasEdgeBetween(a1.ID(), a2.ID()))
	m := g.CreateMap(myarr)
	g.AddElementToMap(as, m)
	g1 := g.CreateGraph(m)
	fmt.Println(g1)
	//fmt.Println(m)
}
