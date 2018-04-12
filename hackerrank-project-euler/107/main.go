package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler107

func main() {
	fmt.Println(getResult(input()))
}

type edge struct {
	u, v, w int
}

//prim's algorithm
func getResult(edges []*edge) int {
	var ed *edge
	done := make(map[int]struct{})
	maxNode := 1
	//select first random node and count them
	done[edges[0].u] = struct{}{}
	for _, ed = range edges {
		if ed.u > maxNode {
			maxNode = ed.u
		}
		if ed.v > maxNode {
			maxNode = ed.v
		}
	}

	var rowIndex, colIndex int

	//initializing the data stores
	matrix := make([][]int, maxNode)
	minimalMatrix := make([][]int, maxNode)
	for rowIndex = 0; rowIndex < maxNode; rowIndex++ {
		matrix[rowIndex] = make([]int, maxNode)
		minimalMatrix[rowIndex] = make([]int, maxNode)

		for colIndex = 0; colIndex < maxNode; colIndex++ {
			matrix[rowIndex][colIndex] = -1        //we have 0 weight edges
			minimalMatrix[rowIndex][colIndex] = -1 //we have 0 weight edges
		}
	}

	//populating the original matrix
	for _, ed = range edges {
		//defaults are -1
		rowIndex = ed.u - 1
		colIndex = ed.v - 1
		currentValue := matrix[rowIndex][colIndex]
		hasValue := matrix[rowIndex][colIndex] > -1
		//deal with duplicate edges, keep the smallest
		if hasValue == false || ed.w < currentValue {
			matrix[rowIndex][colIndex] = ed.w
			matrix[colIndex][rowIndex] = ed.w
		}
	}

	//while we haven't visited all the nodes...
	for len(done) < maxNode {
		//get all edges from all done nodes
		//get the lowest edge
		//add the edge and node to the result

		nextNodeA, nextNodeB, nextWeight := 0, 0, -1
		for nodeDone := range done {
			for nextIndex, weight := range matrix[nodeDone-1] {
				if weight == -1 {
					continue //no edge
				}

				nodeB := nextIndex + 1
				if _, isDone := done[nodeB]; isDone {
					continue
				}

				if nextWeight == -1 || weight < nextWeight {
					nextNodeA = nodeDone
					nextNodeB = nodeB
					nextWeight = weight
				}
			}
		}
		done[nextNodeB] = struct{}{}
		minimalMatrix[nextNodeA-1][nextNodeB-1] = nextWeight
		minimalMatrix[nextNodeB-1][nextNodeA-1] = nextWeight
	}

	//new sum of weights
	s := 0
	for rowIndex = 0; rowIndex < maxNode; rowIndex++ {
		for colIndex = 0; colIndex < rowIndex; colIndex++ {
			weight := minimalMatrix[rowIndex][colIndex]
			if weight == -1 {
				continue
			}
			s += weight
		}
	}
	return s
}

func input() (edges []*edge) {
	var maxNode, T int

	_, err := fmt.Scanf("%d %d", &maxNode, &T)
	if err != nil {
		log.Panic(err)
	}

	for l := T; l > 0; l-- {
		ed := &edge{}
		_, err = fmt.Scanf("%d %d %d", &ed.u, &ed.v, &ed.w)
		if err != nil {
			log.Panic(err)
		}

		edges = append(edges, ed)
	}

	return
}
