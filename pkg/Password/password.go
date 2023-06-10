package Password

import (
	"errors"

	"github.com/alexedwards/argon2id"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Plaintext *string
	Hash      string
}

// ArgonSet encrypt and hashes the input password
func (p *Password) ArgonSet(plaintextPassword string) error {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

// BcryptSet encrypt and hashes the input password
func (p *Password) BcryptSet(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = string(hash)
	return nil
}

// ArgonMatches compares the existing hash with the newly hashed password to see if they match
func (p *Password) ArgonMatches(plaintextPassword string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plaintextPassword, p.Hash)
}

// BcryptMatches compares the existing hash with the newly hashed password to see if they match
func (p *Password) BcryptMatches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.Hash), []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
