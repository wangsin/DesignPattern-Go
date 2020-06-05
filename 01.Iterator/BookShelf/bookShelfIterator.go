/*
	BookShelfIterator: 实现了Iterator的遍历器 由Aggregate生成
	使用Iterator模式可以忽略底层Map实现方式 进而减少上游代码修改 提高代码复用性
*/
package BookShelf

type BookShelfIterator struct {
	bookShelf *BookShelf
	index int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bookShelf, index: 0}
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.GetSize()
}

func (bsi *BookShelfIterator) Next() interface{} {
	defer func() {
		bsi.index++
	}()
	return bsi.bookShelf.GetBookByIndex(bsi.index)
}