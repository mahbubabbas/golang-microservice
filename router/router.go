package router

import (
	"context"
	"golang-microsvc/controllers"
	"golang-microsvc/models"
	"golang-microsvc/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func ConfigRouter() *mux.Router {
	router := mux.NewRouter()

	//Router for students - these should be secured
	router.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/students", controllers.InsertStudentsMany).Methods("POST")
	router.HandleFunc("/students/{id}", controllers.UpdateStudents).Methods("PATCH")
	router.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	router.HandleFunc("/students/print", controllers.PrintStudent).Methods("POST")

	//Router for user - login etc
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	//router.Use(jwtTokenMiddleware)

	return router
}

var jwtTokenMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		noAuthUrls := []string{
			"/login",
		}

		reqPath := r.URL.Path
		for _, url := range noAuthUrls {
			if url == reqPath {
				next.ServeHTTP(rw, r)
				return
			}
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			rw.WriteHeader(http.StatusForbidden)
			utils.Respond(rw, utils.ErrorResponse("Missing token"))
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			rw.WriteHeader(http.StatusForbidden)
			utils.Respond(rw, utils.ErrorResponse("Malformed token"))
			return
		}

		tokenString := tokenParts[1]
		tk := models.MyToken{}

		token, err := jwt.ParseWithClaims(tokenString, &tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(models.PASS_KEY), nil
		})

		if err != nil {
			rw.WriteHeader(http.StatusForbidden)
			utils.Respond(rw, utils.ErrorResponse("Malformed token"))
			return
		}

		if !token.Valid {
			rw.WriteHeader(http.StatusForbidden)
			utils.Respond(rw, utils.ErrorResponse("Invalid token"))
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.User)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
