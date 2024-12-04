package token
import (
    "fmt"
    "time"
    "errors"
    "github.com/golang-jwt/jwt/v4"
)
var secret_key = []byte("kavi@123")
type token_data struct{
    Username string `json:"username"`
    Role string `json:"role"`
    DB string   `json:"db"`
    jwt.RegisteredClaims
}
func Generate_Token(user string ,  role string, db string)(string, error){
    token_claims := token_data{
        Username: user,
        Role: role,
        DB: db,
        RegisteredClaims:jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }
    // creating token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, token_claims)
    // sign with secret key
    output , err := token.SignedString([]byte(secret_key))
    if err != nil {
        return "", err
    }else{
        return output, nil
        }
}
func Validate_Token(data string)(*token_data, error ){
    token, err := jwt.ParseWithClaims(data, &token_data{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secret_key, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*token_data); ok && token.Valid {
        // fmt.Println("Username:", claims.Username)
        // fmt.Println("Role:", claims.Role)
        // fmt.Println("Database:", claims.DB)
        return claims, nil
    } else {
        err := errors.New("invalid Token")
        return nil, err
    }
}




