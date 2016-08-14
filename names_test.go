package main

import (
	"fmt"
	"strings"
	"testing"
)

var nameTests = []struct {
	group         string
	structures    structureList
	vorthoMapping string
	corthoMapping string
}{
	{"lang 1", structureList{"C?VL?C"}, "Diphthongs", "Slavic"},
	{"lang 2", structureList{"S?CVC?"}, "Ümlauts", "German"},
	{"lang 3", structureList{"CVV?C", "C?VL?C"}, "Welsh", "Chinese"},
	{"lang 4", structureList{"CL?VF"}, "Doubles", "French"},
	{"lang 5", defaultSyllableStructures, "Welsh", "German"},
}

func TestNames(t *testing.T) {
	lang := OrthoLanguage()

	lang.ApplyMorph = true
	lang.SyllableRestrictions = restrictionSets["Hard clusters"]

	const joiners = "  -"
	const minLength = 4
	const maxLength = 16
	const minSyllables = 1
	const maxSyllables = 4

	for _, val := range nameTests {
		lang.VowelOrtho = vowelOrthSets[val.vorthoMapping]
		lang.ConsOrtho = consonantOrthSets[val.corthoMapping]

		list := []string{}

		lang.generateCommon()

		p := wordParams{
			minSyllables,
			maxSyllables,
			val.structures,
		}

		for i := 0; i < 10; i++ {
			list = append(list, lang.makeName(minLength, maxLength, p, joiners, val.group))
		}

		fmt.Printf("Make names [%v]: %v\n", val.group, strings.Join(list[:], ", "))
	}
}
