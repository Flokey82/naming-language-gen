package naming

import (
	"strings"
)

type NameParams struct {
	MinLength  int
	MaxLength  int
	WordParams *WordParams
	Joiners    string
	Group      string
}

// Clone returns a deep copy of the NameParams.
func (p *NameParams) Clone() *NameParams {
	return &NameParams{
		MinLength:  p.MinLength,
		MaxLength:  p.MaxLength,
		WordParams: p.WordParams.Clone(),
		Joiners:    p.Joiners,
		Group:      p.Group,
	}
}

func (lang *Language) MakeName(params *NameParams) (name string) {
	if params.MinLength <= 0 {
		params.MinLength = 5
	}

	if params.MaxLength < params.MinLength {
		params.MaxLength = params.MinLength
	} else if params.MaxLength <= 0 {
		params.MaxLength = 12
	}

	if params.WordParams.MinSyllables <= 0 {
		params.WordParams.MinSyllables = 1
	}

	if params.WordParams.MaxSyllables < params.WordParams.MinSyllables {
		params.WordParams.MaxSyllables = params.WordParams.MinSyllables
	} else if params.MaxLength <= 0 {
		params.WordParams.MaxSyllables = 2
	}

	joinersLen := len(params.Joiners)

	for {
		// NOTE: In the original JavaScript version we use Math.random(),
		// which returns from 0.0 to 1.0. In Go, rand.Float64() returns a
		// value from -1.0 to +1.0.
		if randFloat64Abs(lang.Rnd) < 0.5 {
			name = strings.Title(lang.GetWord(params.WordParams, params.Group))
		} else {
			var group string
			if randFloat64Abs(lang.Rnd) < 0.6 {
				group = params.Group
			}
			word1 := strings.Title(lang.GetWord(params.WordParams, group))
			group = ""
			if randFloat64Abs(lang.Rnd) < 0.6 {
				group = params.Group
			}
			word2 := strings.Title(lang.GetWord(params.WordParams, group))
			if word1 == word2 {
				continue
			}

			if joinersLen > 0 {
				join := RandomRuneFromString(params.Joiners, lang.Rnd)
				if randFloat64Abs(lang.Rnd) > 0.5 {
					name = strings.Join([]string{word1, word2}, join)
				} else {
					name = strings.Join([]string{word1, lang.Words.Genitive, word2}, join)
				}
			}
		}

		if joinersLen > 0 {
			join := RandomRuneFromString(params.Joiners, lang.Rnd)
			if randFloat64Abs(lang.Rnd) < 0.1 {
				name = strings.Join([]string{lang.Words.Definite, name}, join)
			}
		}

		if (len(name) < params.MinLength) || (len(name) > params.MaxLength) {
			continue
		}

		var used bool
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
