package handlers

import (
	profiledto "bewaysbuck/dto/profile"
	dto "bewaysbuck/dto/result"
	"bewaysbuck/models"
	"bewaysbuck/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

// var profileimage = "http://localhost:5000/uploads/"

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) FindProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	profiles, err := h.ProfileRepository.FindProfile(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	// for i, p := range profiles {
	// 	profiles[i].Image = profileimage + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: profiles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var profile models.Profile
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// profile.Image = profileimage + profile.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: profile}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	postcode, _ := strconv.Atoi(r.FormValue("postcode"))

	request := profiledto.ProfileRequest{
		Phone:    r.FormValue("phone"),
		Address:  r.FormValue("address"),
		Gender:   r.FormValue("gender"),
		Postcode: postcode,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Declare Context Background, Cloud Name, API Key, API Secret ...
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysbuck"})

	if err != nil {
		fmt.Println(err.Error())
	}

	profile := models.Profile{
		Image:    resp.SecureURL,
		Phone:    request.Phone,
		Gender:   request.Gender,
		Address:  request.Address,
		Postcode: request.Postcode,
		UserID:   userId,
	}

	profile, err = h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	profile, _ = h.ProfileRepository.GetProfile(profile.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: profile}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContext := r.Context().Value("dataFile")
	filename := dataContext.(string)

	request := profiledto.ProfileRequest{
		Phone:   r.FormValue("phone"),
		Address: r.FormValue("address"),
		Gender:  r.FormValue("gender"),
		UserID:  userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	profile, _ := h.ProfileRepository.GetProfile(id)

	profile.Phone = request.Phone
	profile.Gender = request.Gender
	profile.Address = request.Address

	if filename != "false" {
		profile.Image = filename
	}

	profile, err = h.ProfileRepository.UpdateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: profile}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteProfile, err := h.ProfileRepository.DeleteProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteProfile}
	json.NewEncoder(w).Encode(response)
}
