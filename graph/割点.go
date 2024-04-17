package graph

import (
	"slices"
)

// 返回是否为割点
func findCutVertices(n int, g [][]int) []bool {
	isCut := make([]bool, n)
	dfn := make([]int, n)

	var tarjan func(v, fa int) int
	dfsClock := 0
	tarjan = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			if dfn[w] == 0 {
				childCnt++
				lowW := tarjan(w, v)
				lowV = min(lowV, lowW)
				if lowW >= dfn[v] { //
					isCut[v] = true
				}
			} else if w != fa {
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 {
			// 只有一个儿子的树根
			isCut[v] = true
		}
		return lowV
	}

	for v, time := range dfn {
		if time == 0 {
			tarjan(v, -1)
		}
	}
	return isCut
}

// https://oi-wiki.org/graph/scc/

// 返回所有的点所在的强连通分量号
func findSCC(n int, g [][]int) []int {
	scc := make([]int, n)
	low := make([]int, n)
	dfn := make([]int, n)
	st := make([]int, n)   // 栈
	vis := make([]bool, n) // 是否在栈中
	dfsClock, tp, sc := 0, 0, 0

	var tarjan func(u int)
	tarjan = func(u int) {
		dfsClock++
		low[u] = dfsClock
		dfn[u] = dfsClock
		st[tp+1] = u
		tp++
		vis[u] = true
		for _, v := range g[u] {
			if dfn[v] == 0 {
				tarjan(v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}
		if dfn[u] == low[u] {
			sc++
			for st[tp] != u {
				scc[st[tp]] = sc
				vis[st[tp]] = false
				tp--
			}
			scc[st[tp]] = sc
			vis[st[tp]] = false
			tp--
		}
	}

	for u, time := range dfn {
		if time == 0 {
			tarjan(u)
		}
	}
	return scc
}

// 缩点，灵神的实现
func sccTarjan(g [][]int) ([][]int, []int) {
	var scc [][]int
	var st []int
	dfn := make([]int, len(g)) // 值从1开始
	dfsClock := 0
	inSt := make([]bool, len(g))

	var tarjan func(int) int
	tarjan = func(v int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		st = append(st, v)
		inSt[v] = true
		for _, w := range g[v] {
			if dfn[w] == 0 {
				lowW := tarjan(w)
				lowV = min(lowV, lowW)
			} else if inSt[w] {
				lowV = min(lowV, dfn[w])
			}
		}
		if dfn[v] == lowV {
			var comp []int
			for {
				w := st[len(st)-1]
				st = st[:len(st)-1]
				inSt[w] = false
				comp = append(comp, w)
				if w == v {
					break
				}
			}
			scc = append(scc, comp)
		}
		return lowV
	}
	for i, time := range dfn {
		if time == 0 {
			tarjan(i)
		}
	}
	slices.Reverse(scc) // 得到的scc是拓扑排序逆序的
	// 缩点
	sid := make([]int, len(g))
	for i, cc := range scc {
		for _, v := range cc {
			sid[v] = i
		}
	}
	// 重新构造图
	ns := len(scc)
	g2 := make([][]int, ns)
	deg := make([]int, ns)
	for v, ws := range g {
		v = sid[v]
		for _, w := range ws {
			w = sid[w]
			if v != w {
				g2[v] = append(g2[v], w)
				deg[w]++
			}
		}
	}
	return scc, sid
}
