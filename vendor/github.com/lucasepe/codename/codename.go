package codename

import (
	"bytes"
	"fmt"
	"math/rand"

	crypto_rand "crypto/rand"
	"encoding/binary"
)

// NewCryptoSeed returns a crypto level random numbers generator seed.
// It returns an error and a seed equals to -1 if the underlying system call fails.
func NewCryptoSeed() (int64, error) {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return -1, fmt.Errorf("cannot seed math/rand package with Crypto RNG: %w", err)
	}

	seed := int64(binary.LittleEndian.Uint64(b[:]))
	return seed, nil
}

// DefaultRNG returns a new pseudo-random source
// seeded by a crypto level seed.
func DefaultRNG() (*rand.Rand, error) {
	seed, err := NewCryptoSeed()
	if err != nil {
		return nil, err
	}

	rng := rand.New(rand.NewSource(seed))
	return rng, nil
}

// Generate generates and returns a random hero name.
// Eventually you can specify a `tokenLength` greater
// then zero to generate and additional token and create
// even more entropy.
func Generate(rng *rand.Rand, tokenLength int) string {
	res := fmt.Sprintf("%s-%s", randomAdjective(rng), randomNoun(rng))
	if tokenLength > 0 {
		res = fmt.Sprintf("%s-%s", res, randomToken(rng, tokenLength))
	}
	return res
}

// randomAdjective returns a random adjective from a list of adjectives.
func randomAdjective(rng *rand.Rand) string {
	return adjectives[rng.Intn(len(adjectives))]
}

// randomNoun returns a random noun from a list of nouns.
func randomNoun(rng *rand.Rand) string {
	return nouns[rng.Intn(len(nouns))]
}

// randomToken creates and builds random token
func randomToken(rng *rand.Rand, size int) string {
	hex := []byte{
		'0', '1', '2', '3',
		'4', '5', '6', '7',
		'8', '9', 'a', 'b',
		'c', 'd', 'e', 'f',
	}

	var buffer bytes.Buffer

	for i := 0; i < size; i++ {
		index := rng.Intn(len(hex))
		buffer.WriteByte(hex[index])
	}

	return buffer.String()
}
