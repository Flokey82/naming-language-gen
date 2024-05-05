package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Flokey82/naming-language-gen/naming"
)

func main() {
	var (
		help = flag.Bool("help", false, "Displays helpful usage info")
	)

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	lang := naming.RandomLanguage(true, true, time.Now().UTC().UnixNano())

	// generate some words based on our new awesome language
	p := &naming.WordParams{
		MinSyllables: 1,
		MaxSyllables: naming.RandomRange(1, 5, lang.Rnd),
		Structure:    naming.DefaultSyllableStructures,
	}

	group := "words"

	for i := 0; i < 25; i++ {
		lang.GetWord(p, group)
	}

	wordList := strings.Join(lang.Words.General[group][:], ", ")
	fmt.Printf("[%v]: %v\n", group, wordList)

	// generate some names
	p = &naming.WordParams{
		MinSyllables: 2,
		MaxSyllables: naming.RandomRange(2, 7, lang.Rnd),
		Structure:    naming.DefaultSyllableStructures,
	}

	params := naming.NameParams{
		MinLength:  naming.RandomRange(3, 5, lang.Rnd),
		MaxLength:  naming.RandomRange(6, 20, lang.Rnd),
		WordParams: p,
		Joiners:    "  -",
		Group:      "words",
	}

	for i := 0; i < 20; i++ {
		lang.MakeName(&params)
	}

	nameList := strings.Join(lang.Words.Names[:], ", ")
	fmt.Printf("[%v]: %v\n", "names", nameList)

	lang.Describe()
}
