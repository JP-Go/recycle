package hasher

import "golang.org/x/crypto/bcrypt"

type Hash interface {
	String() string
}

/*
A Hasher is something that can compute hashes based on a string and
verify if some hash can be computed from some given data
*/
type Hasher interface {
	// Compute receives the string representation of the data to be hashed and
	// returns
	Compute(data string) (string, error)
	Verify(data string, hash string) (bool, error)
}

const defaultBcryptCost = 12

type BcryptHasher struct {
	cost int
}

type BcryptHasherConfig func(hasher *BcryptHasher)

func WithCost(cost ...int) BcryptHasherConfig {
	return func(hasher *BcryptHasher) {
		if len(cost) == 0 {
			return
		}
		hasher.cost = cost[0]
	}
}

func NewBcryptHasher(configs ...BcryptHasherConfig) *BcryptHasher {
	baseHasher := &BcryptHasher{cost: defaultBcryptCost}
	for _, cfg := range configs {
		cfg(baseHasher)
	}
	return baseHasher
}

func (bh *BcryptHasher) Compute(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bh.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (bh *BcryptHasher) Verify(data string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil, err
}
