package book

type Book struct {
	Pages uint
}

func (b Book) Less(a Book) bool {
	return b.Pages < a.Pages
}
