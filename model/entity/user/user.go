package user

import (
	"encoding/hex"
	"errors"

	// valid "github.com/asaskevich/govalidator"
	"github.com/russross/meddler"
	"golang.org/x/crypto/bcrypt"

	"github.com/hachibeeDI/tiny-akasha/model/account/github"
	"github.com/hachibeeDI/tiny-akasha/model/entity"
)

type User struct {
	Id           int    `meddler:"id,pk" json:"id" valid:"required"`
	Name         string `meddler:"name" json:"name" valid:"required"`
	Password     string `meddler:"password" json:"password" valid:"ascii"`
	ImageUrl     string `meddler:"image_url" json:"image_url" valid:"url"`
	Email        string `meddler:"email" json:"email" valid:"required,email"`
	Introduction string `meddler:"introduction" json:"introduction" valid:"required"`
	// TODO: テーブル分けるほどのものじゃない？
	GithubId            int    `meddler:"github_id" json:"github_id"`
	GithubUrl           string `meddler:"github_url" json:"github_url"`
	GithubAccessToken   string `meddler:"github_access_token" json:"github_access_token"`
	FaceBookId          int    `meddler:"facebook_id" json:"facebook_id"`
	FaceBookUrl         string `meddler:"facebook_url" json:"facebook_url"`
	FaceBookAccessToken string `meddler:"facebook_access_token" json:"facebook_access_token"`
	TwitterId           int    `meddler:"twitter_id" json:"twitter_id"`
	TwitterUrl          string `meddler:"twitter_url" json:"twitter_url"`
	TwitterAccessToken  string `meddler:"twitter_access_token" json:"twitter_access_token"`
}

func DisposeTable(db entity.DB) {
	if _, err := db.Exec("drop table if exists user"); err != nil {
		panic(err)
	}
}

func CreateTableIfNotExists(db entity.DB) {
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS
			user(
				id integer AUTO_INCREMENT primary key
				, name varchar(40)
				, password varchar(40)
				, image_url varchar(255)
				, email varchar(255)
				, introduction MEDIUMTEXT
				, github_id int UNIQUE
				, github_url varchar(255)
				, github_access_token varchar(255)
				, facebook_id int UNIQUE
				, facebook_url varchar(255)
				, facebook_access_token varchar(255)
				, twitter_id int UNIQUE
				, twitter_url varchar(255)
				, twitter_access_token varchar(255)
			)CHARSET=utf8;`); err != nil {
		panic(err)
	}
}

func Init(Id int, Name, Password, ImageUrl, Introduction string) *User {
	return &User{Id: Id, Name: Name, Password: Password, ImageUrl: ImageUrl, Introduction: Introduction}
}

func InitByGithubAccount(guser github.UserAccount, accessToken string) *User {
	return &User{
		Name:              guser.Name,
		Password:          "",
		ImageUrl:          guser.AvatarUrl,
		Introduction:      guser.Bio,
		GithubId:          guser.Id,
		GithubUrl:         guser.Url,
		GithubAccessToken: accessToken,
	}
}

func (u *User) Insert(db entity.DB) error {
	// result, err := valid.ValidateStruct(u)
	// if err != nil {
	// 	return err
	// }
	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = hex.EncodeToString(cryptedPass)
	return meddler.Insert(db, "user", u)
}

func Update(db entity.DB, id int, givenPassword, username, imageUrl, introduction string) error {
	user := FindById(db, id)
	if user == nil {
		return errors.New("the user does not exists that is accord with id.")
	}
	pass, err := hex.DecodeString(user.Password)
	if err != nil {
		return err
	}
	if compared := bcrypt.CompareHashAndPassword(pass, []byte(givenPassword)); compared != nil {
		return compared
	}

	user.Name = username
	user.ImageUrl = imageUrl
	user.Introduction = introduction
	return meddler.Update(db, "user", user)
}

//
// func Delete(db entity.DB, id int) error {
// 	result, err := db.Exec("DELETE FROM user WHERE id = ?", id)
// 	if err != nil {
// 		return err
// 	}
// 	af, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if af != 1 {
// 		return errors.New("failed to delete")
// 	}
// 	return nil
// }
//
// func DeleteByQuestionID(db entity.DB, questionID int) error {
// 	result, err := db.Exec("DELETE FROM user WHERE question_id = ?", questionID)
// 	if err != nil {
// 		return err
// 	}
// 	af, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if af != 1 {
// 		return errors.New("failed to delete")
// 	}
// 	return nil
// }
//
// func SelectAll(db entity.DB) []*User {
// 	var users []*User
// 	err := meddler.QueryAll(db, &users, "SELECT * FROM user")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return users
// }
//
func FindById(db entity.DB, id int) *User {
	user := new(User)
	err := meddler.QueryRow(db, user, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return user
}

func FindByGithubId(db entity.DB, githubId int) (*User, error) {
	user := new(User)
	err := meddler.QueryRow(db, user, "SELECT * FROM user WHERE github_id = ?", githubId)
	return user, err
}

// func SelectByQuestionId(db entity.DB, question_id int) []*User {
// 	var users []*User
// 	err := meddler.QueryAll(db, &users, "SELECT * FROM user WHERE question_id = ?", question_id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if users == nil {
// 		return []*User{}
// 	}
// 	return users
// }
