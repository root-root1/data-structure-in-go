package set

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(data int) {
	if !set.ContainsElement(data) {
		set.integerMap[data] = true
	}
}

func (set *Set) DeleteElement(data int) {
	delete(set.integerMap, data)
}

func (set *Set) ContainsElement(data int) bool {
	_, exist := set.integerMap[data]
	return exist
}
