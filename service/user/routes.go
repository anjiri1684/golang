package user

import (
	"net/http"

	"github.com/anjiri1684/ecom/service/auth"
	"github.com/anjiri1684/ecom/types"
	"github.com/anjiri1684/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){}



func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	//get json payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if the user exists
_, err:=h.store.GetUserById(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest,  payload.Email)
		return
	}

	hasshedPassword, err:= auth.HasshedPassword(payload.Password)

	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	//if not create a new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hasshedPassword,
	})

	utils.WriteJSON(w, http.StatusInternalServerError, err)

}