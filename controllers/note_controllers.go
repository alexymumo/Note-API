package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/alexymumo/models"
	"github.com/alexymumo/responses"
	"github.com/gorilla/mux"
)

func (server *Server) CreateNote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	note := models.Note{}
	err = json.Unmarshal(body, &note)
	if err != nil {
		return
	}

	output, err := note.SaveNote(server.DB)
	if err != nil {
		return

	}
	responses.JSON(w, http.StatusCreated, output)
}

func (server *Server) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	note := models.Note{}
	err = server.DB.Debug().Model(models.Note{}).Where("note_id = ?", id).Take(&note).Error
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, err)
		return
	}

	_, err = note.DeleteNote(server.DB, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, "")
}

func (server *Server) UpdateNote(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server working")

}

func (server *Server) FindNotes(w http.ResponseWriter, r *http.Request) {
	note := models.Note{}
	notes, err := note.GetAllNotes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, notes)
}
