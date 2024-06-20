package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michee/authentificationApi/model"
	"github.com/michee/authentificationApi/provider"
	"github.com/michee/authentificationApi/utils"
)


var NewUser model.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	utils.ParseBody(r, user)
	u := user.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := model.GetUserByEmail(loginReq.Email)
	if err != nil || !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := provider.GenerateJWT(user.UserId)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Login successful", map[string]string{"token": token})
}


func GetUser(w http.ResponseWriter, r *http.Request){
	u := model.GetAllUser()
	res, _ := json.Marshal(u)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	u, _ := model.GetUserById(userId)
	
	res, _ := json.Marshal(u)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateUser(w http.ResponseWriter, r *http.Request){
	userUpdate := &model.User{}
	utils.ParseBody(r, &userUpdate)
	userId := chi.URLParam(r, "userId")

	userDetail, db := model.GetUserById(userId)

	if userDetail == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if userUpdate.Name != "" {
		userDetail.Name = userUpdate.Name
	}

	if userUpdate.UserName != "" {
		userDetail.UserName = userUpdate.UserName
	}

	if userUpdate.Email != "" {
		userDetail.Email = userUpdate.Email
	}

	if userDetail.Password != ""{
		hashedPassword, _ := utils.HashPassword(userUpdate.Password)
		userDetail.Password = hashedPassword
	}

	db.Save(&userDetail)

	res, _ := json.Marshal(userDetail)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	userId := chi.URLParam(r, "userId")
	u := model.DeleteUser(userId)

	res, _ := json.Marshal(u)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}