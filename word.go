package main

import "math/rand"

type structureList []string

type wordParams struct {
	minSyllables int
	maxSyllables int
	structure    structureList
}

func (list *structureList) random() string {
	l := len(*list)
	i := rand.Intn(l)
	return (*list)[i]
}

func (lang *Language) makeWord(p wordParams, group string) string {
	numSyllables := randomRange(p.minSyllables, p.maxSyllables+1)

	keys := make([]string, numSyllables)

	keys[randomRange(0, numSyllables)] = group

	w := ""

	for i := 0; i < numSyllables; i++ {
		w += lang.getMorpheme(p.structure.random(), keys[i])
	}

	return w
}

func (lang *Language) getWord(p wordParams, group string) string {
	ws := []string{}

	if val, ok := lang.Words.General[group]; ok {
		ws = val
	}

	extras := 3
	if len(group) > 0 {
		extras = 2
	}

	for {
		n := randomRange(0, len(ws)+extras)

		if n < len(ws) {
			return ws[n]
		}

		w := lang.makeWord(p, group)
		exists := false

		for _, v := range lang.Words.General {
			for _, item := range v {
				if item == w {
					exists = true
					break
				}

				if exists {
					break
				}
			}
		}

		if exists {
			continue
		}

		lang.Words.General[group] = append(ws, w)

		return w
	}
}
