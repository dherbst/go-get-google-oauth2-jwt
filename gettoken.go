package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func main() {

	conf := &jwt.Config{
		Email:      "email",
		PrivateKey: []byte("privatekey"),
		Scopes: []string{
			"https://www.googleapis.com/auth/projecthosting",
			"https://www.googleapis.com/auth/appengine.admin",
		},
		TokenURL: google.JWTTokenURL,
		// If you would like to impersonate a user, you can
		// create a transport with a subject. The following GET
		// request will be made on the behalf of user@example.com.
		// Optional.
		//Subject: "email@example.com",
	}

	tokensource := conf.TokenSource(oauth2.NoContext)
	fmt.Printf("tokenSrc=%v\n", tokensource)
	token, err := tokensource.Token()
	if err != nil {
		fmt.Printf("error getting token %v\n", err)
		return
	}
	fmt.Printf("access_token=%v\nexpiry=%v\ntoken_type=%v\n", token.AccessToken, token.Expiry,
		token.TokenType,
	)
}
