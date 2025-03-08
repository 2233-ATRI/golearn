package main

import (
	"fmt"
)

type Pao struct {
	x, y int
	id   int
}

func main() {
	var n int
	fmt.Scan(&n)
	paos := make([]Pao, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		paos[i] = Pao{x: x, y: i + 1, id: i + 1}
	}
	for i := 0; i < n; i++ {
		pao := paos[i]
		x1 := pao.x
		y1 := pao.y
		targets := make(map[int]bool)

		var rackright *Pao
		for j := 0; j < n; j++ {
			p := paos[j]
			if p.id == pao.id {
				continue
			}
			if p.x > x1 && (rackright == nil || p.x < rackright.x) {
				rackright = &p
			}
		}
		if rackright != nil {
			var target *Pao
			for j := 0; j < n; j++ {
				p := paos[j]
				if p.id == pao.id {
					continue
				}
				if p.x > rackright.x && (target == nil || p.x < target.x) {
					target = &p
				}
			}
			if target != nil {
				targets[target.id] = true
			}
		}

		var rackleft *Pao
		for j := 0; j < n; j++ {
			p := paos[j]
			if p.id == pao.id {
				continue
			}
			if p.x < x1 && (rackleft == nil || p.x > rackleft.x) {
				rackleft = &p
			}
		}
		if rackleft != nil {
			var target *Pao
			for j := 0; j < n; j++ {
				p := paos[j]
				if p.id == pao.id {
					continue
				}
				if p.x < rackleft.x && (target == nil || p.x > target.x) {
					target = &p
				}
			}
			if target != nil {
				targets[target.id] = true
			}
		}
		var rackup *Pao
		for j := 0; j < n; j++ {
			p := paos[j]
			if p.id == pao.id {
				continue
			}
			if p.x == x1 && p.y > y1 && (rackup == nil || p.y < rackup.y) {
				rackup = &p
			}
		}
		if rackup != nil {
			var target *Pao
			for j := 0; j < n; j++ {
				p := paos[j]
				if p.id == pao.id {
					continue
				}
				if p.x == x1 && p.y > rackup.y && (target == nil || p.y < target.y) {
					target = &p
				}
			}
			if target != nil {
				targets[target.id] = true
			}
		}

		var rackdown *Pao
		for j := 0; j < n; j++ {
			p := paos[j]
			if p.id == pao.id {
				continue
			}
			if p.x == x1 && p.y < y1 && (rackdown == nil || p.y > rackdown.y) {
				rackdown = &p
			}
		}
		if rackdown != nil {
			var target *Pao
			for j := 0; j < n; j++ {
				p := paos[j]
				if p.id == pao.id {
					continue
				}
				if p.x == x1 && p.y < rackdown.y && (target == nil || p.y > target.y) {
					target = &p
				}
			}
			if target != nil {
				targets[target.id] = true
			}
		}
		fmt.Println(len(targets))
	}
} //对每一个炮都查看其他的横纵坐标相同的元素的位置，记录下当前的最大值，当遍历结束输出该值，然后查看下一个炮
