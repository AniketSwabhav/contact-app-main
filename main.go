package main

import (
	contactController "contact-app-main/components/contact/controller"
	contactDetailController "contact-app-main/components/contactDetail/controller"
	"contact-app-main/components/security"
	userController "contact-app-main/components/user/controller"
	"contact-app-main/models/credential"
	"contact-app-main/models/user"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Create a hardcoded admin user
	defaultEmail := "god@father.com"
	defaultPassword := "admin1234"

	// Create credential
	cred, err := credential.CreateCredential(defaultEmail, defaultPassword)
	if err != nil {
		log.Fatalf("Failed to create default admin credential: %v", err)
	}

	// Create admin user
	user.CreateAdmin("Super", "Admin", defaultEmail, cred.CredentialID)

	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter().StrictSlash(true)

	// ----------------------------------------------------------------------------------------------------------------------------------

	adminRouter := router.PathPrefix("/").Subrouter()
	// Middleware for security
	adminRouter.Use(security.MiddlewareAdmin)

	// login
	router.HandleFunc("/login", userController.Login).Methods("POST")

	// Admin creation
	adminRouter.HandleFunc("/admins", userController.RegisterAdmin).Methods("POST")

	// User management
	adminRouter.HandleFunc("/users", userController.RegisterUser).Methods("POST")
	adminRouter.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	adminRouter.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	adminRouter.HandleFunc("/users/{id}", userController.UpdateUserByID).Methods("PUT")
	adminRouter.HandleFunc("/users/{id}", userController.DeleteUserByID).Methods("DELETE")

	// ----------------------------------------------------------------------------------------------------------------------------------

	userRouter := router.PathPrefix("/users/{userId}/contact").Subrouter()
	userRouter.Use(security.MiddlewareContact)

	userRouter.HandleFunc("", contactController.CreateContact).Methods("POST")
	userRouter.HandleFunc("", contactController.GetAllContacts).Methods("GET")
	userRouter.HandleFunc("/{contactId}", contactController.GetContactByID).Methods("GET")
	userRouter.HandleFunc("/{contactId}", contactController.UpdateContactById).Methods("PUT")
	userRouter.HandleFunc("/{contactId}", contactController.DeleteContactByID).Methods("DELETE")

	// ----------------------------------------------------------------------------------------------------------------------------------

	contactDetailRouter := router.PathPrefix("/users/{userId}/contacts/{contactId}/details").Subrouter()
	contactDetailRouter.Use(security.MiddlewareContact)

	contactDetailRouter.HandleFunc("", contactDetailController.CreateContactDetail).Methods("POST")
	contactDetailRouter.HandleFunc("", contactDetailController.GetAllContactDetails).Methods("GET")
	contactDetailRouter.HandleFunc("/{ContactDetailId}", contactDetailController.GetContactDetailById).Methods("GET")
	contactDetailRouter.HandleFunc("/{contactDetailId}", contactDetailController.UpdateContactDetail).Methods("PUT")
	contactDetailRouter.HandleFunc("/{contactDetailId}", contactDetailController.DeleteContactDetail).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":2611", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
