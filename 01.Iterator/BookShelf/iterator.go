/*
	Iterator: 遍历集合的接口 有集合本身创建实现
*/
package BookShelf

type Iterator interface {
	HasNext() bool
	Next() interface{}
}