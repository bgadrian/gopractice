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

	get := func(position *p) int {
		return m[position.row][position.col]
	}
	set := func(position *p, newVal int) {
		m[position.row][position.col] = newVal
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

			tmp := get(bottom)
			set(bottom, get(right))
			set(right, get(top))
			set(top, get(left))
			set(left, tmp)
		}
	}
}
