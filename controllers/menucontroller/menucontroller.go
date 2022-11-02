package menucontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Furqon-M/gopesan/helper"
	"github.com/Furqon-M/gopesan/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var menus []models.Menu

	if err := models.DB.Find(&menus).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, menus)
}

func Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var menu models.Menu
	if err := models.DB.First(&menu, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Menu Tidak Ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
		}
	}
	ResponseJson(w, http.StatusOK, menu)
}

func Create(w http.ResponseWriter, r *http.Request) {

	var menu models.Menu

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&menu); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&menu).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, menu)

}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var menu models.Menu

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&menu); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&menu).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak Dapat Mengupdate Menu")
		return
	}
	menu.Id = id

	ResponseJson(w, http.StatusOK, menu)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var menu models.Menu
	if models.DB.Delete(&menu, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak Dapat Menghapus Menu")
		return
	}
	response := map[string]string{"message": "Menu Berhasil Dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
