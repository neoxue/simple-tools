package simhash

import (
	"bytes"
	"regexp"
)

var Alg int = Shingled

const (
	Weighted int = iota
	Shingled
)

// FeatureSet represents a set of features in a given document
type FeatureSet interface {
	GetFeatures() []Feature
}

// Feature consists of a 64-bit hash and a weight
type Feature interface {
	// Sum returns the 64-bit sum of this WeightedFeature
	Sum() [16]byte
	// Weight returns the weight of this WeightedFeature
	Weight() int
}

// WordFeatureSet is a WeightedFeature set in which each word is a WeightedFeature,
// all equal weight.
type WordFeatureSet struct {
	b []byte
}

func NewWordFeatureSet(b []byte) *WordFeatureSet {
	fs := &WordFeatureSet{b}
	fs.normalize()
	return fs
}

func (w *WordFeatureSet) normalize() {
	w.b = bytes.ToLower(w.b)
}

//var boundaries = regexp.MustCompile(`[\w']+(?:\://[\w\./]+){0,1}`)
//var unicodeBoundaries = regexp.MustCompile(`[\pL-_']+`)
//add chinese characters, only basic character, 20902 characters
//https://www.qqxiuzi.cn/zh/hanzi-unicode-bianma.php
var unicodeBoundaries = regexp.MustCompile("[\u4e00-\u9fa5]|[a-zA-Z0-9]+")

// Returns a []Feature representing each word in the byte slice
//
func (w *WordFeatureSet) GetFeatures() []Feature {
	switch Alg {
	case Weighted:
		return getWeightedFeatures(w.b, unicodeBoundaries)
	case Shingled:
		return getShingledFeatures(w.b, unicodeBoundaries)
	default:
		return getShingledFeatures(w.b, unicodeBoundaries)
	}
}
