package main

import "fmt"

type dsu struct {
	pare []int
	rank []int
}

func newdsu(n int) *dsu {
	pare := make([]int, n)
	rank := make([]int, n)
	for i := range pare {
		pare[i] = i
		rank[i] = 1
	}
	return &dsu{pare, rank}
}

func (d *dsu) find(x int) int {
	for d.pare[x] != x {
		d.pare[x] = d.pare[d.pare[x]]
		x = d.pare[x]
	}
	return x
}

func (d *dsu) union(x, y int) {
	rootx := d.find(x)
	rooty := d.find(y)
	if rootx != rooty {
		if d.rank[rootx] > d.rank[rooty] {
			d.pare[rooty] = rootx
		} else if d.rank[rootx] < d.rank[rooty] {
			d.pare[rootx] = rooty
		} else {
			d.pare[rooty] = rootx
			d.rank[rootx]++
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	iptoid := make(map[string]int)
	idtoip := make(map[int]string)

	for i := 0; i < n; i++ {
		var ip string
		var id int
		fmt.Scan(&ip, &id)
		iptoid[ip] = id
		idtoip[id] = ip
	}

	dsu := newdsu(n + 1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		dsu.union(a, b)
	}

	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var ip1, ip2 string
		fmt.Scan(&ip1, &ip2)
		id1 := iptoid[ip1]
		id2 := iptoid[ip2]
		if dsu.find(id1) == dsu.find(id2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
