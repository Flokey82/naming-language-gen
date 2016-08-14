package main

type orthoMapping map[string]string

var defaultOrtho = orthoMapping{
	"ʃ": "sh",
	"ʒ": "zh",
	"ʧ": "ch",
	"ʤ": "j",
	"ŋ": "ng",
	"j": "y",
	"x": "kh",
	"ɣ": "gh",
	"ʔ": "‘",
	"A": "á",
	"E": "é",
	"I": "í",
	"O": "ó",
	"U": "ú",
}

type orthoSet map[string]orthoMapping

func (o *orthoSet) random() orthoMapping {
	i := randomRange(1, len(*o))

	for _, val := range *o {
		i--
		if i == 0 {
			return val
		}
	}

	return orthoMapping{}
}

var vowelOrthSets = orthoSet{
	"Ácutes": orthoMapping{
		"A": "á",
		"E": "é",
		"I": "í",
		"O": "ó",
		"U": "ú",
	},
	"Ümlauts": orthoMapping{
		"A": "ä",
		"E": "ë",
		"I": "ï",
		"O": "ö",
		"U": "ü",
	},
	"Welsh": orthoMapping{
		"A": "â",
		"E": "ê",
		"I": "y",
		"O": "ô",
		"U": "w",
	},
	"Diphthongs": orthoMapping{
		"A": "au",
		"E": "ei",
		"I": "ie",
		"O": "ou",
		"U": "oo",
	},
	"Doubles": orthoMapping{
		"A": "aa",
		"E": "ee",
		"I": "ii",
		"O": "oo",
		"U": "uu",
	},
}

var consonantOrthSets = orthoSet{
	"Slavic": orthoMapping{
		"ʃ": "š",
		"ʒ": "ž",
		"ʧ": "č",
		"ʤ": "ǧ",
		"j": "j",
	},
	"German": orthoMapping{
		"ʃ": "sch",
		"ʒ": "zh",
		"ʧ": "tsch",
		"ʤ": "dz",
		"j": "j",
		"x": "ch",
	},
	"French": orthoMapping{
		"ʃ": "ch",
		"ʒ": "j",
		"ʧ": "tch",
		"ʤ": "dj",
		"x": "kh",
	},
	"Chinese": orthoMapping{
		"ʃ": "x",
		"ʧ": "q",
		"ʤ": "j",
	},
}

type consonantSet map[string]string

func (c *consonantSet) random() string {
	i := randomRange(1, len(*c))

	for _, val := range *c {
		i--
		if i == 0 {
			return val
		}
	}

	return ""
}

var consonantSets = consonantSet{
	"Minimal":              "ptkmnls",
	"English-ish":          "ptkbdgmnlrsʃzʒʧ",
	"Pirahã (very simple)": "ptkmnh",
	"Hawaiian-ish":         "hklmnpwʔ",
	"Greenlandic-ish":      "ptkqvsgrmnŋlj",
	"Arabic-ish":           "tksʃdbqɣxmnlrwj",
	"Arabic-lite":          "tkdgmnsʃ",
	"English-lite":         "ptkbdgmnszʒʧhjw",
}

type vowelSet map[string]string

func (v *vowelSet) random() string {
	i := randomRange(1, len(*v))

	for _, val := range *v {
		i--
		if i == 0 {
			return val
		}
	}

	return ""
}

var vowelSets = vowelSet{
	"Standard 5-vowel":  "aeiou",
	"3-vowel a i u":     "aiu",
	"Extra A E I":       "aeiouAEI",
	"Extra U":           "aeiouU",
	"5-vowel a i u A I": "aiuAI",
	"3-vowel e o u":     "eou",
	"Extra A O U":       "aeiouAOU",
}

type phonemeSet map[string]string

func (p *phonemeSet) random() string {
	i := randomRange(1, len(*p))

	for _, val := range *p {
		i--
		if i == 0 {
			return val
		}
	}

	return ""
}

var phonemeSSets = phonemeSet{
	"Just s": "s",
	"s ʃ":    "sʃ",
	"s ʃ f":  "sʃf",
}

var phonemeFSets = phonemeSet{
	"m n":     "mn",
	"s k":     "sk",
	"m n ŋ":   "mnŋ",
	"s ʃ z ʒ": "sʃzʒ",
}

var phonemeLSets = phonemeSet{
	"r l":     "rl",
	"Just r":  "r",
	"Just l":  "l",
	"w j":     "wj",
	"r l w j": "rlwj",
}

var defaultSyllableStructures = structureList{
	"CVC",
	"CVV?C",
	"CVVC?", "CVC?", "CV", "VC", "CVF", "C?VC", "CVF?",
	"CL?VC", "CL?VF", "S?CVC", "S?CVF", "S?CVC?",
	"C?VF", "C?VC?", "C?VF?", "C?L?VC", "VC",
	"CVL?C?", "C?VL?C", "C?VLC?",
}

type restrictionSet map[string][]string

func (r *restrictionSet) random() []string {
	i := randomRange(1, len(*r))

	for _, val := range *r {
		i--
		if i == 0 {
			return val
		}
	}

	return []string{}
}

var restrictionSets = restrictionSet{
	"None": []string{},
	// "Double sounds": []string{"/(.)\\1/"},  // backreferences do not work with this regexp lib
	"Hard clusters": []string{"[sʃf][sʃ]", "[rl][rl]"},
	// "Doubles and hard clusters": []string{"[sʃf][sʃ]", "/(.)\\1/", "[rl][rl]"}, // backreferences do not work with this regexp lib
}
