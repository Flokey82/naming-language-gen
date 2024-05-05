package naming

import "math/rand"

type structureList []string

func (list *structureList) random(rnd *rand.Rand) string {
	l := len(*list)
	i := rnd.Intn(l)
	return (*list)[i]
}

type WordParams struct {
	MinSyllables int
	MaxSyllables int
	Structure    structureList
}

// Clone returns a deep copy of the WordParams.
func (p *WordParams) Clone() *WordParams {
	clone := &WordParams{
		MinSyllables: p.MinSyllables,
		MaxSyllables: p.MaxSyllables,
		Structure:    make(structureList, len(p.Structure)),
	}
	copy(clone.Structure, p.Structure)
	return clone
}

// makeWord generates a new random word based on the given parameters.
func (lang Language) makeWord(p *WordParams, group string) (word string) {
	numSyllables := RandomRange(p.MinSyllables, p.MaxSyllables+1, lang.Rnd)

	keys := make([]string, numSyllables)
	keys[RandomRange(0, numSyllables, lang.Rnd)] = group

	for i := 0; i < numSyllables; i++ {
		word += lang.makeMorpheme(p.Structure.random(lang.Rnd), keys[i])
	}
	return
}

func (lang *Language) GetWord(p *WordParams, group string) (word string) {
	words := []string{}

	// Get the word list for the given group.
	if val, ok := lang.Words.General[group]; ok {
		words = val
	}

	// Extras defines the chance of a word being a new
	// word, rather than a word from the existing list.
	extras := 4
	if len(group) > 0 {
		extras = 3
	}

	for {
		// If random range returns a number larger than
		// the length of the current word list.
		n := RandomRange(0, len(words)+extras, lang.Rnd)
		if n < len(words) {
			return words[n] // Return a word from the list.
		}

		word = lang.makeWord(p, group)
		exists := false

		// Check if the generated word is already present
		// in any group. We iterate in stable order over
		// the group keys.
		keys := sortedKeys(lang.Words.General)
		for _, k := range keys {
			v := lang.Words.General[k]
			if contains(v, word) {
				exists = true
				break
			}
		}

		// If the word already exists, try again.
		if exists {
			continue
		}

		// Since it is a new word, just add it to the list for the
		// current group.
		lang.Words.General[group] = append(words, word)
		return word
	}
}
