# go-get-google-oauth2-jwt
Example for getting the google oauth2/jwt for deploying appengine

If you create a service account for Google AppEngine and you want to let jenkins auto deploy your project, you will need an oauth2/jwt token to do so.

Download your service account json credentials.

    `go run gettoken.go`
