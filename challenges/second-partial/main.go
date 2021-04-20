package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var info = gin.H{
	"username": gin.H{"email": "username@gmail.com", "token": ""},
}

var tokens = make(map[string]string)

func main() {

	r := gin.Default()
	r.Use()

	auth := r.Group("/", gin.BasicAuth(gin.Accounts{"username": "password"}))

	auth.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/status", status)
	r.GET("/upload", upload)
	r.Run()

}

func login(c *gin.Context) {

	userToken := c.MustGet(gin.AuthUserKey).(string)

	print(userToken)

	user := c.MustGet(gin.AuthUserKey).(string)
	token := GenerateSecureToken(1)

	tokens[user] = token

	if _, userOk := info[user]; userOk {
		c.JSON(http.StatusOK, gin.H{"message": "Hi " + user + " welcome to the DPIP System", "token": tokens[user]})
	} else {
		c.AbortWithStatus(401)
	}

}

func logout(c *gin.Context) {

	exist, user, _ := auth(c)

	if exist == true {
		delete(tokens, user)
		c.AbortWithStatus(401)
		c.JSON(http.StatusOK, gin.H{"message": "Bye " + user + ", your token has been revoked"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Invalid Token"})
		c.AbortWithStatus(401)
	}

}

func status(c *gin.Context) {

	exist, user, _ := auth(c)

	if exist == true {
		current := time.Now()
		c.JSON(http.StatusOK, gin.H{"message": "Hi " + user + ", the DPIP System is Up and Running", "time": current.Format("2006-01-02 15:04:05")})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Invalid Token"})
		c.AbortWithStatus(401)
	}

}

func upload(c *gin.Context) {
	exist, _, _ := auth(c)

	if exist == true {
	_,header,err :=c.Request.FormFile("data")
    if err!=nil{
		c.AbortWithStatus(401)

        return
    }
    size:= strconv.Itoa(int(header.Size))
    c.JSON(http.StatusOK,gin.H{"status":"SUCCESS","Filename":header.Filename,"filesize":size+" bytes"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Invalid Token"})
		c.AbortWithStatus(401)
	}
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	//fmt.Println(hex.EncodeToString(b))
	return hex.EncodeToString(b)
}

func auth(c *gin.Context) (bool, string, string) {

	exist := false

	bearer := c.Request.Header["Authorization"]
	bearerToken := bearer[0]
	splitedToken := strings.Split(bearerToken, " ")
	token := string(splitedToken[1])

	userName := ""
	userToken := ""

	for user, tokenList := range tokens {

		if token == tokenList {
			exist = true
			userToken = tokenList
			userName = user
		}

	}

	return exist, userName, userToken
}
