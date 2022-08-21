package simhash

import (
	"fmt"
	"testing"
)

func TestWeightBytes(t *testing.T) {
	bts := []byte("test a ")
	if weightBytes(bts) != 5 {
		t.Error("weighted btes english should be 5")
	}
	bts = []byte("的")
	if weightBytes(bts) != 0 {
		t.Error("weighted btes 的 should be 0")
	}
	bts = []byte("一")
	if weightBytes(bts) != 1 {
		t.Error("weighted btes 一 should be 1")
	}
	bts = []byte("也")
	if weightBytes(bts) != 2 {
		t.Error("weighted btes 也 should be 2")
	}
	bts = []byte("铩")
	if weightBytes(bts) != 5 {
		t.Error("weighted btes 铩 should be 5")
	}
	bts = []byte("an")
	if weightBytes(bts) != 1 {
		t.Error("weighted btes 'an' should be 1")
	}
}
func TestNewWeightedFeature(t *testing.T) {
	bts := []byte("我")
	f := newWeightedFeature(bts)
	fmt.Println(f.Sum())
	fmt.Println(f.Weight())
	bts2 := []byte("们")
	f2 := newWeightedFeature(bts2)
	fmt.Println(f2.Sum())
	fmt.Println(f2.Weight())
}

func TestGetWeightedFeatures(t *testing.T) {
	bts := []byte("我们来自北京")
	fs := getWeightedFeatures(bts, unicodeBoundaries)
	bts2 := []byte("来自北京我们")
	fs2 := getWeightedFeatures(bts2, unicodeBoundaries)

	f1 := Fingerprint(Vectorize(fs))
	f2 := Fingerprint(Vectorize(fs2))
	a := Compare(f1, f2)
	if a > 1 {
		t.Error("weighted features should be the same , 我们来自北京/来自北京我们")
	}

}
