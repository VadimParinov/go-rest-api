package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api/cmd/internal/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.IndexHandler)
	router.GET(userURL, h.getUserByID)
	router.POST(usersURL, h.createUser)
	router.PUT(userURL, h.updateUser)
	router.DELETE(userURL, h.deleteUser)
	router.PATCH(userURL, h.patchUser)
}

func (h *handler) getUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Hello, %s", id)))
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("Create user"))
}

func (h *handler) updateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Update user"))
}

func (h *handler) deleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Delete user"))
}

func (h *handler) patchUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Patch user"))
}

func (h *handler) IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Get All users"))
}
