package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/michee/authentificationApi/controller"
)


const port = ":8080"

var tokenAuth *jwtauth.JWTAuth


func main() {
	tokenAuth = jwtauth.New("HS256", []byte("ksQD5adHXZ-5SSJCupcHwBzDi6q5kfr5hdU7Eq5tMmo"), nil)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.RequestID)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", controller.CreateUser)
		r.Post("/login", controller.LoginHandler)

	})

	r.Route("/user", func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/", controller.GetUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Post("/", controller.LogoutUser)
			r.Get("/", controller.GetUserById)
			r.Patch("/", controller.UpdateUser)
			r.Delete("/", controller.DeleteUser)
		})

	})



	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	http.ListenAndServe(port, r)

}