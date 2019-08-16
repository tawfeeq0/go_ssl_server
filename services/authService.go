package services


import (
	  "github.com/dgrijalva/jwt-go"
	  "time"
	  "fmt"
	  "encoding/base64"
  )
  
// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")
// User : struct
type User struct{
	BadgeNo string  `json:"BadgeNo"`
	CategoryNames []string `json:"CategoryNames"`
	jwt.StandardClaims
}

  
// Signin : method
func Signin(BadgeNo string) (string, error) {
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["BadgeNo"] = BadgeNo
	claims["CategoryNames"] = [ ... ]string{"69KV","DOA","NOA"}
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	pass,_ := base64.StdEncoding.DecodeString("U2NhZGE0UGk=")
	fmt.Println(string(pass))
	tokenString, err := token.SignedString(pass)
	fmt.Println(tokenString,err)
    return tokenString, err
}

// Verify : method
func Verify(reqToken string) bool{
	fmt.Println(reqToken)
	pass,_ := base64.StdEncoding.DecodeString("U2NhZGE0UGk=")
    token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
        return pass, nil
    })
    if err == nil && token.Valid {
        fmt.Println("valid token")
        return true
    } else {
        fmt.Println("invalid token")
        return false
    }
}

// GetUserInfo : method
func GetUserInfo(reqToken string) User{
	fmt.Println(reqToken)
	pass,_ := base64.StdEncoding.DecodeString("U2NhZGE0UGk=")
    token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
        return pass, nil
	})
	if err != nil {
		return User{}
	}

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("valid token")
		fmt.Println(claims["BadgeNo"])
		//return User{BadgeNo: claims["BadgeNo"], CategoryNames: claims["CategoryNames"]}
		return User{}

    } else {
        fmt.Println("invalid token")
        return User{}
    }
}