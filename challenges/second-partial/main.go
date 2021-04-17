package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var info = gin.H{
	"username": gin.H{"email": "username@gmail.com", "token": ""},
	"mariana":  gin.H{"email": "mariana@gmail.com", "token": ""},
}

var tokens = make(map[string]string)

func main() {

	print("inicia\n")

	r := gin.Default()
	r.Use()

	auth := r.Group("/", gin.BasicAuth(gin.Accounts{"username": "password", "mariana": "gomez"}))

	print("crea grupos\n")

	auth.GET("/login", login)
	auth.GET("/logout", logout)
	auth.GET("/status", status)
	auth.GET("/upload", upload)
	r.Run(":8080")

	print("termina\n")
}

func login(c *gin.Context) {

	print("inicia login\n")

	user := c.MustGet(gin.AuthUserKey).(string)
	token := GenerateSecureToken(11)

	tokens[user] = token

	if _, usok := info[user]; usok {
		print("el token es: ", token, " el usuario es: ", user, "\n ")
		c.JSON(http.StatusOK, gin.H{"message": "Hi username welcome to the DPIP System", "token": tokens[user]})
	} else {
		c.AbortWithStatus(401)
	}
	print("termina login\n")
}

func logout(c *gin.Context) {

	print("inicio de logoff\n")
	user := c.MustGet(gin.AuthUserKey).(string)

	if _, usok := tokens[user]; usok {
		print("se pudo cerrar session\n")

		delete(tokens, user)
		c.AbortWithStatus(401)
		c.JSON(http.StatusOK, gin.H{"message": "Bye username, your token has been revoked"})
		return
	} else {
		print("no se pudo\n")
		c.AbortWithStatus(401)

	}

	print("final de logoff \n")
}

func status(c *gin.Context) {
	print("inicio de status\n")
	user := c.MustGet(gin.AuthUserKey).(string)
	dt := time.Now()
	fmt.Println(tokens)
	if _, usok := tokens[user]; usok {
		print("el tiempo es \n")
		c.JSON(http.StatusOK, gin.H{"message": "Hi username, the DPIP System is Up and Running", "time": dt})
	} else {
		c.AbortWithStatus(401)
	}
	print("fin de status\n")
}

func upload(c *gin.Context) {
	print("inicio de upload\n")
	user := c.MustGet(gin.AuthUserKey).(string)
	if _, usok := tokens[user]; usok {
		_, header, err := c.Request.FormFile("image")
		if err != nil {
			return
		}
		size := strconv.Itoa(int(header.Size))
		c.JSON(http.StatusOK, gin.H{"status": "SUCCESS", "Filename": header.Filename, "filesize": size + " bytes"})
		print("nombre de archivo es ", header.Filename, " y su peso es de ", size, "\n")

	} else {
		c.AbortWithStatus(401)
	}
	print("fin de upload\n")
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	//fmt.Println(hex.EncodeToString(b))
	return hex.EncodeToString(b)
}
