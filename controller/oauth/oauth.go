package oauth

import (
	// "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	// "golang.org/x/net/context"
	"github.com/hachibeeDI/tiny-akasha/model/entity"
	// "github.com/zenazn/goji/web"
	"github.com/hachibeeDI/tiny-akasha/model/account/github"
	"github.com/hachibeeDI/tiny-akasha/model/entity/user"
)

const (
	clientId = "36df85e7be84b6f6055d"
)

var clientSecret = os.Getenv("GITHUB_OAUTH_SECRET")

func NewRequestForGithub(method, url, accessToken string, param url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(param.Encode()))
	req.Header.Add("Authorization", fmt.Sprintf("token %s", accessToken))
	return req, err
}

func getGithubAccessToken(clientId, clientSecret, code string) (*http.Response, error) {
	authVal := url.Values{}
	authVal.Add("client_id", clientId)
	authVal.Add("client_secret", clientSecret)
	authVal.Add("code", code)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(authVal.Encode()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	return client.Do(req)
}

func getGithubUserInfo(access_token string) (github.UserAccount, error) {
	var guser github.UserAccount
	fmt.Printf("access_token is %s \n", access_token)
	if access_token == "" {
		return guser, errors.New("access_token is empty")
	}
	client := &http.Client{}
	req, _ := NewRequestForGithub("GET", "https://api.github.com/user", access_token, url.Values{})
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return guser, err
	}
	err = json.Unmarshal(body, &guser)
	if err != nil {
		fmt.Println(err)
		return guser, err
	}
	return guser, nil
}

func GithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	fmt.Printf("github code = %s \n", code)
	resp, err := getGithubAccessToken(clientId, clientSecret, code)
	body, _ := ioutil.ReadAll(resp.Body)
	authed, err := simplejson.NewJson(body)
	if err != nil {
		fmt.Fprintf(w, "err %s \n", err)
		return
	}
	fmt.Printf("calb body = %s \n", string(body))
	fmt.Printf("authed = %s \n", authed)
	if err != nil {
		log.Fatal(err)
		fmt.Fprint(w, err)
		return
	}
	// access_token=e72e16c7e42f292c6912e7710c838347ae178b4a&scope=user%2Cgist&token_type=bearer
	accessToken := authed.Get("access_token").MustString()
	gobj, err := getGithubUserInfo(accessToken)
	if err != nil {
		log.Println("get Github user information is failed")
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}
	// fmt.Fprintf(w, "your token is %s \n", accessToken)
	// fmt.Fprintf(w, "your name is %s \n", gobj.Name)
	u, err := GithubSignUpOrSignIn(gobj, accessToken)
	if err != nil || u == nil {
		fmt.Fprintf(w, "failed to sign in / up on github account = %s \n", err)
		return
	}
	fmt.Fprintf(w, "hello !  %s \n", u.Name)
}

func GithubSignUpOrSignIn(guser github.UserAccount, accessToken string) (*user.User, error) {
	db := entity.Db
	// NOTE: err may sql.ErrNoRows
	authedUser, err := user.FindByGithubId(db, guser.Id)
	log.Printf("find by github = %s\n", err)
	if err == nil && authedUser != nil {
		log.Printf("authed user = %+v\n", authedUser)
		return authedUser, nil
	}
	u := user.InitByGithubAccount(guser, accessToken)
	log.Printf("not authed user so create =  %+v\n", u)
	err = u.Insert(db)
	if err != nil {
		log.Printf("insert failed = %s\n", err)
		return nil, err
	}
	log.Printf("insert success = %+v\n", u)
	return u, err
}
