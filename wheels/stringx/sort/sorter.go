package sort

type Sorter interface {
	Sort(s1, s2 string) bool
}

type SorterType uint8

const (
	CustomizedSorter SorterType = iota + 1
)

func NewSorter(sType SorterType) Sorter {
	switch sType {
	//case CustomizedSorter:
	//	return &customizedSorter{order: order}
	default:
		return &defaultSorter{order: order}
	}
}
