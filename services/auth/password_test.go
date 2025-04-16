package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hashedPassword, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hashedPassword == "" {
		t.Errorf("expected hashed password to not be empty")
	}

	if hashedPassword == "password" {
		t.Errorf("expected hashed password to not be equal to the original password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Errorf("expected passwords to match hash")
	}

	if ComparePasswords(hash, []byte("wrong-password")) {
		t.Errorf("expected passwords to not match hash")
	}
}
