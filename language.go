package main

import (
	"fmt"
	"strings"
)

type generatedWords struct {
	Genitive string
	Definite string
	General  map[string][]string
	Names    []string
}

type Language struct {
	ApplyOrtho bool
	ApplyMorph bool

	Phonemes map[string]string

	SyllableRestrictions []string

	ConsOrtho  orthoMapping
	VowelOrtho orthoMapping

	Morphemes map[string][]string

	Words generatedWords
}

func (lang *Language) generateCommon() {
	lang.Words.Genitive = lang.getMorpheme("C?VC?", "of")
	lang.Words.Definite = lang.getMorpheme("C?VC?", "the")
}

func BasicLanguage() Language {
	lang := Language{
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

		ConsOrtho:  orthoMapping{},
		VowelOrtho: orthoMapping{},

		Morphemes: map[string][]string{},

		Words: generatedWords{
			Genitive: "",
			Definite: "",
			General:  map[string][]string{},
			Names:    []string{},
		},
	}

	lang.generateCommon()

	return lang
}

func OrthoLanguage() Language {
	lang := BasicLanguage()

	lang.ApplyOrtho = true

	lang.generateCommon()

	return lang
}

func RandomLanguage(ortho bool, morph bool) Language {
	lang := BasicLanguage()

	lang.Phonemes["C"] = consonantSets.random()
	lang.Phonemes["V"] = vowelSets.random()
	lang.Phonemes["S"] = phonemeSSets.random()
	lang.Phonemes["F"] = phonemeFSets.random()
	lang.Phonemes["L"] = phonemeLSets.random()

	lang.ApplyOrtho = ortho
	lang.ApplyMorph = morph

	lang.ConsOrtho = consonantOrthSets.random()
	lang.VowelOrtho = vowelOrthSets.random()

	lang.Morphemes = map[string][]string{}
	lang.SyllableRestrictions = restrictionSets.random()

	lang.generateCommon()
	return lang
}

func (lang *Language) Describe() {
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
