package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWords(t *testing.T) {
	words := []string{
		"qwe",
		"qWe",
		"qwE",
		"qw1",
		"w1",
		"w2",
		"w1",
		"q1",
		"Q1",
		"q1",
		"q1",
		"1q",
	}

	list := Words{}
	for _, s := range words {
		list.Add(s)
	}

	tops := list.Tops(2)
	require.Equal(t, Word{Text: "q1", Count: 4}, tops[0])
	require.Equal(t, Word{Text: "qwe", Count: 3}, tops[1])
}
