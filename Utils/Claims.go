package Utils

import (
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"wan-api-verify-user/DTO"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/youmark/pkcs8"
)

func GenerateJWT(claims DTO.DicodeClaims, tokenType string) (string, error) {
	// Step 1: Defined expired time for the token based on the token type (1 day for access token, 7 days for refresh token)
	var expiredTime time.Time
	if tokenType == "ACCESS" {
		expiredTime = time.Now().Add(time.Hour * 24)
	} else if tokenType == "REFRESH" {
		expiredTime = time.Now().Add(time.Hour * 24 * 7)
	}

	// Step 2: Generate the token using the claims
	claims.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "Dicode_Authen_Service",
		Audience:  []string{"Dicode_User"},
		ExpiresAt: jwt.NewNumericDate(expiredTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(),
	}

	if tokenType == "ACCESS" {
		claims.Subject = "Access_Token"
	} else if tokenType == "REFRESH" {
		claims.Subject = "Refresh_Token"
	} else {
		return "", fmt.Errorf("invalid token type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	secretKey, err := ParsePrivateKey([]byte("decodebe"))
	if err != nil {
		return "", err
	}

	// Step 3: Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParsePrivateKey(passphrase []byte) (*rsa.PrivateKey, error) {
	// Read the encrypted private key from the PEM file
	privateKeyBytes, err := ioutil.ReadFile("./Utils/private_key.pem")
	if err != nil {
		return nil, fmt.Errorf("error reading the private key: %v", err)
	}

	// Decode the PEM block
	pemBlock, _ := pem.Decode(privateKeyBytes)
	if pemBlock == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}

	// Decrypt and parse the private key using pkcs8 library
	privateKey, err := pkcs8.ParsePKCS8PrivateKeyRSA(pemBlock.Bytes, passphrase)
	if err != nil {
		return nil, fmt.Errorf("error parsing the private key: %v", err)
	}

	return privateKey, nil
}

func VerifyJWT(tokenString string) (*DTO.DicodeClaims, error) {
	publicKey, err := os.ReadFile("public_key.pem")
	if err != nil {
		log.Fatalf("Error reading the public key: %v", err)
		return nil, err
	}
	publicKeyParsed, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &DTO.DicodeClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKeyParsed, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*DTO.DicodeClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}