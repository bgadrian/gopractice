package main

//https://www.hackerrank.com/contests/projecteuler/challenges/euler107
//
//func main() {
//	for _, n := range input() {
//		fmt.Println(getResult(n))
//	}
//}

type edge struct {
	u, v, w int
}

//prim's algorithm
func getResult(edges []*edge) int {
	done := make(map[int]struct{})
	maxNode := 0
	//select first random node and count them
	done[edges[0].u] = struct{}{}
	for _, e := range edges {
		if e.u > maxNode {
			maxNode = e.u
		}
		if e.v > maxNode {
			maxNode = e.u
		}
	}

	//initializing the data stores
	matrix := make([][]int, maxNode-1)
	minimalMatrix := make([][]int, maxNode-1)
	for row := 0; row < maxNode; row++ {
		matrix[row] = make([]int, maxNode)
		minimalMatrix[row] = make([]int, maxNode)

		for col := 0; col < maxNode; col++ {
			matrix[row][col] = -1        //we have 0 weight edges
			minimalMatrix[row][col] = -1 //we have 0 weight edges
		}
	}

	//populating the original matrix
	for _, edge := range edges {
		//defaults are -1
		currentValue := matrix[edge.u][edge.v]
		hasValue := matrix[edge.u][edge.v] > -1
		if hasValue == false || currentValue < edge.w {
			matrix[edge.u][edge.v] = edge.w
		}
	}

	//while we haven't visited all the nodes...
	for len(done) < maxNode {
		//get all edges from all done nodes
		//get the lowest edge
		//add the edge and node to the result
	}

	//new sum of weights
	s := 0
	for _, row := range minimalMatrix {
		for _, weight := range row {
			if weight == -1 {
				continue
			}
			s += weight
		}
	}
	return s
}

//
//func input() (ints []int) {
//	var T int
//
//	_, err := fmt.Scanln(&T)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	var n int
//
//	for l := T; l > 0; l-- {
//		_, err = fmt.Scanln(&n)
//		if err != nil {
//			log.Panic(err)
//		}
//
//		ints = append(ints, n)
//	}
//
//	return
//}
