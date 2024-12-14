package Utils

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

/*
	* GenerateSalt generates a random salt of the specified length.
	* It returns the salt as a hexadecimal string and an error if any occurs during the random byte generation.
	@param length: The length of the salt to be generated.
	@return string: The generated salt in hexadecimal format.
		- string: The generated salt in hexadecimal format.
		- error: An error if there is an issue generating the random bytes.
*/
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

/*
	* Generated hashed password using bcrypt and the salt provided.
	? Salt will also be stored inside the database for future reference.
*/
func HashPassword(password string, salt string) (string, error) {
	combinedPass := password + salt
	
	// Hash the password using bcrypt, use the default cost of 10, if the system is larger, increase the cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(combinedPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}