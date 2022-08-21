package simhash

// a simhash lib
// compare words(chinese, english), pics
// use 128bit, we can

type Vector [128]int

// Vectorize generates 128 dimension vectors given a set of features.
// Vectors are initialized to zero. The i-th element of the vector is then
// incremented by weight of the i-th WeightedFeature if the i-th bit of the WeightedFeature
// is set, and decremented by the weight of the i-th WeightedFeature otherwise.
func Vectorize(features []Feature) Vector {
	var v Vector
	for _, feature := range features {
		sum := feature.Sum()
		weight := feature.Weight()
		for i := uint8(0); i < 128; i++ {
			j := i / 8
			k := i % 8
			bit := ((sum[j] >> (8 - k - 1)) & 1)
			if bit == 1 {
				v[i] += weight
			} else {
				v[i] -= weight
			}
		}
	}
	return v
}

// Fingerprint returns a 128-bit fingerprint of the given vector.
// The fingerprint f of a given 128-dimension vector v is defined as follows:
//   f[i] = 1 if v[i] >= 0
//   f[i] = 0 if v[i] < 0
func Fingerprint(v Vector) [16]byte {
	var f [16]byte
	for i := uint8(0); i < 128; i++ {
		j := i / 8
		k := i % 8
		if v[i] >= 0 {
			f[j] |= 1 << (8 - k - 1)
		}
	}
	return f
}

// Compare calculates the Hamming distance between two 64-bit integers
//
// Currently, this is calculated using the Kernighan method [1]. Other methods
// exist which may be more efficient and are worth exploring at some point
//
// [1] http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
func Compare(a, b [16]byte) uint8 {
	var c uint8
	c = 0
	for i := 0; i < 16; i++ {
		v := a[i] ^ b[i]
		for ; v != 0; c++ {
			v &= v - 1
		}
	}
	return c
}

// Returns a 64-bit simhash of the given WeightedFeature set
func Simhash(fs FeatureSet) [16]byte {
	return Fingerprint(Vectorize(fs.GetFeatures()))
}

func SimhashByteArray(bts []byte) [16]byte {
	fs := NewWordFeatureSet(bts)
	return Simhash(fs)
}
func SimhashString(str string) [16]byte {
	bts := []byte(str)
	return SimhashByteArray(bts)
}
