package usecase

import (
	"sort"
	"stockbit/module/anagram/model"
	"strings"
)

func UseCaseAnagram(b *model.Anagram) (rs interface{}, err error) {
	list := make(map[string][]string)
	r := make([][]string, 0)

	for _, x := range b.Data {
		key := sortStr(x)
		list[key] = append(list[key], x)
	}

	for _, y := range list {
		r = append(r, y)
	}

	return model.ResponseAnagram{
		Request: b.Data,
		Result:  r,
	}, nil
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
