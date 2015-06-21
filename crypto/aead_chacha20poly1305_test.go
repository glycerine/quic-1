package crypto

import "testing"
import "bytes"

func Test_KeyGeneratorPoly1305(t *testing.T) {
	var buf [64]byte
	var cipher *ChaCha20Cipher
	var err error

	key := []byte{0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
		0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f}
	noncePrefix := []byte{0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7}

	if cipher, err = NewChaCha20Cipher(key, noncePrefix, 0); err != nil {
		t.Error("Key Generator test for Poly1305 : error when calling NewChaCha20Cipher")
	}
	cipher.GetNextKeystream(&buf)

	test := []byte{
		0x8a, 0xd5, 0xa0, 0x8b, 0x90, 0x5f, 0x81, 0xcc, 0x81, 0x50, 0x40, 0x27, 0x4a, 0xb2, 0x94, 0x71,
		0xa8, 0x33, 0xb6, 0x37, 0xe3, 0xfd, 0x0d, 0xa5, 0x08, 0xdb, 0xb8, 0xe2, 0xfd, 0xd1, 0xa6, 0x46}
	if !bytes.Equal(buf[:32], test) {
		t.Error("NewAEAD_ChaCha20Poly1305 : bad Poly1305 Key Generation test vector")
	}
}
