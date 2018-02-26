package ast

//定义一个Node，定义vistor访问方式，该Node可以接受vistor，vistor会遍历这个节点
type Node interface {
	//每一个Node都需要能接受一个Visitor进行访问
	//返回的Node是由Vistor来定义的
	//bool表示是否要继续向下访问，用于控制子节点的遍历
	Accept(v Visitor) (Node, bool)
}

//实现walk为vistor，该vistor对当前的Node进行处理，给出我们想要的操作
type Visitor interface {
	//判断Node是否需要访问
	//使用switch判断节点是否需要继续遍历自己的子节点
	Notify(Node) bool
	//处理节点
	//返回处理后的新的节点以及是否继续往下运行
	Visit(Node) (Node, bool)
}

//定义一个表达式类型的接口
type ExprNode interface {
	//实现Node接口
	Node
	//由struct node实现
	SetTag(int)
	GetTag() int
}
