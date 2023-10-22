package naming

import (
	"fmt"
	"sort"
	"strings"
)

type generatedWords struct {
	Genitive string
	Definite string
	General  map[string][]string
	Names    []string
}

type Language struct {
	ApplyOrtho           bool
	ApplyMorph           bool
	Phonemes             map[string]string
	Morphemes            map[string][]string
	SyllableRestrictions []string
	ConsOrtho            orthoMapping
	VowelOrtho           orthoMapping
	Words                generatedWords
}

func BasicLanguage() *Language {
	lang := &Language{
		ApplyOrtho: false,
		ApplyMorph: false,
		Phonemes: orthoMapping{
			"C": "ptkmnls",
			"V": "aeiou",
			"S": "s",
			"F": "mn",
			"L": "rl",
		},
		SyllableRestrictions: []string{},
		ConsOrtho:            orthoMapping{},
		VowelOrtho:           orthoMapping{},
		Morphemes:            map[string][]string{},
		Words: generatedWords{
			General: map[string][]string{},
			Names:   []string{},
		},
	}
	lang.generateCommon()
	return lang
}

func OrthoLanguage() *Language {
	lang := BasicLanguage()
	lang.ApplyOrtho = true
	lang.generateCommon()
	return lang
}

func RandomLanguage(ortho bool, morph bool) (lang *Language) {
	lang = BasicLanguage()

	lang.Phonemes["C"] = randomItem(consonantSets)
	lang.Phonemes["V"] = randomItem(vowelSets)
	lang.Phonemes["S"] = randomItem(phonemeSSets)
	lang.Phonemes["F"] = randomItem(phonemeFSets)
	lang.Phonemes["L"] = randomItem(phonemeLSets)
	lang.ApplyOrtho = ortho
	lang.ApplyMorph = morph
	lang.ConsOrtho = randomItem(consonantOrthSets)
	lang.VowelOrtho = randomItem(vowelOrthSets)
	lang.Morphemes = map[string][]string{}
	lang.SyllableRestrictions = randomItem(restrictionSets)
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

// randomItem returns a random item from a map that uses string keys.
func randomItem[V any](m map[string]V) V {
	keys := sortedKeys[V](m)
	i := RandomRange(0, len(keys)-1)
	key := keys[i]
	return m[key]
}

// sortedKeys returns the sorted keys of a map that uses string keys.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
