package main

import (
	"strings"
	"unicode"
)

type Cipher interface {
	Encrypt(clearText string) (cipherText string, err error)
	Decrypt(cipherText string) (clearText string, err error)
}

const (
	cipherAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetBase   = int('A')
)

var (
	alphabetLen = len(cipherAlphabet)
)

type caesarCipher struct {
	shift          int
	alphabet       []rune
	rotatedChToPos map[rune]int
}

func NewCaesarCipher(shift int) Cipher {
	cc := &caesarCipher{
		shift:    shift,
		alphabet: []rune(cipherAlphabet),
	}

	// if cc.shift < 0 {
	// 	cc.shift = -cc.shift
	// }

	abl := len(cc.alphabet)
	if abl < cc.shift {
		cc.shift %= abl
	}

	cc.rightRotate()

	cc.rotatedChToPos = make(map[rune]int)
	for i, r := range cc.alphabet {
		cc.rotatedChToPos[r] = i
	}

	return cc
}

// func (cc *caesarCipher) leftRotate() {
// 	var res []rune
// 	lft := cc.alphabet[0:cc.shift]
// 	res = append(res, cc.alphabet[cc.shift:]...)
// 	res = append(res, lft...)
// 	cc.alphabet = res
// }

func (cc *caesarCipher) rightRotate() {
	var res []rune
	rght := cc.alphabet[len(cc.alphabet)-cc.shift:]
	res = append(res, rght...)
	res = append(res, cc.alphabet[:len(cc.alphabet)-cc.shift]...)
	cc.alphabet = res
}

func (cc *caesarCipher) Encrypt(clearText string) (cipherText string, err error) {
	// log.Printf("Rotated alphabet: %s", string(cc.alphabet))
	cipher := make([]rune, len(clearText))
	for i, ch := range []rune(strings.ToUpper(clearText)) {
		if !unicode.IsLetter(ch) {
			cipher[i] = ch
		} else {
			chb := int(ch) - alphabetBase
			cipher[i] = cc.alphabet[chb%alphabetLen]
		}
	}
	// log.Printf("Encrypted: \"%s\"", string(cipher))
	return string(cipher), err
}

func (cc *caesarCipher) Decrypt(cipherText string) (clearText string, err error) {
	clear := make([]rune, len(cipherText))

	ab := []rune(cipherAlphabet)
	for i, ch := range []rune(strings.ToUpper(cipherText)) {
		if !unicode.IsLetter(ch) {
			clear[i] = ch
		} else {
			chb := cc.rotatedChToPos[ch]
			clear[i] = ab[chb]
		}
	}
	// log.Printf("Decrypted: \"%s\"", string(clear))
	return string(clear), err
}
