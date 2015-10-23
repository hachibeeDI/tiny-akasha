package oauth

import (
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
	// "github.com/zenazn/goji/web"
	// "github.com/hachibeeDI/tiny-akasha/model/entity/user"
)

const (
	clientId = "36df85e7be84b6f6055d"
)

// via: https://developer.github.com/v3/users/#get-a-single-user
type GithubUserObj struct {
	Login              string `json:"login"`
	Id                 int    `json:"id"`
	AvatarUrl          string `json:"avatar_url"`
	GravatarId         string `json:"gravatar_id"`
	Url                string `json:"url"`
	HtmlUrl            string `json:"html_url"`
	FollowersUrl       string `json:"followers_url"`
	FollowingUrl       string `json:"following_url"`
	GistsUrl           string `json:"gists_url"`
	StarredUrl         string `json:"starred_url"`
	SubscriptionsUrl   string `json:"subscriptions_url"`
	OrganizationsUrl   string `json:"organizations_url"`
	ReposUrl           string `json:"repos_url"`
	EventsUrl          string `json:"events_url"`
	Received_eventsUrl string `json:"received_events_url"`
	Type               string `json:"type"`
	SiteAdmin          bool   `json:"site_admin"`
	Name               string `json:"name"`
	Company            string `json:"company"`
	Blog               string `json:"blog"`
	Location           string `json:"location"`
	Email              string `json:"email"`
	Hireable           bool   `json:"hireable"`
	Bio                string `json:"bio"`
	PublicRepos        int    `json:"public_repos"`
	PublicGists        int    `json:"public_gists"`
	Followers          int    `json:"followers"`
	Following          int    `json:"following"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

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

func getGithubUserInfo(access_token string) (GithubUserObj, error) {
	var guser GithubUserObj
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
	fmt.Printf("g user obj  = %s \n", string(body))
	fmt.Printf("g user obj  = %s \n", guser)
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
	gobj, err := getGithubUserInfo(authed.Get("access_token").MustString())
	if err != nil {
		log.Println("get Github user information is failed")
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "your token is %s \n", authed.Get("access_token").MustString())
	fmt.Fprintf(w, "your name is %s \n", gobj.Name)
}
