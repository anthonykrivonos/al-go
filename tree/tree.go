package tree

type Tree interface {
	PreOrderPrint()
	InOrderPrint()
	PostOrderPrint()
}

type Node interface {
	Value() interface{}
}
