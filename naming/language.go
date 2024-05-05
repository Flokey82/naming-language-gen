package naming

import (
	"fmt"
	"math/rand"
	"strings"
)

type generatedWords struct {
	Genitive string
	Definite string
	General  map[string][]string
	Names    []string
}

// Clone returns a deep copy of the generatedWords.
func (gw generatedWords) Clone() generatedWords {
	clone := generatedWords{
		Genitive: gw.Genitive,
		Definite: gw.Definite,
		General:  make(map[string][]string),
		Names:    make([]string, len(gw.Names)),
	}
	for k, v := range gw.General {
		clone.General[k] = make([]string, len(v))
		copy(clone.General[k], v)
	}
	copy(clone.Names, gw.Names)
	return clone
}

type Language struct {
	Seed                 int64
	Rnd                  *rand.Rand
	ApplyOrtho           bool
	ApplyMorph           bool
	Phonemes             map[string]string
	Morphemes            map[string][]string
	SyllableRestrictions []string
	ConsOrtho            orthoMapping
	VowelOrtho           orthoMapping
	Words                generatedWords
}

// Fork returns a deep copy of the language.
func (lang *Language) Fork(newSeed int64) *Language {
	clone := &Language{
		Seed:                 newSeed,
		Rnd:                  rand.New(rand.NewSource(newSeed)),
		ApplyOrtho:           lang.ApplyOrtho,
		ApplyMorph:           lang.ApplyMorph,
		Phonemes:             make(map[string]string),
		Morphemes:            make(map[string][]string),
		SyllableRestrictions: make([]string, len(lang.SyllableRestrictions)),
		ConsOrtho:            lang.ConsOrtho,
		VowelOrtho:           lang.VowelOrtho,
		Words:                lang.Words.Clone(),
	}

	copy(clone.SyllableRestrictions, lang.SyllableRestrictions)

	for k, v := range lang.Phonemes {
		clone.Phonemes[k] = v
	}

	for k, v := range lang.Morphemes {
		clone.Morphemes[k] = make([]string, len(v))
		copy(clone.Morphemes[k], v)
	}

	return clone
}

func BasicLanguage(seed int64) *Language {
	lang := &Language{
		Seed:       seed,
		Rnd:        rand.New(rand.NewSource(seed)),
		ApplyOrtho: false,
		ApplyMorph: false,
		Phonemes: orthoMapping{
			"C": "ptkmnls",
			"V": "aeiou",
			"S": "s",
			"F": "mn",
			"L": "rl",
		},
		SyllableRestrictions: nil,
		ConsOrtho:            make(orthoMapping),
		VowelOrtho:           make(orthoMapping),
		Morphemes:            make(map[string][]string),
		Words: generatedWords{
			General: make(map[string][]string),
			Names:   nil,
		},
	}
	lang.generateCommon()
	return lang
}

func OrthoLanguage(seed int64) *Language {
	lang := BasicLanguage(seed)
	lang.ApplyOrtho = true
	lang.generateCommon()
	return lang
}

func RandomLanguage(ortho, morph bool, seed int64) (lang *Language) {
	lang = BasicLanguage(seed)
	lang.Phonemes["C"] = randomItem(consonantSets, lang.Rnd)
	lang.Phonemes["V"] = randomItem(vowelSets, lang.Rnd)
	lang.Phonemes["S"] = randomItem(phonemeSSets, lang.Rnd)
	lang.Phonemes["F"] = randomItem(phonemeFSets, lang.Rnd)
	lang.Phonemes["L"] = randomItem(phonemeLSets, lang.Rnd)
	lang.ApplyOrtho = ortho
	lang.ApplyMorph = morph
	lang.ConsOrtho = randomItem(consonantOrthSets, lang.Rnd)
	lang.VowelOrtho = randomItem(vowelOrthSets, lang.Rnd)
	lang.Morphemes = make(map[string][]string)
	lang.SyllableRestrictions = randomItem(restrictionSets, lang.Rnd)
	lang.generateCommon()
	return
}

func (lang Language) Describe() {
	fmt.Printf("-> apply ortho: %v\n", lang.ApplyOrtho)
	fmt.Printf("-> apply morph: %v\n", lang.ApplyMorph)

	fmt.Printf("-> phonemes:\n")
	fmt.Printf("      C:  %v\n", lang.Phonemes["C"])
	fmt.Printf("      V:  %v\n", lang.Phonemes["V"])
	fmt.Printf("      S:  %v\n", lang.Phonemes["S"])
	fmt.Printf("      F:  %v\n", lang.Phonemes["F"])
	fmt.Printf("      L:  %v\n", lang.Phonemes["L"])

	fmt.Printf("-> restrictions: %v\n", strings.Join(lang.SyllableRestrictions[:], ", "))

	if lang.ApplyOrtho {
		fmt.Printf("-> consonant ortho:\n")
		for k, v := range lang.ConsOrtho {
			fmt.Printf("      %v  =>  %v\n", k, v)
		}

		fmt.Printf("-> vowel ortho:\n")
		for k, v := range lang.VowelOrtho {
			fmt.Printf("      %v  =>  %v\n", k, v)
		}
	}

	if lang.ApplyMorph {
		fmt.Printf("-> morphemes:\n")
		for k, v := range lang.Morphemes {
			fmt.Printf(" >  '%v'\n", k)
			fmt.Printf("      %v\n", strings.Join(v[:], ", "))
		}
	}
}

func (lang *Language) generateCommon() {
	lang.Words.Genitive = lang.makeMorpheme("C?VC?", "of")
	lang.Words.Definite = lang.makeMorpheme("C?VC?", "the")
}
