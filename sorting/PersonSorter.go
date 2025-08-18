package main

import "sort"

type PersonSorter struct {
	people []Person
	by     By
}

func (s *PersonSorter) Len() int {
	return len(s.people)
}

func (s *PersonSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *PersonSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

type By func(p1, p2 *Person) bool

func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}
