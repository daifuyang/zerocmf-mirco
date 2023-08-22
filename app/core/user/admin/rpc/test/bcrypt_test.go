package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("zerocmf123456"), bcrypt.DefaultCost)
	fmt.Println("hashedPassword", string(hashedPassword))
}
