package hmac

// llgo:skipall
import (
	"crypto/subtle"
	"hash"

	"github.com/goplus/llgo/internal/lib/crypto/internal/boring"
)

type marshalable interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

type hmac struct {
	opad, ipad   []byte
	outer, inner hash.Hash

	// If marshaled is true, then opad and ipad do not contain a padded
	// copy of the key, but rather the marshaled state of outer/inner after
	// opad/ipad has been fed into it.
	marshaled bool
}

func (h *hmac) Sum(in []byte) []byte {
	origLen := len(in)
	in = h.inner.Sum(in)

	if h.marshaled {
		if err := h.outer.(marshalable).UnmarshalBinary(h.opad); err != nil {
			panic(err)
		}
	} else {
		h.outer.Reset()
		h.outer.Write(h.opad)
	}
	h.outer.Write(in[origLen:])
	return h.outer.Sum(in[:origLen])
}

func (h *hmac) Write(p []byte) (n int, err error) {
	return h.inner.Write(p)
}

func (h *hmac) Size() int      { return h.outer.Size() }
func (h *hmac) BlockSize() int { return h.inner.BlockSize() }

func (h *hmac) Reset() {
	if h.marshaled {
		if err := h.inner.(marshalable).UnmarshalBinary(h.ipad); err != nil {
			panic(err)
		}
		return
	}

	h.inner.Reset()
	h.inner.Write(h.ipad)

	// If the underlying hash is marshalable, we can save some time by
	// saving a copy of the hash state now, and restoring it on future
	// calls to Reset and Sum instead of writing ipad/opad every time.
	//
	// If either hash is unmarshalable for whatever reason,
	// it's safe to bail out here.
	marshalableInner, innerOK := h.inner.(marshalable)
	if !innerOK {
		return
	}
	marshalableOuter, outerOK := h.outer.(marshalable)
	if !outerOK {
		return
	}

	imarshal, err := marshalableInner.MarshalBinary()
	if err != nil {
		return
	}

	h.outer.Reset()
	h.outer.Write(h.opad)
	omarshal, err := marshalableOuter.MarshalBinary()
	if err != nil {
		return
	}

	// Marshaling succeeded; save the marshaled state for later
	h.ipad = imarshal
	h.opad = omarshal
	h.marshaled = true
}

// New returns a new HMAC hash using the given [hash.Hash] type and key.
// New functions like sha256.New from [crypto/sha256] can be used as h.
// h must return a new Hash every time it is called.
// Note that unlike other hash implementations in the standard library,
// the returned Hash does not implement [encoding.BinaryMarshaler]
// or [encoding.BinaryUnmarshaler].
func New(h func() hash.Hash, key []byte) hash.Hash {
	if boring.Enabled {
		hm := boring.NewHMAC(h, key)
		if hm != nil {
			return hm
		}
		// BoringCrypto did not recognize h, so fall through to standard Go code.
	}

	hm := new(hmac)
	hm.outer = h()
	hm.inner = h()
	unique := true
	func() {
		defer func() {
			// The comparison might panic if the underlying types are not comparable.
			_ = recover()
		}()
		if hm.outer == hm.inner {
			unique = false
		}
	}()
	if !unique {
		panic("crypto/hmac: hash generation function does not produce unique values")
	}
	blocksize := hm.inner.BlockSize()
	hm.ipad = make([]byte, blocksize)
	hm.opad = make([]byte, blocksize)
	if len(key) > blocksize {
		// If key is too big, hash it.
		hm.outer.Write(key)
		key = hm.outer.Sum(nil)
	}
	copy(hm.ipad, key)
	copy(hm.opad, key)
	for i := range hm.ipad {
		hm.ipad[i] ^= 0x36
	}
	for i := range hm.opad {
		hm.opad[i] ^= 0x5c
	}
	hm.inner.Write(hm.ipad)
	return hm
}

// Equal compares two MACs for equality without leaking timing information.
func Equal(mac1, mac2 []byte) bool {
	// We don't have to be constant time if the lengths of the MACs are
	// different as that suggests that a completely different hash function
	// was used.
	return subtle.ConstantTimeCompare(mac1, mac2) == 1
}
