package meta

import "strings"

type Index struct {
	HasPrimary bool
	Idx        map[string][]string
}

func NewIndex() *Index {
	return &Index{
		HasPrimary: false,
		Idx:        make(map[string][]string),
	}
}

func (i *Index) AddIndex(name string, column string) {
	if _, ok := i.Idx[name]; ok {
		i.Idx[name] = append(i.Idx[name], strings.ToLower(column))
	} else {
		i.Idx[name] = []string{strings.ToLower(column)}
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
	if i.HasPrimary{
		return isSameAs(i.Idx["PRIMARY"],idx)
	}else{
		return false
	}
}

func (i *Index) HaveSameIndex(idx []string) bool {
	for _, index := range i.Idx {
		if isSameAs(index, idx) && len(index)==len(idx){
			return true
		}
	}

	return false
}