package main

type UnionFind struct {
	Fa     []int
	Groups int // 连通量的个数
}

func NewUnionFind(n int) UnionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return UnionFind{fa, n}
}

func (u *UnionFind) FindR(x int) int {
	if u.Fa[x] != x {
		u.Fa[x] = u.FindR(u.Fa[x])
	}
	return u.Fa[x]
}

func (u *UnionFind) Merge(from, to int) (newRoot int) {
	x, y := u.FindR(from), u.FindR(to)
	if x == y {
		return -1
	}
	u.Fa[x] = y
	u.Groups--
	return y
}

func (u *UnionFind) Same(x, y int) bool {
	return u.FindR(x) == u.FindR(y)
}
