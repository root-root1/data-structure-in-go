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
	var intersectionSet = &Set{}
	intersectionSet.New()
	var value int
	for value = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectionSet.AddElement(value)
		}
	}
	return intersectionSet
}

func (set *Set) UnionSet(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value = range set.integerMap {
		unionSet.AddElement(value)
	}
	for value = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}
