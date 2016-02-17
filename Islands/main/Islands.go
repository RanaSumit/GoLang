package main

import "fmt"

func main() {

	var row int

	var col int

	fmt.Scanf("%d", &row)
	fmt.Scanf("%d", &col)

	a := make([][]int, row)
	for i := range a {
		a[i] = make([]int, col)
	}


	for i := 0; i < row ; i++ {
		for j := 0; j < col ; j++{
			fmt.Scanf("%d",&a[i][j])
		}
	}
	for i := 0; i < row ; i++ {
		for j := 0; j < col ; j++{
			fmt.Print(a[i][j])
		}
		fmt.Println();
	}




	fmt.Println("Count: ", countIslands(a, row, col))

}
func countIslands(grid [][]int, row int, col int) (result int) {

	if len(grid) == 0 {
		return 0
	}
	result = 0
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	var length int
	for _, h := range grid{
		length = len(h)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < length; j++ {
			if (grid[i][j] == 1 && visited[i][j] == false){
				DFS(grid, visited, i, j)
				result++
			}
		}
	}

	return result
}
func DFS(grid [][]int, visited [][]bool, i int, j int){

	var length int
	for _, h := range grid{
		length = len(h)
	}
	if (i < 0 || i >= len(grid) || j < 0 || j >= length || visited[i][j] || grid[i][j] == 0) {
		return;
	}

	visited[i][j] = true;
	DFS(grid, visited, i - 1, j);
	DFS(grid, visited, i + 1, j);
	DFS(grid, visited, i, j - 1);
	DFS(grid, visited, i, j + 1);
}