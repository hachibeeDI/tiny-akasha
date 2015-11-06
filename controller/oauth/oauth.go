package oauth

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	// "golang.org/x/net/context"
	"github.com/hachibeeDI/tiny-akasha/model/entity"
	// "github.com/zenazn/goji/web"
	"github.com/hachibeeDI/tiny-akasha/model/account/github"
	"github.com/hachibeeDI/tiny-akasha/model/account/jwt"
	"github.com/hachibeeDI/tiny-akasha/model/entity/user"
)

func getGithubUserInfo(accessToken *oauth2.Token) (github.UserAccount, error) {
	var guser github.UserAccount
	fmt.Printf("access_token is %+v \n", accessToken)
	client := github.OAuthConf.Client(oauth2.NoContext, accessToken)
	resp, err := client.Get("https://api.github.com/user")
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
	accessToken, err := github.OAuthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	gobj, err := getGithubUserInfo(accessToken)
	if err != nil {
		log.Println("get Github user information is failed")
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}
	// fmt.Fprintf(w, "your token is %s \n", accessToken)
	// fmt.Fprintf(w, "your name is %s \n", gobj.Name)
	u, err := githubSignUpOrSignIn(gobj, accessToken)
	if err != nil || u == nil {
		fmt.Fprintf(w, "failed to sign in / up on github account = %s \n", err)
		return
	}
	token := jwt.InitFromUser(u)
	tmpl := template.Must(template.ParseFiles("template/oauth-callback.html"))
	tmpl.Execute(w, token)
}

func githubSignUpOrSignIn(guser github.UserAccount, accessToken *oauth2.Token) (*user.User, error) {
	db := entity.Db
	// NOTE: err may sql.ErrNoRows
	authedUser, err := user.FindByGithubId(db, guser.Id)
	log.Printf("find by github = %s\n", err)
	if err == nil && authedUser != nil {
		log.Printf("authed user = %+v\n", authedUser)
		return authedUser, nil
	}
	u := user.InitByGithubAccount(guser, accessToken.AccessToken)
	log.Printf("not authed user so create =  %+v\n", u)
	err = u.Insert(db)
	if err != nil {
		log.Printf("insert failed = %s\n", err)
		return nil, err
	}
	log.Printf("insert success = %+v\n", u)
	return u, err
}
