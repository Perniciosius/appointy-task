package utils_test

import (
	utils "appointy-task/utils/hashing"
	"testing"
)

func TestHash(t *testing.T) {
	password := "abc123"
	hashedPassword := utils.HashPassword(password, nil)
	if !utils.CompareHashedPassword(hashedPassword, password) {
		t.Error("Failed")
	}
}
