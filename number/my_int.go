package number

type MyInt int

func (i MyInt) Less(j MyInt) bool {
	return i < j
}
