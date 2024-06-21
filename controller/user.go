package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michee/authentificationApi/model"
	"github.com/michee/authentificationApi/provider"
	"github.com/michee/authentificationApi/utils"
	"gorm.io/gorm"
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

	user.Token = token
    // Mettre à jour le modèle User dans la base de données avec le nouveau token
    model.DB.Save(&user)

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


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUpdate := &model.User{}
	utils.ParseBody(r, userUpdate)

	userId := chi.URLParam(r, "userId")

	// Vérifier le token avant de continuer
	if !provider.VerificationToken(r, userId) {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
	}

	var userDetail model.User
	err := model.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("user_id = ?", userId).First(&userDetail).Error; err != nil {
					return err
			}

			// Mettre à jour les champs modifiés
			if userUpdate.Name != "" {
					userDetail.Name = userUpdate.Name
			}
			if userUpdate.UserName != "" {
					userDetail.UserName = userUpdate.UserName
			}
			if userUpdate.Email != "" {
					userDetail.Email = userUpdate.Email
			}
			if userUpdate.Password != "" {
					hashedPassword, err := utils.HashPassword(userUpdate.Password)
					if err != nil {
							return err
					}
					userDetail.Password = hashedPassword
			}

			return tx.Save(&userDetail).Error
	})

	if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
					http.Error(w, "User not found", http.StatusNotFound)
			} else {
					http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
			}
			return
	}

	res, err := json.Marshal(userDetail)
	if err != nil {
			http.Error(w, "Failed to serialize user data", http.StatusInternalServerError)
			return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}





func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	if !provider.VerificationToken(r, userId) {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	u := model.DeleteUser(userId)
	res, _ := json.Marshal(u)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



func LogoutUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	u, _ := model.GetUserById(userId)

	if u == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
	}

	if !provider.VerificationToken(r, userId) {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// Révoquer le token en effaçant celui enregistré dans la base de données
	u.Logout()

	utils.RespondWithJSON(w, http.StatusOK, "Logout successful", nil)
}
