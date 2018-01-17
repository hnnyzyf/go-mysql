package myast



//定义List类型
type List []Node


func NewList(n Node) List{
	switch node:=n.(type){
		default:
			return List{node}
	}
}
