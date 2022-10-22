package naming

import (
	"math"
	"math/rand"
	"strings"
)

type NameParams struct {
	MinLength  int
	MaxLength  int
	WordParams *WordParams
	Joiners    string
	Group      string
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
		// which returns from 0.0 to 1.0.
		// In Go, rand.Float32() returns a value from -1.0 to +1.0
		if randFloat32Abs() < 0.5 {
			name = strings.Title(lang.GetWord(params.WordParams, params.Group))
		} else {
			g := ""
			if randFloat32Abs() < 0.6 {
				g = params.Group
			}
			w1 := strings.Title(lang.GetWord(params.WordParams, g))
			g = ""
			if randFloat32Abs() < 0.6 {
				g = params.Group
			}
			w2 := strings.Title(lang.GetWord(params.WordParams, g))
			if w1 == w2 {
				continue
			}

			if joinersLen > 0 {
				join := RandomRuneFromString(params.Joiners)
				if randFloat32Abs() > 0.5 {
					name = strings.Join([]string{w1, w2}, join)
				} else {
					name = strings.Join([]string{w1, lang.Words.Genitive, w2}, join)
				}
			}
		}

		if joinersLen > 0 {
			join := RandomRuneFromString(params.Joiners)
			if randFloat32Abs() < 0.1 {
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

func randFloat32Abs() float32 {
	return float32(math.Abs(float64(rand.Float32())))
}
