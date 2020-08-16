package list

type List interface {
	Push(x interface{})
	Unshift(x interface{})
	Pop() interface{}
	Poll() interface{}
	Peek() interface{}
	Check() interface{}
	Set(i int, x interface{})
	Get(i int) interface{}
	Reverse()
	Length() int
}
