package main

import "fmt"

func main() {

		var row int
		var col int
		fmt.Scanf("%d", &row)
		fmt.Scanf("%d", &col)
		a := make([][]int, row)
		for i := range a{
			a[i] = make([]int, col)
		}
		for i := 0; i < row; i++{
			for j := 0; j < col; j++{
				fmt.Scanf("%d", &a[i][j])
			}
		}




	//fmt.Println("Count: ", countIslands(a, row, col))

}
/*func countIslands(grid [][]int, row int, col int) (result int) {

	if len(grid) == 0 {
		return 0
	}
	result = 0
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}

	return result
}*/
