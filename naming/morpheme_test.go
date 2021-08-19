package naming

import (
	"fmt"
	"testing"
)

func TestGetMorphemeWithMorph(t *testing.T) {
	lang := OrthoLanguage()

	lang.ApplyMorph = true

	groups := []string{"", "1", "2"}

	for _, group := range groups {
		list := []string{}

		for i := 0; i < 20; i++ {
			list = append(list, lang.makeMorpheme("CVC", group))
		}

		fmt.Printf("Random morpheme [%v]: %v\n", group, list)
	}
}
