package routes

import (
	"GST/billlingSystem/handler"
	"GST/billlingSystem/service"
	"GST/billlingSystem/utils"
	"context"
	"fmt"
	"log"
	"strings"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var skippedUrls = []string{"/login"}

func InitRouter(deps service.Dependencies) (router *mux.Router) {
	router = mux.NewRouter()
	router.Use(middleware)

	router.HandleFunc("/login", handler.Login(deps)).Methods(http.MethodPost)
	router.HandleFunc("/addTask", handler.AddTask(deps)).Methods(http.MethodPost)
	router.HandleFunc("/tasks", handler.GetAllTasks(deps)).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}", handler.GetTaskByID(deps)).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.UpdateTaskByID(deps)).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.DeleteTaskByID(deps)).Methods("DELETE")

	return
}

func skipUrls(url string) (toSkip bool) {

	for _, v := range skippedUrls {
		if url == v {
			toSkip = true
			return
		}
	}
	return false
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if skipUrls(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		if tokenString == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return utils.SecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			log.Printf("Token parse error: %v", err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userId := claims["user_id"].(string)
		log.Printf("Authenticated user: %s", userId)

		ctx := r.Context()
		ctx = context.WithValue(ctx, "UserID", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
