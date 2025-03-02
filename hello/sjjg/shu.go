package main

import "fmt"

type TREEnode struct {
	data  string
	left  *TREEnode
	rigth *TREEnode
}

func preorder(root *TREEnode) { //先序
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preorder(root.left)
	preorder(root.rigth)
}

func inorder(root *TREEnode) { //中序
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.data)
	inorder(root.rigth)
}

func postorder(root *TREEnode) { //后序
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.rigth)
	fmt.Println(root.data)
}

func ccorder(root *TREEnode) { //层次遍历的广度优先
	if root == nil {
		return
	}
	que := new(linkqueue) //创建队列
	que.add(root)
	for {
		if que.len > 0 {
			ele := que.pop()
			fmt.Println(ele.data)
			if ele.left != nil {
				que.add(ele.left)
			}
			if ele.right != nil {
				que.add(ele.right)
			}
			if ele.right == nil && ele.lift == nil {
				break
			}
		}
	}
}

func main() {

}
