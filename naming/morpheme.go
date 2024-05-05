package naming

func (lang *Language) makeMorpheme(structure string, group string) string {
	if !lang.ApplyMorph {
		return lang.makeSyllable(structure)
	}

	// Get the morpheme list for the given group.
	var list []string
	if val, ok := lang.Morphemes[group]; ok {
		list = val
	}

	// Extras defines the chance of a morpheme being a new
	// morpheme, rather than a morpheme from the existing list.
	extras := 10
	if len(group) > 0 {
		extras = 1
	}

	for {
		// If random range returns a number larger than
		// the length of the current morpheme list.
		n := lang.Rnd.Intn(len(list) + extras)
		if n < len(list) {
			return list[n] // Return a morpheme from the list.
		}

		morph := lang.makeSyllable(structure)

		// Check if the generated morpheme is already present
		// in any group. We iterate in stable order over
		// the group keys.
		var exists bool
		for _, k := range sortedKeys(lang.Morphemes) {
			if contains(lang.Morphemes[k], morph) {
				exists = true
				break
			}
		}

		// If the morpheme already exists, try again.
		if exists {
			continue
		}

		// Since it is a new morpheme, just add it to the list for the
		// current group.
		lang.Morphemes[group] = append(list, morph)
		return morph
	}
}
