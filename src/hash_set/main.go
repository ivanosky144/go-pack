package hashset

type HashSet struct {
	items map[int]struct{}
}

func NewSet() *HashSet {
	return &HashSet{
		items: make(map[int]struct{}),
	}
}
