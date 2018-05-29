package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

// GetToken -
func GetToken(w http.ResponseWriter, r *http.Request) {
	// log.Printf("\nShops: %#v\n", shops)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	log.Printf("user.GetToken token:\n%#v", tokenString)
	// w.Write([]byte(tokenString))
	json.NewEncoder(w).Encode(tokenString)
}
