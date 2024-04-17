package graph

import "container/heap"

type dijkstraPair struct{ x, y, dis int }
type dijkstraHeap []dijkstraPair

func (h dijkstraHeap) Len() int           { return len(h) }
func (h dijkstraHeap) Less(i, j int) bool { return h[i].dis > h[j].dis } // 最大堆，修改符号可以创建最小堆
func (h dijkstraHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现 heap.interface 接口

func (h *dijkstraHeap) Push(v any)          { *h = append(*h, v.(dijkstraPair)) }
func (h *dijkstraHeap) Pop() (v any)        { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *dijkstraHeap) push(v dijkstraPair) { heap.Push(h, v) }
func (h *dijkstraHeap) pop() dijkstraPair   { return heap.Pop(h).(dijkstraPair) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

// 模板题： https://leetcode.cn/problems/path-with-maximum-minimum-value/description/

func DijkstraShortestPathOnGrid(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dis := make([][]int, n)
	const inf int = 1e18
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = -inf
		}
	}
	h := dijkstraHeap{{0, 0, grid[0][0]}}
	for h.Len() > 0 {
		cur := h.pop()
		if dis[cur.x][cur.y] != cur.dis {
			continue
		}
		for _, d := range dirs {
			nx, ny := cur.x+d[0], cur.y+d[1]
			if nx < 0 || ny < 0 || nx >= n || ny >= m {
				continue
			}
			if dis[nx][ny] < min(cur.dis, grid[nx][ny]) {
				dis[nx][ny] = min(cur.dis, grid[nx][ny])
				h.push(dijkstraPair{nx, ny, dis[nx][ny]})
			}
		}
	}
	return dis[n-1][m-1]
}
