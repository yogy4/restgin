package atur

import (
	"fmt"
	"net/http"
	"restgin/basecon"
	"restgin/entitas"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

func CreateUser(c *gin.Context) {
	var err error
	dec := entitas.User{Password: c.PostForm("password")}
	hash, _ := HashPass(dec)
	user := entitas.User{Name: c.PostForm("name"), Email: c.PostForm("email"), Password: hash}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
	basecon.Db.Save(&user)
	var us entitas.User
	tk := &Token{UserId: us.ID}
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
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

func CheckPass(ps, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ps), []byte(hash))
	return err == nil
}

func Login(c *gin.Context) {
	var us entitas.User

	email := c.PostForm("email")
	pass := c.PostForm("password")

	cmail := basecon.Db.Where("email = ?", email).First(&us).Error
	if cmail != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "email not found"})
		return
	}

	cpass := CheckPass(us.Password, pass)
	str := strconv.FormatBool(cpass)
	fmt.Printf("Match : %s \n", str)

	if str != "true" {

		c.JSON(http.StatusNotFound, gin.H{"message": "wrong pass or username"})
		return

	}

	tk := &Token{UserId: us.ID}
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "token": token})

}
