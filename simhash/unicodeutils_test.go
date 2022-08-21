package simhash

import (
	"strconv"
	"testing"
)

func TestMd5sumUnicodeByte(t *testing.T) {
	f := []byte("我们")
	k := md5sumUnicodeByte(f)
	if len(k) != 16 {
		t.Errorf("expected byte length 16, actual %d", len(k))
	}
	i := 0
	for true {
		i++
		n := []byte(strconv.Itoa(i))
		go md5sumUnicodeByte(n)
		if i > 100000 {
			break
		}
	}

}
