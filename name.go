package main

import (
	"math/rand"
	"strings"
)

func (lang *Language) makeName(minLen int, maxLen int, p wordParams, joiners string, group string) string {
	if minLen <= 0 {
		minLen = 5
	}

	if maxLen < minLen {
		maxLen = minLen
	} else if maxLen <= 0 {
		maxLen = 12
	}

	if p.minSyllables <= 0 {
		p.minSyllables = 1
	}

	if p.maxSyllables < p.minSyllables {
		p.maxSyllables = p.minSyllables
	} else if maxLen <= 0 {
		p.maxSyllables = 2
	}

	for {
		name := ""

		if rand.Float32() < 0.5 {
			name = strings.Title(lang.getWord(p, group))
		} else {
			g := ""
			if rand.Float32() < 0.6 {
				g = group
			}
			w1 := strings.Title(lang.getWord(p, g))
			g = ""
			if rand.Float32() < 0.6 {
				g = group
			}
			w2 := strings.Title(lang.getWord(p, g))
			if w1 == w2 {
				continue
			}

			if len(joiners) > 0 {
				join := randomRuneFromString(joiners)

				if rand.Float32() > 0.5 {
					name = strings.Join([]string{w1, w2}, join)
				} else {
					name = strings.Join([]string{w1, lang.Words.Genitive, w2}, join)
				}
			}
		}

		if len(joiners) > 0 {
			join := randomRuneFromString(joiners)

			if rand.Float32() < 0.1 {
				name = strings.Join([]string{lang.Words.Definite, name}, join)
			}
		}

		if (len(name) < minLen) || (len(name) > maxLen) {
			continue
		}

		used := false
		for _, name2 := range lang.Words.Names {
			if strings.Contains(name, name2) || strings.Contains(name2, name) {
				used = true
				break
			}
		}

		if used {
			continue
		}

		lang.Words.Names = append(lang.Words.Names, name)

		return name
	}
}
