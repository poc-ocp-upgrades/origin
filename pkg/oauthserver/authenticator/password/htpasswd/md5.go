package htpasswd

import "crypto/md5"

const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var wordOutputs = [][][2]int{{{0, 16}, {6, 8}, {12, 0}}, {{1, 16}, {7, 8}, {13, 0}}, {{2, 16}, {8, 8}, {14, 0}}, {{3, 16}, {9, 8}, {15, 0}}, {{4, 16}, {10, 8}, {5, 0}}, {{11, 0}}}
var magic = []byte("$apr1$")

func aprMD5(password, salt []byte) []byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx := md5.New()
	ctx.Write(password)
	ctx.Write(magic)
	ctx.Write(salt)
	ctx1 := md5.New()
	ctx1.Write(password)
	ctx1.Write(salt)
	ctx1.Write(password)
	final := ctx1.Sum(nil)
	for i := len(password); i > 0; i -= md5.Size {
		if i > md5.Size {
			ctx.Write(final)
		} else {
			ctx.Write(final[:i])
		}
	}
	for i := len(password); i != 0; i >>= 1 {
		if i&1 != 0 {
			ctx.Write([]byte{0})
		} else {
			ctx.Write([]byte{password[0]})
		}
	}
	final = ctx.Sum(nil)
	for i := 0; i < 1000; i++ {
		ctx1 := md5.New()
		if i&1 != 0 {
			ctx1.Write(password)
		} else {
			ctx1.Write(final)
		}
		if i%3 != 0 {
			ctx1.Write(salt)
		}
		if i%7 != 0 {
			ctx1.Write(password)
		}
		if i&1 != 0 {
			ctx1.Write(final)
		} else {
			ctx1.Write(password)
		}
		final = ctx1.Sum(nil)
	}
	result := []byte{}
	result = append(result, magic...)
	result = append(result, salt...)
	result = append(result, '$')
	for _, word := range wordOutputs {
		l := uint64(0)
		for _, chunk := range word {
			index := chunk[0]
			offset := chunk[1]
			l |= (uint64(final[index]) << uint(offset))
		}
		result = append(result, to64(l, len(word)+1)...)
	}
	return result
}
func to64(v uint64, n int) []byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := make([]byte, n)
	for i := 0; i < n; i++ {
		r[i] = itoa64[v&0x3f]
		v >>= 6
	}
	return r
}
