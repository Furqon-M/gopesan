package pesancontroller

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
	var pesandbs []models.Pesandb

	if err := models.DB.Find(&pesandbs).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, pesandbs)
}

func Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var pesandb models.Pesandb
	if err := models.DB.First(&pesandb, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Pesanan Tidak Ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
		}
	}
	ResponseJson(w, http.StatusOK, pesandb)
}

func Create(w http.ResponseWriter, r *http.Request) {

	var pesandb models.Pesandb

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pesandb); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&pesandb).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, pesandb)

}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var pesandb models.Pesandb

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pesandb); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&pesandb).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak Dapat Mengupdate Pesanan")
		return
	}
	pesandb.Id = id

	ResponseJson(w, http.StatusOK, pesandb)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var pesandb models.Pesandb
	if models.DB.Delete(&pesandb, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak Dapat Menghapus Pesanan")
		return
	}
	response := map[string]string{"message": "Pesanan Berhasil Dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
