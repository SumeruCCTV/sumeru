// Package argon2id provides a convience wrapper around Go's golang.org/x/crypto/argon2
// implementation, making it simpler to securely hash and verify passwords
// using Argon2.
//
// It enforces use of the Argon2id algorithm variant and cryptographically-secure
// random salts.
package argon2id

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	StringVersion = strconv.Itoa(argon2.Version)

	// ErrInvalidHash in returned by ComparePasswordAndHash if the provided
	// hash isn't in the expected format.
	ErrInvalidHash = errors.New("argon2id: hash is not in the correct format")

	// ErrIncompatibleVariant is returned by ComparePasswordAndHash if the
	// provided hash was created using a unsupported variant of Argon2.
	// Currently only argon2id is supported by this package.
	ErrIncompatibleVariant = errors.New("argon2id: incompatible variant of argon2")

	// ErrIncompatibleVersion is returned by ComparePasswordAndHash if the
	// provided hash was created using a different version of Argon2.
	ErrIncompatibleVersion = errors.New("argon2id: incompatible version of argon2")
)

func CompareHashes(hash1, hash2 string) (bool, error) {
	key1, err := DecodeHash(hash1)
	if err != nil {
		return false, err
	}
	key2, err := DecodeHash(hash2)
	if err != nil {
		return false, err
	}
	key1Len := int32(len(key1))
	key2Len := int32(len(key2))

	if subtle.ConstantTimeEq(key1Len, key2Len) == 0 {
		return false, nil
	}
	if subtle.ConstantTimeCompare(key1, key2) == 1 {
		return true, nil
	}
	return false, nil
}

func DecodeHash(hash string) (key []byte, err error) {
	vars := strings.Split(hash, "$")
	if len(vars) != 6 {
		return nil, ErrInvalidHash
	}

	if vars[1] != "argon2id" {
		return nil, ErrIncompatibleVariant
	}

	var version int
	_, err = fmt.Sscanf(vars[2], "v=%d", &version)
	if err != nil {
		return nil, err
	}
	if version != argon2.Version {
		return nil, ErrIncompatibleVersion
	}

	key, err = base64.RawStdEncoding.Strict().DecodeString(vars[5])
	if err != nil {
		return nil, err
	}

	return key, nil
}

func ValidHash(hash string) bool {
	vars := strings.SplitN(hash, "$", 4)
	if len(vars) != 4 || vars[1] != "argon2id" {
		return false
	}
	if len(vars[2]) < 3 || vars[2][2:] != StringVersion {
		return false
	}
	return true
}
