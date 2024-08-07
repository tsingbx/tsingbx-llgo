package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	h := hmac.New(sha256.New, []byte{
		0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
		0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
		0x0b, 0x0b, 0x0b, 0x0b,
	})

	if s := h.Size(); s != sha256.Size {
		fmt.Printf("Size: got %v, want %v\n", s, sha256.Size)
	}
	if b := h.BlockSize(); b != sha256.BlockSize {
		fmt.Printf("BlockSize: got %v, want %v\n", b, sha256.BlockSize)
	}
	in := []byte("Hi There")
	out := "b0344c61d8db38535ca8afceaf0bf12b881dc200c9833da726e9376c2e32cff7"
	buf := make([]byte, sha256.Size)
	n, err := h.Write(in)
	if n != len(in) || err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for k := 0; k < 2; k++ {
		sum := fmt.Sprintf("%x", h.Sum(buf))
		if sum != out {
			fmt.Printf("err: %s want %s\n", sum[2*len(buf):], out)
		}
	}
}
