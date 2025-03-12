package main

import (
	"cmp"
	"slices"
	"strings"
)

type (
	Words map[string]int

	Word struct {
		Text  string
		Count int
	}
)

func (w Words) Add(s string) {
	w[strings.ToLower(s)] += 1
}

func (w Words) Tops(count int) []Word {
	if count <= 0 {
		return []Word{}
	}

	tops := make([]Word, 0, len(w))
	for k, v := range w {
		tops = append(tops, Word{Text: k, Count: v})
	}

	slices.SortFunc(tops, func(i, j Word) int {
		return cmp.Compare(j.Count, i.Count)
	})

	if count < len(tops) {
		tops = tops[:count]
	}
	return tops
}
