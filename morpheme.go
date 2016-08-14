package main

import "math/rand"

func (lang *Language) getMorpheme(structure string, group string) string {
	if !lang.ApplyMorph {
		return lang.makeSyllable(structure)
	}

	list := []string{}

	if val, ok := lang.Morphemes[group]; ok {
		list = val
	}

	extras := 10
	if len(group) > 0 {
		extras = 1
	}

	for {
		n := rand.Intn(len(list) + extras)

		if n < len(list) {
			return list[n]
		}

		morph := lang.makeSyllable(structure)

		exists := false

		for _, v := range lang.Morphemes {
			for _, item := range v {
				if item == morph {
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

		lang.Morphemes[group] = append(list, morph)

		return morph
	}
}
