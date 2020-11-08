package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pet_store_rest_api/interfaces"
	"pet_store_rest_api/models"
	"pet_store_rest_api/responses"
	"strconv"

	"github.com/gorilla/mux"
)

type PetController struct {
	interfaces.IPetRepository
}

func (controller *PetController) CreateNewPet(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Pet struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	var newPet models.Pet
	err = json.Unmarshal(reqBody, &newPet)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = controller.AddNewPet(newPet)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusCreated, newPet)
}

func (controller *PetController) ReadPetInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.ParseInt(vars["petId"], 10, 64)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	pet, err := controller.GetPet(petId)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, pet)
}

func (controller *PetController) ReadPetsInfoByStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for i, val := range r.Form {
		if i == "status" {
			pets := controller.GetPetByStatus(val)
			responses.JSON(w, http.StatusOK, pets)
		}
	}
	responses.JSON(w, http.StatusBadRequest, nil)
}

func (controller *PetController) UpdatePetData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusMethodNotAllowed, err)
		return
	}
	var petInfo models.Pet
	err = json.Unmarshal(reqBody, &petInfo)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = controller.UpdatePet(petInfo)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, petInfo)
}

func (controller *PetController) UpdatePetWithFormData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.ParseInt(vars["petId"], 10, 64)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	r.ParseForm()
	petInfoPtr, err := controller.GetPet(petId)
	petInfoVal := petInfoPtr.(models.Pet)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	formValue := r.FormValue("name")
	if formValue != "" {
		petInfoVal.Name = formValue
	}
	formValue = r.FormValue("status")
	if formValue != "" {
		petInfoVal.Name = formValue
	}
	// Making Assumption that category_id and category_name can be supplied with form.
	formValue = r.FormValue("category_id")
	if formValue != "" {
		petInfoVal.Category.Id, err = strconv.ParseInt(formValue, 10, 64)
		if err != nil {
			log.Println(err)
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
	}
	formValue = r.FormValue("category_name")
	if formValue != "" {
		petInfoVal.Category.Name = formValue
	}

	// Making Assumption that tag_id and tag_name can be supplied with form.
	var tag models.Tag
	formValue = r.FormValue("tag_id")
	if formValue != "" {
		tag.Id, err = strconv.ParseInt(formValue, 10, 64)
		if err != nil {
			log.Println(err)
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
	}
	formValue = r.FormValue("tag_name")
	if formValue != "" {
		tag.Name = formValue
	}

	err = controller.UpdatePet(petInfoVal)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, petInfoVal)
}

// DeletePet deletes a Pet
func (controller *PetController) DeletePetData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.ParseInt(vars["petId"], 10, 64)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = controller.DeletePet(petId)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, models.ApiResponse{
		Code:    int32(http.StatusOK),
		Type:    "success",
		Message: "Pet with id " + vars["petId"] + " deleted"})
}
