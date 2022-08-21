package simhash

import (
	"crypto/md5"
	"sync"
)

//unicode utils
var mu sync.RWMutex
var unicodewordkv = map[string][]byte{}

func md5sumUnicodeByte(f []byte) []byte {
	str := string(f)
	mu.RLock()
	k, ok := unicodewordkv[str]
	mu.RUnlock()
	if ok {
		return k
	}
	hasher := md5.New()
	hasher.Write(f)
	g := hasher.Sum(nil)
	hasher.Reset()
	mu.Lock()
	unicodewordkv[str] = g
	mu.Unlock()
	return g
}
