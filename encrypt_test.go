package anet

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	src := "hello world2021dlrow olleh"
	enc, err := Encrypt([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := Decrypt(enc)
	if err != nil {
		t.Fatal(err)
	}
	if string(dec) != src {
		t.Fatalf("%s", string(dec))
	}
}
