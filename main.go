package main

import (
	"log"
	"os"
	"strconv"
)

func usage() {
	log.Fatalf("\n\nUSAGE: %s SHIFT CLEAR_TEXT\n\n\tWhere SHIFT is positive number (right shift)\n\n", os.Args[0])
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	shift, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("SHIFT parameter must be positive number. Got %s. Error: %s", os.Args[1], err.Error())
	}
	if shift < 0 {
		shift = -shift
	}

	cipher := NewCaesarCipher(shift)

	input := os.Args[2]

	encrypted, err := cipher.Encrypt(input)
	if err != nil {
		log.Fatalf("Can't encrypt input \"%s\"; Error: %s", input, err.Error())
	}
	log.Printf("Input:   \"%s\"   ->   \"%s\"", input, encrypted)

	clearText, err := cipher.Decrypt(encrypted)
	if err != nil {
		log.Fatalf("Can't decrypt input \"%s\"; Error: %s", encrypted, err.Error())
	}
	log.Printf("Input:   \"%s\"   ->   \"%s\"", encrypted, clearText)
}
