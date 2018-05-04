package meta

import "strings"

type Index struct {
	HasPrimary  bool
	Idx         map[string][]string
	cardinality map[string]int
}

func NewIndex() *Index {
	return &Index{
		HasPrimary:  false,
		Idx:         make(map[string][]string),
		cardinality: make(map[string]int),
	}
}

func (i *Index) AddIndex(name string, column string, c int) {
	if _, ok := i.Idx[name]; ok {
		i.Idx[name] = append(i.Idx[name], strings.ToLower(column))
		if c > i.cardinality[name] {
			i.cardinality[name] = c
		}
	} else {
		i.Idx[name] = []string{strings.ToLower(column)}
		i.cardinality[name] = c
	}
}

func (i *Index) HaveIndex(idx []string) bool {
	for _, index := range i.Idx {
		if isSameAs(index, idx) {
			return true
		}
	}

	return false
}

func isSameAs(org []string, src []string) bool {
	if len(org) != len(src) {
		return false
	}

	for _, oarg := range org {
		flag := 0
		for _, sarg := range src {
			if strings.ToLower(oarg) == strings.ToLower(sarg) {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return false
		}
	}

	return true

}

func (i *Index) HaveSamePrimary(idx []string) bool {
	if i.HasPrimary {
		return isSameAs(i.Idx["PRIMARY"], idx)
	} else {
		return false
	}
}

func (i *Index) HaveSameIndex(idx []string) bool {
	for _, index := range i.Idx {
		if isSameAs(index, idx) && len(index) == len(idx) {
			return true
		}
	}

	return false
}

func (i *Index) Bestkey() (string, []string) {
	if len(i.Idx) == 0 {
		return "", nil
	}

	//有主键的情况下
	if i.HasPrimary {
		return "PRIMARY", i.Idx["PRIMARY"]
	}

	//无主键的情况下
	maxCardinality := -1

	res := ""
	for name, c := range i.cardinality {
		if c > maxCardinality {
			maxCardinality = c
			res = name
		}
	}

	if res != "" {
		return res, i.Idx[res]
	}

	panic("获得BestKey的过程中出现问题")
	return "", nil

}
