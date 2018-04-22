package __7_rotate_matrix

/* 1.7
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
*/

//is making in-place swaps! for real world problems clone the input to avoid side effects
func solution(m [][]int) {
	/* 90 degrees means the values from X to go Y
	left -> top
	top -> right
	right -> bottom
	bottom -> left
	*/

	type p struct {
		row, col int
	}

	swap := func(p1, p2 *p) {
		m[p1.row][p1.col], m[p2.row][p2.col] = m[p2.row][p2.col], m[p1.row][p1.col]
	}

	for layer := 0; layer < (len(m) / 2); layer++ {
		onionMinIndex := 0 + layer          //for top and left
		onionMaxIndex := len(m) - layer - 1 //for bottom and right

		//corners defined just for code readability
		topLeft := &p{onionMinIndex, onionMinIndex}
		topRight := &p{onionMinIndex, onionMaxIndex}
		bottomRight := &p{onionMaxIndex, onionMaxIndex}
		bottomLeft := &p{onionMaxIndex, onionMinIndex}

		//switch the entire layer
		for i := 0; i < onionMaxIndex-onionMinIndex; i++ { //we need to switch only the layerLength - 1 positions, they overlap
			left := &p{bottomLeft.row - i, bottomLeft.col}     //bottom to top
			top := &p{topLeft.row, topLeft.col + i}            //left to right
			right := &p{topRight.row + i, topRight.col}        //top to bottom
			bottom := &p{bottomRight.row, bottomRight.col - i} //right to left

			//instead of using a temporary variable, we can use the awesome
			//go feature: a,b = b,a and rotate the values
			swap(top, right)
			swap(bottom, left)
			swap(bottom, top)
		}
	}
}
