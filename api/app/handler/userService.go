package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BomaxChen/restfulGorilla/api/app/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	fmt.Println(user)
	respondJSON(w, http.StatusOK, user)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	idInt, _ := strconv.Atoi(id)
	if err := db.First(&user, model.User{ID: idInt}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
