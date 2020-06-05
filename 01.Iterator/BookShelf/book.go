/*
	Book: 图书类 单个元素对象
*/
package BookShelf

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (b *Book) GetName() string {
	return b.name
}