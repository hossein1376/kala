package password

import (
	"errors"

	"github.com/alexedwards/argon2id"
)

type Password struct {
	Plaintext *string
	Hash      []byte
}

// ArgonSet encrypt and hashes the input password
func (p *Password) ArgonSet(plaintextPassword string) error {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = []byte(hash)
	return nil
}

// ArgonMatches compares the existing hash with the newly hashed password to see if they match
func (p *Password) ArgonMatches() (bool, error) {
	if p.Plaintext == nil {
		return false, errors.New("'plaintext' field must be provided")
	}
	return argon2id.ComparePasswordAndHash(*p.Plaintext, string(p.Hash))
}
