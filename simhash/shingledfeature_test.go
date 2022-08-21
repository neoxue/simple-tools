package simhash

import (
	"testing"
)

func TestNewShingledFeature(t *testing.T) {
	w := []byte("test")
	wlast := []byte("我们")
	a := newShingledFeature(w, wlast)
	b := newShingledFeature(w, wlast)
	if a.Sum() != b.Sum() {
		t.Error("a,b sum not equal")
	}
}
func TestGetShingledFeatures(t *testing.T) {
	stringa := "南京，   超级amazing. 超666!"
	words := [][]byte{[]byte("南"), []byte("京"), []byte("超"), []byte("级"), []byte("amazing"), []byte("超"), []byte("666")}

	features := []Feature{}
	wlast := []byte{}
	for _, w := range words {
		features = append(features, newShingledFeature(w, wlast))
		wlast = w
	}
	retfetures := getShingledFeatures([]byte(stringa), unicodeBoundaries)
	if len(features) != len(retfetures) {
		t.Error("shingled features retfeatures not equal")
	}
	for i, f := range features {
		if f.Sum() != retfetures[i].Sum() {
			t.Error("shingled features retfeatures not equal")
		}
	}
}
