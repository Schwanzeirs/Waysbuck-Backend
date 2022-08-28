package routes

import (
	"bewaysbuck/handlers"
	"bewaysbuck/pkg/middleware"
	"bewaysbuck/pkg/mysql"
	"bewaysbuck/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profiles", middleware.Auth(h.FindProfile)).Methods("GET")
	r.HandleFunc("/profile/{id}", middleware.Auth(h.GetProfile)).Methods("GET")
	r.HandleFunc("/profile", middleware.Auth(middleware.UploadFile(h.CreateProfile))).Methods("POST")
	r.HandleFunc("/profile/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProfile))).Methods("PATCH")
	r.HandleFunc("/profile/{id}", middleware.Auth(h.DeleteProfile)).Methods("DELETE")
}
