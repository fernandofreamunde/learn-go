package main

// go get golang.org/x/crypto/bcrypt
import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	badPassword := "secreto"
	hash, _ := HashPassword(password) // ignore errors for simplicity

	fmt.Println("Password: ", password)
	fmt.Println("Hash:     ", hash)

	mathch := CheckPasswordHash(password, hash)
	fmt.Println("Match:    ", mathch)

	fmt.Println("\n----- now with bad password ----\n")

	mathch = CheckPasswordHash(badPassword, hash)
	fmt.Println("Match:    ", mathch)
}
