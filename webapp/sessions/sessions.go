package sessions

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
)

type Session struct{
	Is_Authenticated bool
}




const tokenLength = 10
var sessionStorage map[string]*Session



func init() {
    sessionStorage = make(map[string]*Session)
}

func GetSession(token string) (session *Session, err error) {
	if session, exists := sessionStorage[token]; !exists {
            return session, errors.New("token does not exist")
	    } 
    return sessionStorage[token], nil
}

func SetSession() string {
	var token string
	for sessionIsUnique := false; sessionIsUnique != true; {
	    token = generateToken(tokenLength)
	    if _, exists := sessionStorage[token]; !exists {
            sessionIsUnique = true
	    }
    }    
    sessionStorage[token] = new(Session)
    sessionStorage[token].Is_Authenticated = false
    return token
}



func generateToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)

	if err != nil {
		panic(err)
	}
    return base32.StdEncoding.EncodeToString(randomBytes)[:length]

}