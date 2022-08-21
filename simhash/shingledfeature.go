package simhash

import (
	"hash/fnv"
	"regexp"
)

// feature that shingled words

// shingledfeature implements Feature type
type ShingledFeature struct {
	sum    [16]byte
	weight int
}

func (f ShingledFeature) Sum() [16]byte {
	return f.sum
}
func (f ShingledFeature) Weight() int {
	return f.weight
}

// Splits the given []byte using the given regexp, then returns a slice
// containing a Feature constructed from each piece matched by the regexp
func getShingledFeatures(b []byte, r *regexp.Regexp) []Feature {
	//r = regexp.MustCompile("[\u4e00-\u9fa5]|[a-zA-Z0-9]+")
	words := r.FindAll(b, -1)
	//TODO should make it more solid
	//features := make([]Feature, len(words))
	features := []Feature{}
	wlast := []byte{}
	for _, w := range words {
		if len(w) < 1 {
			continue
		}
		features = append(features, newShingledFeature(w, wlast))
		wlast = w
	}
	return features
}

func newShingledFeature(w, wlast []byte) ShingledFeature {
	weight := 1
	if unicodeBoundaries.Match(w) {
		w = md5sumUnicodeByte(w)
	}
	if unicodeBoundaries.Match(wlast) {
		wlast = md5sumUnicodeByte(wlast)
	}
	h := fnv.New128()
	h.Write(append(w, wlast...))
	b := []byte{}
	b = h.Sum(b)
	c := [16]byte{}
	for index, bi := range b {
		c[index] = bi
	}
	return ShingledFeature{c, weight}
}
