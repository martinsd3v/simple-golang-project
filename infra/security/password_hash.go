package security

import (
	"golang.org/x/crypto/bcrypt"
)

//IHash interface of hash
type IHash interface {
	Create(payload string) string
	Compare(hashed string, payload string) bool
}

//Hash struct for assign IHasg
type Hash struct{}

//Assign interface
var _ IHash = Hash{}

//Create hashed password from payload
func (hash Hash) Create(payload string) string {
	crypted, _ := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	return string(crypted)
}

//Compare hashed and playload
func (hash Hash) Compare(hashed string, payload string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(payload))
	return check == nil
}
