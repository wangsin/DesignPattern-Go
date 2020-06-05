/*
	BookShelf: 图书集合类 需要实现Aggregate接口
*/
package BookShelf

type BookShelf struct {
	books []*Book
}

func NewEmptyBookShelf() *BookShelf {
	return &BookShelf{books: make([]*Book, 0)}
}

func (bs *BookShelf) GetBookByIndex(index int) *Book {
	if index >= len(bs.books) {
		return nil
	}
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookShelf) GetSize() int {
	return len(bs.books)
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}