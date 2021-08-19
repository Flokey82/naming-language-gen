package naming

import (
	"fmt"
	"testing"
)

var orthoSpellTests = []struct {
	in  string
	out string
}{
	{"AEIOU", "áéíóú"},
	{"aeiou", "aeiou"},
	{"ʧʔɣan", "ch‘ghan"},
}

func TestSpellOrtho(t *testing.T) {
	lang := OrthoLanguage()

	for _, tt := range orthoSpellTests {
		s := lang.spell(tt.in)

		if s != tt.out {
			t.Errorf("[TestSpellOrtho] Does not match: [%v] -> [%v] => expecting [%v]", tt.in, s, tt.out)
		}
	}
}

var vowelOrthoSpellTests = []struct {
	name string
	in   string
	out  string
}{
	{"Welsh", "AEIOU", "âêyôw"},
	{"Welsh", "aeiou", "aeiou"},
	{"Welsh", "ʧʔɣan", "ch‘ghan"},
	{"Doubles", "AEIOU", "aaeeiioouu"},
}

func TestSpellVowelOrtho(t *testing.T) {
	lang := OrthoLanguage()

	for _, tt := range vowelOrthoSpellTests {
		lang.ConsOrtho = vowelOrthSets[tt.name]

		s := lang.spell(tt.in)

		if s != tt.out {
			t.Errorf("[TestSpellVowelOrtho] Does not match: [%v] -> [%v] => expecting [%v]", tt.in, s, tt.out)
		}
	}
}

var consonantOrthoSpellTests = []struct {
	name string
	in   string
	out  string
}{
	{"Slavic", "AEIOU", "áéíóú"},
	{"Slavic", "ʧʔɣan", "č‘ghan"},
	{"French", "aeiou", "aeiou"},
	{"French", "ʤazzy ʃeese", "djazzy cheese"},
	{"Chinese", "ʧʔɣan", "q‘ghan"},
}

func TestSpellConsonantOrtho(t *testing.T) {
	lang := OrthoLanguage()

	for _, tt := range consonantOrthoSpellTests {
		lang.ConsOrtho = consonantOrthSets[tt.name]

		s := lang.spell(tt.in)

		if s != tt.out {
			t.Errorf("[TestSpellConsonantOrtho] Does not match: [%v] -> [%v] => expecting [%v]", tt.in, s, tt.out)
		}
	}
}

func TestMakeSyllable(t *testing.T) {
	lang := OrthoLanguage()

	syllable := lang.makeSyllable("CVC")

	fmt.Printf("Random syllable: %v\n", syllable)
}
