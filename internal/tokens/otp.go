package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
)

type Token struct {
	Secret string
	Hash   string
}

func GenerateOTP() (*Token, error) {
	generator, err := rand.Int(rand.Reader, big.NewInt(9000000))
	if err != nil {
		return nil, err
	}
	secret := generator.Int64() + 100000

	token := Token{
		Secret: fmt.Sprintf("%06d", secret),
	}

	hash := sha256.Sum256([]byte(token.Secret))
	token.Hash = fmt.Sprintf("%x\n", hash)

	return &token, nil
}

func FormatOTP(s string) string {
	length := len(s)
	half := length / 2
	first := s[:half]
	second := s[half:]
	words := []string{first, second}

	return strings.Join(words, " ")
}
