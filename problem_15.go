package main

import ("fmt")

const PATHS int64 = 137846528820
const MAX int = 20

var path_count = 0

func approach_goal(cur_x int, cur_y int, grid [MAX][MAX]int) {
  grid[cur_x][cur_y] = 1
  if cur_x + 1 == MAX && cur_y + 1 == MAX {
//    display_maze(grid);
    path_count++
    if path_count % 100000 == 0 {
      fmt.Printf("%d - %d = %d\r", PATHS, path_count, PATHS - int64(path_count))
    }
  }
/*  if cur_x + 1 < MAX && cur_y + 1 < MAX {
    approach_goal(cur_x + 1, cur_y + 1, grid)
  } */
  if cur_x + 1 < MAX {
    approach_goal(cur_x + 1, cur_y, grid)
  }
  if cur_y + 1 < MAX {
    approach_goal(cur_x, cur_y + 1, grid)
  }
}

func display_maze(grid [MAX][MAX]int) {
  var i, j int

  fmt.Println()
  for i = 0; i < MAX; i ++ {
    for j = 0; j < MAX; j++ {
      fmt.Printf("[%d]", grid[i][j])
    }
    fmt.Println()
  }
}


func main() {
  var grid [MAX][MAX]int

  approach_goal(0, 0, grid)
  fmt.Printf("path count: %d\n", path_count)
}
