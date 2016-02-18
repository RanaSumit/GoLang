package main

func countIslands(grid [][]int) int{
	 var count, row,col int
	if len(grid) == 0 {
		return 0
	}
	count = 0
	row = len(grid)
	for _, h := range grid{
		col = len(h)
	}
	//make a 2D boolean matrix of size [row][col]
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	//traverse the matrix and look for 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (grid[i][j] == 1 && visited[i][j] == false){
				//Call DFS for checking adjacent 1 and mark them visited and finally raise the counter
				DFS(grid, visited, i, j)
				count++
			}
		}
	}

	return count
}
func DFS(grid [][]int, visited [][]bool, i int, j int){

	var length int
	for _, h := range grid{
		length = len(h)
	}
	//Check if the traversal is safe
	if (i < 0 || i >= len(grid) || j < 0 || j >= length || visited[i][j] || grid[i][j] == 0) {
		return;
	}
	//mark visited as true
	visited[i][j] = true;
	DFS(grid, visited, i - 1, j);
	DFS(grid, visited, i + 1, j);
	DFS(grid, visited, i, j - 1);
	DFS(grid, visited, i, j + 1);
}