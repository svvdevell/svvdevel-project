package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    password := "eleganceautoadminpass777!"
    
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
    
    fmt.Println("Замініть в auth.go:")
    fmt.Printf("ADMIN_PASSWORD_HASH = \"%s\"\n", string(hash))
}