package main

func main() {

}

// 图 中等题 岛屿数量
// 方法1 深度优先搜索
func numIslands(grid [][]byte) int {
	// 特殊情况处理
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	// 返回值 记录岛屿数量
	resultNum := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				resultNum++
				dfs(grid, i, j, row, col)
			}
		}
	}
	return resultNum
}

// 深度优先遍历函数 关键点 在深度优先搜索的过程中，每个搜索到的 1 都会被重新标记为 0
func dfs(grid [][]byte, a, b, row, col int) {
	// 下标溢出就退出到上一层递归
	if a < 0 || b < 0 || a == row || b == col || grid[a][b] == '0' { //关键点 grid[a][b] == '0'
		return
	}
	//遍历该节点的上下左右
	//关键点 在深度优先搜索的过程中，每个搜索到的 1 都会被重新标记为 0
	grid[a][b] = '0'
	dfs(grid, a-1, b, row, col)
	dfs(grid, a+1, b, row, col)
	dfs(grid, a, b-1, row, col)
	dfs(grid, a, b+1, row, col)
}
