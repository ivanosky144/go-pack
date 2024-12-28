package hashset

func (set *HashSet) Add(value int) {
	set.items[value] = struct{}{}
}

func (set *HashSet) Remove(value int) {
	delete(set.items, value)
}

func (set *HashSet) Size() int {
	return len(set.items)
}
