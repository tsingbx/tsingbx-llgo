package boring

import (
	"bytes"
	"hash"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/openssl"
	"github.com/goplus/llgo/internal/lib/crypto/sha1"
	"github.com/goplus/llgo/internal/lib/crypto/sha256"
	"github.com/goplus/llgo/internal/lib/crypto/sha512"
)

func hashToMD(h hash.Hash) *openssl.EVP_MD {
	md := sha1.HashToMD(h)
	if md != nil {
		return md
	}
	md = sha256.HashToMD(h)
	if md != nil {
		return md
	}
	return sha512.HashToMD(h)
}

func NewHMAC(h func() hash.Hash, key []byte) hash.Hash {
	ch := h()
	md := hashToMD(ch)
	if md == nil {
		return nil
	}
	// Note: Could hash down long keys here using EVP_Digest.
	hkey := bytes.Clone(key)
	hmac := &boringHMAC{
		md:        md,
		size:      ch.Size(),
		blockSize: ch.BlockSize(),
		key:       hkey,
	}
	hmac.Reset()
	return hmac
}

type boringHMAC struct {
	ctx       *openssl.HMAC_CTX
	size      int
	blockSize int
	key       []byte
	md        *openssl.EVP_MD
	sum       []byte
}

func (h *boringHMAC) Reset() {
	h.ctx.Reset()
}

func (h *boringHMAC) Finalize() {
	h.ctx.Free()
}

func (h *boringHMAC) Write(p []byte) (int, error) {
	return int(h.ctx.UpdateBytes(p)), nil
}

func (h *boringHMAC) Size() int {
	return h.size
}

func (h *boringHMAC) BlockSize() int {
	return h.blockSize
}

func (h *boringHMAC) Sum(in []byte) []byte {
	var digestLen c.Uint = c.Uint(h.Size())
	h.sum = make([]byte, digestLen)
	ctx := openssl.NewHMAC_CTX()
	if ctx == nil {
		c.Fprintf(c.Stderr, c.Str("%s\n"), c.Str("Error creating HMAC_CTX"))
		return nil
	}
	defer ctx.Free()
	var lenKey c.Int = c.Int(len(h.key))
	var ret c.Int = ctx.InitEx(unsafe.Pointer(unsafe.SliceData(h.key)), lenKey, openssl.EVP_sha256(), nil)
	if ret == 0 {
		c.Fprintf(c.Stderr, c.Str("%s\n"), c.Str("Error initializing HMAC_CTX"))
		return nil
	}
	h.ctx.InitEx(unsafe.Pointer(unsafe.SliceData(h.key)), c.Int(len(h.key)), h.md, nil)
	h.ctx.UpdateBytes(in)
	h.ctx.Final(unsafe.SliceData(h.sum), &digestLen)
	return append(in, h.sum...)
}
