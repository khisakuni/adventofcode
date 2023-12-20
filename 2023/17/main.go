package main

// Credit to https://topaz.github.io/paste/#XQAAAQDeDQAAAAAAAAA4GEiZzRd1JAgz+whYRQxSFI7XvmlfhtGDiniujRhadVR1sGwdDMRmMnTvGTTnUsqSF87cjLYSuKArccOM5JNNFUtIqckouw1flfajW9jALD0zQpcy28s3Hv/Fr0E5H1cWDs0q3mlLjo8IBGB+uf7pwtQmcuGz/CnqsKai6NlFIztcgR6IvOhEcg42MH4mAUn4HkZ7cHdDWpxmo5RN0Fl1YlRxG1OMLt60zqp4VtQl+RoIzQgW+ooAMSibO9/nmFnuzgcVxnEoeHxkoZzcW8av5jEpzgy3jF7sJMEbwz3xjjeUknrHG4rLMfwQ53Wr+P2teVjNN2WRl/whNVKzrOUQFPwzh/ViuHDDf1ZcR7n79ROMoXpcHhhH0BZ3CGlm655mmJ+4qoznt6y2xqQ7vh/+9sI7RjQU5QBTWn6JxhweKJW0dfT5CsKSAY2ZzySlC1HK6zVS4DP9mBljKSN2ZA7emZYIbHZIJEXfMkhf+z/GRPyHeytaRUODFI4XigeaVJAYeRj6eaRwEEK8iu2X6LoY31DwKEEqvypdIhT5aLp3iQ2NzXfxIH+Q5qPN3/Wd87/3y6XVds7PzQkcsPoH7e6Gx41dkuFf1z1/CDlU5PI1L9PjNYrpDyGzG0E0YKSDuTW3X0FRnA/3oNxg+Uu9grpoAs422pT6lOEd9uL2KA3fm2xRX+ma6ocBwUGRai6ct2I018IZdvD4ng9UctjDkZVD1IDdP8yKHXPK/TUsF0YcoeUp/rblTIlYLv804wYpAapKSEeeemGE0Za5Ok7HBx4b9stUFrvjx2NVvWDfAUkjqvUUWHj0Cueg/NVbj5QmzvZE5EfjpiHC5/SPKL2V8Wzb3XricH6slyMMUgEYIaXoIX6DzO2v/8InH4BwS1iHNINtyvV39bKn0BAiz2RA8MfAGUKQCFp2I3bMe9sMp7qQjaO2nLx4iqUGt+JztlEci/z57eHigObx7XhWcrvyYtZNjKGLOuUJB2GWCuQRdOYBZikEQq4n4pLNSuExTBmXtEhUIHVpMU+kQh9IDQ6Xcd4Y31B0WZvA/5eHysr5Z1NTrzYkkJCKj6+60SK9yk8dxXmcZxcIOTgnLfH/DzQZGGgCrreTdOH1O8ZqMIZoqe4xrdZYH6LlJZtS6vBtG4gT37i+JnvEsM07DnkSOhd/HMyM2ea6/LSbjTAQ7rKur7LzAEOKefU/2JxIP1lyCvx3LT2xONmQmtILL/vxWKvCsfUH7L7jERm+2RMhKSy65l90y9qQjFNHCmXq5nClPCYxbKUuJ1/KyhDPelJWJUVSJcmTGS8cYttkc+ybxmkDsta2BlSz0/swY7AKlRWd8Tr7jbOK/ye4ubDPWKmZDU6VwzR8CsPT8eRk8kQ0mQYjUY890Ist+YMqSG8s0wgxq5n6DEPc7ZcXc8Ieje3oU1Bqg0xjBbZulIcpNvhUJsWidUslL2aEfgdGqpAJDRXJSO4ZLcHmoIZKZXr1CDDm5RJGlKhgTkKv5xRzpsQWMwij7SPEGvDAdS/MQDczDexFKVzRr+KlH0O8LBGtUUsT+MbnQCsZvRxHzzkUScRZbzPzVDvDZxPpKPvUpGkdY4wu2PC9tFBuduj0wCEC+Pl7wQkXNOq1gJF6sBIR28OO20W7HXU3WYcp6x8zbPJ9ILY0ZQbbW//dkEaYUqhQpqwT/Pxw18ER+jFhrTnSNkMD/ozosEdOn5QB+63Qoo+1aBiv/ZAPiBWz8T10E5jSTp+w6bd6WqSUQ+0cELNE9TnmiVWIZzZ/xHKU//44vAY=

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type direction struct {
	row, col int
}

var (
	left  = direction{row: 0, col: -1}
	right = direction{row: 0, col: 1}
	up    = direction{row: -1, col: 0}
	down  = direction{row: 1, col: 0}
)

var rotations = map[direction][]direction{
	left:  {up, down},
	right: {up, down},
	up:    {left, right},
	down:  {left, right},
}

var reverse = map[direction]direction{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type state struct {
	row   int
	col   int
	dir   direction
	moves int
}

type item struct {
	heatLoss int
	state    state
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(b)

	data := strings.Split(input, "\n")
	grid := data[:len(data)-1]
	out := dijkstra(grid, 10)

	fmt.Println("total:", out)
}

func dijkstra(grid []string, maxConsecutive int) int {
	m := len(grid)
	n := len(grid[0])
	startRight := state{row: 0, col: 0, dir: right, moves: 0}
	startDown := state{row: 0, col: 0, dir: down, moves: 0}
	pq := priorityQueue{
		&item{heatLoss: 0, state: startRight, index: 0},
		&item{heatLoss: 0, state: startDown, index: 1},
	}

	minCost := map[state]int{startRight: 0, startDown: 0}
	heap.Init(&pq)

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*item)
		if minCost[curr.state] < curr.heatLoss {
			continue
		}

		if curr.state.row == m-1 && curr.state.col == n-1 && curr.state.moves >= 4 {
			return curr.heatLoss
		}

		for _, dir := range [4]direction{left, right, up, down} {
			isReverse := dir == reverse[curr.state.dir]
			isRotated := slices.Contains(rotations[curr.state.dir], dir)

			if curr.state.moves == maxConsecutive && !isRotated || isReverse {
				continue
			}

			ni, nj := curr.state.row+dir.row, curr.state.col+dir.col
			nextMoves := curr.state.moves

			if curr.state.moves < 4 {
				if dir != curr.state.dir {
					continue
				}
				nextMoves += 1
			} else {
				if dir != curr.state.dir {
					nextMoves = 1
				} else {
					nextMoves = nextMoves%maxConsecutive + 1
				}
			}

			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}

			nextState := state{row: ni, col: nj, moves: nextMoves, dir: dir}
			nextHeatLoss := int(rune(grid[ni][nj]) - '0')
			if _, ok := minCost[nextState]; ok && minCost[nextState] <= curr.heatLoss+nextHeatLoss {
				continue
			}

			minCost[nextState] = curr.heatLoss + nextHeatLoss
			heap.Push(&pq, &item{heatLoss: curr.heatLoss + nextHeatLoss, state: nextState})
		}
	}

	return -1
}
