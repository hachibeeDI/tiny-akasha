package jwt

import (
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/hachibeeDI/tiny-akasha/model/entity/user"
)

var (
	defaultPrivKey, _ = ioutil.ReadFile("./key_for_token")
	defaultPubKey, _  = ioutil.ReadFile("./key_for_token.pub")
)

type UserToken struct {
	JWT   string `json:"jwt"`
	Error string `json:"error"`
}

func InitFromUser(u *user.User) *UserToken {
	jwt := jwt.New(jwt.SigningMethodRS512)
	jwt.Claims["iss"] = "tiny-akasha"
	jwt.Claims["user_id"] = u.Id
	jwt.Claims["image_url"] = u.ImageUrl
	jwt.Claims["name"] = u.Name
	// expire := time.Hour * 800
	// token.Claims["exp"] = time.Now().Add(expire).Unix()
	token, err := jwt.SignedString(defaultPrivKey)
	if err != nil {
		log.Print(err)
		return &UserToken{Error: err.Error()}
	}
	return &UserToken{JWT: token}
}
