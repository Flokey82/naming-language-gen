package naming

import (
	"github.com/dlclark/regexp2"
	"math/rand"
)

func (lang Language) spell(syllable string) string {
	if !lang.ApplyOrtho {
		return syllable
	}

	var s []rune
	for _, c := range syllable {
		str := string(c)
		if val, ok := lang.ConsOrtho[str]; ok {
			s = append(s, []rune(val)...)
		} else if val, ok := lang.VowelOrtho[str]; ok {
			s = append(s, []rune(val)...)
		} else if val, ok := defaultOrtho[str]; ok {
			s = append(s, []rune(val)...)
		} else {
			s = append(s, c)
		}
	}

	return string(s)
}

func (lang Language) makeSyllable(structure string) string {
	for {
		syllable := ""
		structureLen := len(structure)

		for i := 0; i < structureLen; i++ {
			ptype := string(structure[i])

			if (i < structureLen-1) && structure[i+1] == '?' {
				i++
				if rand.Float32() < 0.5 {
					continue
				}
			}

			syllable += RandomRuneFromString(lang.Phonemes[ptype])
		}

		var bad bool
		for _, restriction := range lang.SyllableRestrictions {
			exp := regexp2.MustCompile(restriction, 0)
			matched, _ := exp.MatchString(syllable)
			if matched {
				bad = true
				break
			}
		}

		if bad {
			continue
		}

		return lang.spell(syllable)
	}
}
