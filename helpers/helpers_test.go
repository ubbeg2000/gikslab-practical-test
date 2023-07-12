package helpers

import (
	"fmt"
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	p := "password"

	h, err := HashPassword(p)
	if err != nil {
		t.Errorf("Error on hashing password : %s\n", err.Error())
	}

	if err := ComparePasswordAndHash(p, h); err != nil {
		t.Errorf("Error on comparing hash : %s\n", err.Error())
	}
}

func TestCreateToken(t *testing.T) {
	os.Setenv("APP_SECRET", "LOREM")
	token, err := CreateToken(1, "board")
	if err != nil {
		t.Errorf("Error on creating token : %s\n", err.Error())
	}

	fmt.Println(token)

	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("Error on parsing token : %s\n", err.Error())
	}

	fmt.Printf("%+v\n", claims)
}
