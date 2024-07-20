package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"strings"
)

func calculatePasswordEntropy(password string) float64 {
	if isHexString(password) {
		return float64(len(password) * 4)
	}

	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	symbols := "!@#$%^&*()_+-=[]{}|;:,.<>"

	usedLowercase := strings.ContainsAny(password, lowercase)
	usedUppercase := strings.ContainsAny(password, uppercase)
	usedDigits := strings.ContainsAny(password, digits)
	usedSymbols := strings.ContainsAny(password, symbols)

	charsetSize := 0
	if usedLowercase {
		charsetSize += 26
	}
	if usedUppercase {
		charsetSize += 26
	}
	if usedDigits {
		charsetSize += 10
	}
	if usedSymbols {
		charsetSize += 26
	}

	if charsetSize == 0 {
		return 0
	}
	return math.Log2(float64(charsetSize)) * float64(len(password))
}

func isHexString(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil && len(s)%2 == 0
}

func printUsage() {
	fmt.Println("Usage: go run entropy.go <password>")
	fmt.Println("Example: go run entropy.go 'MyPassword123!'")
	fmt.Println("For hex strings: go run entropy.go '1a2b3c4d5e6f'")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	password := os.Args[1]
	entropy := calculatePasswordEntropy(password)
	
	fmt.Printf("Password: %s\n", password)
	if isHexString(password) {
		fmt.Println("Detected as hex string")
	}
	fmt.Printf("Entropy: %.2f bits\n", entropy)
}
