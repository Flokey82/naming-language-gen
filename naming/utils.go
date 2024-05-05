package naming

import (
	"math"
	"math/rand"
	"sort"
	"unicode/utf8"
)

func RandomRange(min, max int, rnd *rand.Rand) int {
	return rnd.Intn(max-min) + min
}

func RandomRuneFromString(str string, rnd *rand.Rand) string {
	n := utf8.RuneCountInString(str)
	r := []rune(str)[rnd.Intn(n)]
	return string(r)
}

func contains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}

	return false
}

// randomItem returns a random item from a map that uses string keys.
func randomItem[V any](m map[string]V, rnd *rand.Rand) V {
	keys := sortedKeys[V](m)
	i := RandomRange(0, len(keys)-1, rnd)
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

func randFloat64Abs(rnd *rand.Rand) float64 {
	return math.Abs(rnd.NormFloat64())
}
