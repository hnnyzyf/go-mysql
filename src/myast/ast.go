package myast



//定义一个Node，定义vistor访问方式，该Node可以接受vistor，vistor会遍历这个节点
type Node interface{
	Accept(v Visitor)
}


//实现walk为vistor，该vistor对当前的Node进行处理，给出我们想要的操作
type Visitor interface {
	Walk()
}


