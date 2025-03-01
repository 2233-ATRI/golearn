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

func main() {

}
