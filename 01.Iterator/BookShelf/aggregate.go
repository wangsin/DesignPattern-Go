/*
	Aggregate: 表示集合的接口 生成适用于集合的遍历器
*/
package BookShelf

type Aggregate interface {
	Iterator() Iterator
}
