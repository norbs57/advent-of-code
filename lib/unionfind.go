package lib

// https://www.cs.princeton.edu/~rs/AlgsDS07/01UnionFind.pdf
// Could also define a Union-Find by "rank"="depth"
// Also called "Disjoint Set"
type UnionFind struct {
	// Union-find data structure with size for balancing
	Parent []int
	Size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{}
	uf.Parent = make([]int, n)
	uf.Size = make([]int, n)
	for i := range uf.Parent {
		uf.Parent[i] = i
		uf.Size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Copy() *UnionFind {
	res := &UnionFind{}
	res.Parent = make([]int, len(uf.Parent))
	res.Size = make([]int, len(uf.Size))
	copy(res.Parent, uf.Parent)
	copy(res.Size, uf.Size)
	return res
}

func (uf *UnionFind) Extend() {
	uf.Parent = append(uf.Parent, len(uf.Parent))
	uf.Size = append(uf.Size, 1)
}

func (uf *UnionFind) Find(i int) int {
	for i != uf.Parent[i] {
		uf.Parent[i] = uf.Parent[uf.Parent[i]]
		i = uf.Parent[i]
	}
	return i
}

func (uf *UnionFind) Merge(p, q int) {
	i := uf.Find(p)
	j := uf.Find(q)
	if i == j {
		return
	}
	if uf.Size[i] < uf.Size[j] {
		uf.Parent[i] = j
		uf.Size[j] += uf.Size[i]
	} else {
		uf.Parent[j] = i
		uf.Size[i] += uf.Size[j]
	}
}

func (union *UnionFind) SameRoot(p, q int) bool {
	return union.Find(p) == union.Find(q)
}
