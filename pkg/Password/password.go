package Password

import (
	"github.com/alexedwards/argon2id"
)

type Password struct {
	Plaintext *string
	Hash      string
}

// Set encrypt and hashes the input password
func (p *Password) Set(plaintextPassword string) error {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

// Matches compares the existing hash with the newly hashed password to see if they match
func (p *Password) Matches(plaintextPassword string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plaintextPassword, p.Hash)
}
