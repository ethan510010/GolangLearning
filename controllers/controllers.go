package controllers

import (
	"GolangAPIPractice/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	usersInfo := models.ListAllUsers()
	resultData := make(map[string][]models.User)
	resultData["data"] = usersInfo
	err := json.NewEncoder(w).Encode(resultData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	validId, _ := strconv.Atoi(id)
	userDetail := models.GetUserProfile(validId)
	err := json.NewEncoder(w).Encode(userDetail)
	if err != nil {
		fmt.Println(err)
		return
	}
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

	avatarURL := ""
	if handler.Filename != "" {
		avatarURL = fmt.Sprintf("http://localhost:3000/static/%s", handler.Filename)
	}
	validAge, err := strconv.Atoi(userAge)
	u := &models.User{
		Name:       userName,
		Age:        validAge,
		AvatarPath: avatarURL,
	}

	models.CreateUserProfile(u)

	w.Write([]byte("create user profile successfully"))
}
