package naming

import "math/rand"

type structureList []string

type WordParams struct {
	MinSyllables int
	MaxSyllables int
	Structure    structureList
}

func (list *structureList) random() string {
	l := len(*list)
	i := rand.Intn(l)
	return (*list)[i]
}

func (lang Language) makeWord(p *WordParams, group string) (word string) {
	numSyllables := RandomRange(p.MinSyllables, p.MaxSyllables+1)

	keys := make([]string, numSyllables)

	keys[RandomRange(0, numSyllables)] = group

	for i := 0; i < numSyllables; i++ {
		word += lang.makeMorpheme(p.Structure.random(), keys[i])
	}

	return
}

func (lang *Language) GetWord(p *WordParams, group string) (word string) {
	words := []string{}

	if val, ok := lang.Words.General[group]; ok {
		words = val
	}

	extras := 3
	if len(group) > 0 {
		extras = 2
	}

	for {
		n := RandomRange(0, len(words)+extras)

		if n < len(words) {
			return words[n]
		}

		word = lang.makeWord(p, group)
		exists := false

		for _, v := range lang.Words.General {
			if contains(v, word) {
				exists = true
				break
			}
		}

		if exists {
			continue
		}

		lang.Words.General[group] = append(words, word)

		return word
	}
}
