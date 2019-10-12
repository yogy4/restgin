package atur

import (
	"fmt"
	"net/http"
	"restgin/basecon"
	"restgin/entitas"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var err error
	dec := entitas.User{Password: c.PostForm("password")}
	hash, _ := HashPass(dec)
	user := entitas.User{Name: c.PostForm("name"), Email: c.PostForm("email"), Password: hash}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
	basecon.Db.Save(&user)
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Ok", "result": user, "token": token})

}

func HashPass(ps entitas.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(ps.Password), 14)
	return string(hash), err
}

func CheckPass(ps string, hash entitas.User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash.Password), []byte(ps))
	return err == nil
}

func Login(c *gin.Context) {
	var err error
	u := &entitas.User{}
	email := entitas.User{Email: c.PostForm("email")}
	pass := entitas.User{Password: c.PostForm("password")}

	basecon.Db.Where("email = ?", email).First(u)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "email not found"})
	}
	cp := CheckPass(u.Password, pass)
	fmt.Println("match", cp)
	u.Password = ""
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Ok", "token": token})

}
