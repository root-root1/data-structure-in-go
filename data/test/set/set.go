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

func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersection = &Set{}
	intersection.New()
	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersection.AddElement(value)
		}
	}
	return intersection
}

func (set *Set) Union(anotherSet *Set) *Set {
	var union = &Set{}
	union.New()

	for value := range set.integerMap {
		union.AddElement(value)
	}
	for value := range anotherSet.integerMap {
		union.AddElement(value)
	}
	return union
}
