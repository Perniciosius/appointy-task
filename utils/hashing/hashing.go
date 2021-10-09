package hashing

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"log"
)

const saltSize = 16

func generateRandomSalt() []byte {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		log.Fatalln(err)
	}
	return salt
}

// Function to generate hashed password using SHA512 algorithm
func HashPassword(password string, salt []byte) string {
	if salt == nil {
		salt = generateRandomSalt()
	}
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	hashedPasswordBytes = append(salt, hashedPasswordBytes...)
	return base64.URLEncoding.EncodeToString(hashedPasswordBytes)
}

// Function to check password is correct
func CompareHashedPassword(hashedPassword, password string) bool {
	hashedPasswordBytes, err := base64.URLEncoding.DecodeString(hashedPassword)
	if err != nil {
		log.Fatalln(err)
	}

	passwordHash := HashPassword(password, hashedPasswordBytes[:saltSize])
	return hashedPassword == passwordHash
}
