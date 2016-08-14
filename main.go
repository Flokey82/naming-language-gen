package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
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

	rand.Seed(time.Now().UTC().UnixNano())

	lang := RandomLanguage(true, true)

	// generate some words based on our new awesome language
	p := wordParams{
		1,
		randomRange(1, 5),
		defaultSyllableStructures,
	}

	group := "words"

	for i := 0; i < 25; i++ {
		lang.getWord(p, group)
	}

	wordList := strings.Join(lang.Words.General[group][:], ", ")
	fmt.Printf("[%v]: %v\n", group, wordList)

	// generate some names
	p = wordParams{
		2,
		randomRange(2, 7),
		defaultSyllableStructures,
	}

	const joiners = "  -"

	minLength := randomRange(3, 5)
	maxLength := randomRange(6, 20)

	for i := 0; i < 20; i++ {
		lang.makeName(minLength, maxLength, p, joiners, "words")
	}

	nameList := strings.Join(lang.Words.Names[:], ", ")
	fmt.Printf("[%v]: %v\n", "names", nameList)

	lang.Describe()
}
