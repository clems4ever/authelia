package authentication

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/authelia/authelia/internal/utils"
	"github.com/simia-tech/crypt"
)

// PasswordHash represents all characteristics of a password hash.
// Authelia only supports salted SHA512 or salted argon2id method, i.e., $6$ mode or $argon2id$ mode.
type PasswordHash struct {
	Algorithm   string
	Iterations  int
	Salt        string
	Key         string
	KeyLength   int
	Memory      int
	Parallelism int
}

// ParseHash extracts all characteristics of a hash given its string representation.
func ParseHash(hash string) (passwordHash *PasswordHash, err error) {
	parts := strings.Split(hash, "$")
	code, parameters, salt, key, err := crypt.DecodeSettings(hash)
	if err != nil {
		return nil, err
	}
	h := &PasswordHash{}

	h.Salt = salt
	h.Key = key
	if h.Key != parts[len(parts)-1] {
		return nil, fmt.Errorf("Cannot parse hash key is not the last parameter, the hash is probably malformed (%s).", hash)
	}
	if code == HashingAlgorithmSHA512 {
		h.Iterations = parameters.GetInt("rounds", HashingDefaultSHA512Iterations)
		h.Algorithm = HashingAlgorithmSHA512
		if h.Key == "" {
			return nil, fmt.Errorf("Cannot parse hash key contains no characters or the field length is invalid (%s)", hash)
		}
		if !utils.IsStringBase64Valid(h.Key) {
			return nil, fmt.Errorf("Cannot parse hash key contains invalid base64 characters.")
		}
		if !utils.IsStringBase64Valid(h.Salt) {
			return nil, fmt.Errorf("Cannot parse hash salt contains invalid base64 characters.")
		}
		if parameters["rounds"] != "" && parameters["rounds"] != strconv.Itoa(h.Iterations) {
			return nil, fmt.Errorf("Cannot parse hash sha512 rounds is not numeric (%s).", parameters["rounds"])
		}
	} else if code == HashingAlgorithmArgon2id {
		version := parameters.GetInt("v", 0)
		if version < 19 {
			if version == 0 {
				return nil, fmt.Errorf("Cannot parse hash argon2id version parameter not found (%s)", hash)
			}
			return nil, fmt.Errorf("Cannot parse hash argon2id versions less than v19 are not supported (hash is version %d).", version)
		} else if version > 19 {
			return nil, fmt.Errorf("Cannot parse hash argon2id versions greater than v19 are not supported (hash is version %d).", version)
		}
		h.Algorithm = HashingAlgorithmArgon2id
		h.Memory = parameters.GetInt("m", HashingDefaultArgon2idMemory)
		h.Iterations = parameters.GetInt("t", HashingDefaultArgon2idTime)
		h.Parallelism = parameters.GetInt("p", HashingDefaultArgon2idParallelism)
		h.KeyLength = parameters.GetInt("k", HashingDefaultArgon2idKeyLength)

		if h.Key == "" {
			return nil, fmt.Errorf("Cannot parse hash key contains no characters or the field length is invalid (%s)", hash)
		}
		if !utils.IsStringBase64Valid(h.Key) {
			return nil, fmt.Errorf("Cannot parse hash key contains invalid base64 characters.")
		}
		if !utils.IsStringBase64Valid(h.Salt) {
			return nil, fmt.Errorf("Cannot parse hash salt contains invalid base64 characters.")
		}
		decodedKey, _ := crypt.Base64Encoding.DecodeString(h.Key)
		if len(decodedKey) != h.KeyLength {
			return nil, fmt.Errorf("Cannot parse hash argon2id key length parameter (%d) does not match the actual key length (%d).", h.KeyLength, len(decodedKey))
		}
	} else {
		return nil, fmt.Errorf("Authelia only supports salted SHA512 hashing ($6$) and salted argon2id ($argon2id$), not $%s$", code)
	}
	return h, nil
}

// HashPassword generate a salt and hash the password with the salt and a constant
// number of rounds.
func HashPassword(password, salt, algorithm string, iterations, memory, parallelism, keyLength, saltLength int) (hash string, err error) {
	var settings string

	if algorithm != HashingAlgorithmArgon2id && algorithm != HashingAlgorithmSHA512 {
		return "", fmt.Errorf("Hashing algorithm input of '%s' is invalid, only values of %s and %s are supported.", algorithm, HashingAlgorithmArgon2id, HashingAlgorithmSHA512)
	}

	if salt == "" {
		if saltLength < 2 {
			return "", fmt.Errorf("Salt length input of %d is invalid, it must be 2 or higher.", saltLength)
		} else if saltLength > 16 {
			return "", fmt.Errorf("Salt length input of %d is invalid, it must be 16 or lower.", saltLength)
		}
	} else if len(salt) > 16 {
		return "", fmt.Errorf("Salt input of %s is invalid (%d characters), it must be 16 or fewer characters.", salt, len(salt))
	} else if len(salt) < 2 {
		return "", fmt.Errorf("Salt input of %s is invalid (%d characters), it must be 2 or more characters.", salt, len(salt))
	} else if !utils.IsStringBase64Valid(salt) {
		return "", fmt.Errorf("Salt input of %s is invalid, only characters [a-zA-Z0-9+/] are valid for input.", salt)
	}
	if algorithm == HashingAlgorithmArgon2id {
		if memory < 8 {
			return "", fmt.Errorf("Memory (argon2id) input of %d is invalid, it must be 8 or higher.", memory)
		}
		if parallelism < 1 {
			return "", fmt.Errorf("Parallelism (argon2id) input of %d is invalid, it must be 1 or higher.", parallelism)
		}
		if memory < parallelism*8 {
			return "", fmt.Errorf("Memory (argon2id) input of %d is invalid with a paraellelism input of %d, it must be %d (parallelism * 8) or higher.", memory, parallelism, parallelism*8)
		}
	}

	if salt == "" {
		salt = utils.RandomString(saltLength, HashingPossibleSaltCharacters)
	}
	if algorithm == HashingAlgorithmArgon2id {
		settings, _ = crypt.Argon2idSettings(memory, iterations, parallelism, keyLength, salt)
	} else if algorithm == HashingAlgorithmSHA512 {
		settings = fmt.Sprintf("$6$rounds=%d$%s", iterations, salt)
	}

	hash, err = crypt.Crypt(password, settings)
	if err != nil {
		log.Fatal(err)
	}
	return hash, nil
}

// CheckPassword check a password against a hash.
func CheckPassword(password, hash string) (ok bool, err error) {
	passwordHash, err := ParseHash(hash)
	if err != nil {
		return false, err
	}
	expectedHash, err := HashPassword(password, passwordHash.Salt, passwordHash.Algorithm, passwordHash.Iterations, passwordHash.Memory, passwordHash.Parallelism, passwordHash.KeyLength, len(passwordHash.Salt))
	if err != nil {
		return false, err
	}
	return hash == expectedHash, nil
}
