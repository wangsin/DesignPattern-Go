package BookShelf

import "fmt"

func Main() {
	bookShelf := NewEmptyBookShelf()
	bookShelf.AppendBook(NewBook("First Book"))
	bookShelf.AppendBook(NewBook("Second Book"))
	bookShelf.AppendBook(NewBook("Third Book"))
	bookShelf.AppendBook(NewBook("Fourth Book"))
	iterator := bookShelf.Iterator()
	for iterator.HasNext() {
		if book, ok := iterator.Next().(*Book); ok {
			fmt.Println(book.GetName())
		}
	}
}
