package controllers

import (
	h "dragon-fruit-expert-system/helpers"
	"dragon-fruit-expert-system/models/db_models"
	"dragon-fruit-expert-system/models/requests"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"net/http"
)

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Conn: db,
	}
}

type AuthController struct {
	Conn *gorm.DB
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	var param requests.Login
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		h.ResError(w, r,400, "Kolom belum terisi, lengkapi terlebih dahulu")
		return
	}

	_, errors := govalidator.ValidateStruct(param)
	if errors != nil {
		h.ResError(w, r,400, "Kolom belum terisi, lengkapi terlebih dahulu")
		return
	}

	var checkUser models.User
	err := a.Conn.Where("email = ?", param.Email).First(&checkUser)
	if err.RecordNotFound() {
		h.ResError(w, r, 400, "Email atau password salah")
		return
	}

	result, errB := h.GenerateUserToken(checkUser.ID)
	if errB != nil {
		h.ResError(w, r, 400, err)
		return
	}

	h.ResWithoutMeta(w, r, 201, result)
}

func (a *AuthController) Register(w http.ResponseWriter, r *http.Request)  {
	var param requests.Register
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		h.ResError(w, r,400, "Kolom belum terisi, lengkapi terlebih dahulu")
		return
	}

	_, errors := govalidator.ValidateStruct(param)
	if errors != nil {
		h.ResError(w, r,400, "Kolom belum terisi, lengkapi terlebih dahulu")
		return
	}

	var checkUser models.User
	err := a.Conn.Where("email = ?", param.Email).First(&checkUser)
	if !err.RecordNotFound() {
		h.ResError(w, r, 400, "Email sudah terdaftar")
		return
	}

	user := models.User{
		Email:    param.Email,
		Name:     param.Name,
		Password: param.Password,
	}
	res := a.Conn.Model(&models.User{}).Create(&user)
	if res.Error != nil {
		h.ResError(w, r, 400, "Cannot create !")
		return
	}

	result, errB := h.GenerateUserToken(user.ID)
	if errB != nil {
		h.ResError(w, r, 400, err)
		return
	}

	h.ResWithoutMeta(w, r, 201, result)
}