package book

type SortableBook struct {
	Pages uint
}

type SortableBooks []SortableBook

func (b SortableBooks) Less(i, j int) bool {
	return b[i].Pages < b[j].Pages
}

func (b SortableBooks) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b SortableBooks) Len() int {
	return len(b)
}
