package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	avatarFile, handler, err := r.FormFile("avatar")
	userName := r.FormValue("userName")
	userAge := r.FormValue("userAge")

	fmt.Println(userName, userAge)
	// check image type
	ext := strings.ToLower(path.Ext(handler.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		fmt.Println("not supperted file image extension")
		return
	}

	// save image
	saveFile, err := os.OpenFile("./static/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(saveFile, avatarFile)

	defer avatarFile.Close()
	defer saveFile.Close()
	//
	w.Write([]byte("create user profile successfully"))
}