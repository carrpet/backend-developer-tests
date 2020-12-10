package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/carrpet/backend-developer-tests/rest-service/pkg/models"
)

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	peopleJSON, err := json.Marshal(models.AllPeople())

	if err != nil {
		log.Printf("Error marshalling all people: %s", err.Error())
		http.Error(w, "Internal error encountered", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleJSON)
}

func peopleByIDHandler(w http.ResponseWriter, r *http.Request) {
	uuid, err := getID(w, r)
	if err != nil {
		log.Printf("Error retrieving UUID: %s", err.Error())
		return
	}

	people, err := models.FindPersonByID(uuid)
	if err != nil {
		errorMsg := fmt.Sprintf("Could not find person with the specified ID: %s", uuid.String())
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	peopleJSON, err := json.Marshal(people)
	if err != nil {
		log.Printf("Error marshalling response: %s", err.Error())
		http.Error(w, "Error constructing response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleJSON)

}

func getID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	if idString, ok := vars["id"]; ok {
		id, err := uuid.FromString(idString)
		if err != nil {
			http.Error(w, "Supplied id was an invalid UUID", http.StatusBadRequest)
			return uuid.UUID{}, err
		}
		return id, nil
	}
	errorMsg := "Server could not process id properly"
	http.Error(w, errorMsg, http.StatusInternalServerError)
	return uuid.UUID{}, errors.New(errorMsg)
}

func peoplePhoneHandler(w http.ResponseWriter, r *http.Request) {
	number, err := getPhoneNumber(w, r)

	if err != nil {
		log.Println(err.Error())
		return
	}

	people := models.FindPeopleByPhoneNumber(number)
	peopleJSON, err := json.Marshal(people)

	if err != nil {
		log.Printf("Error marshalling peopleJSON: %s", err.Error())
		http.Error(w, "Internal error encountered", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleJSON)

}

func getPhoneNumber(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	if pNumber, ok := vars["phone_number"]; ok {
		return pNumber, nil
	}
	errorMsg := "Internal error retrieving phone number"
	http.Error(w, errorMsg, http.StatusInternalServerError)
	return "", errors.New(errorMsg)

}

func peopleNameHandler(w http.ResponseWriter, r *http.Request) {
	name, err := getName(w, r)

	if err != nil {
		log.Printf("Error retrieving name: %s", err.Error())
		return
	}

	people := models.FindPeopleByName(name.first, name.last)
	peopleJSON, err := json.Marshal(people)

	if err != nil {
		log.Printf("Error marshalling peopleJSON in name handler: %s", err.Error())
		http.Error(w, "Internal error encounterd", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleJSON)

}

type name struct {
	first string
	last  string
}

func getName(w http.ResponseWriter, r *http.Request) (name, error) {
	vars := mux.Vars(r)
	if fName, ok := vars["first_name"]; ok {
		if lName, ok := vars["last_name"]; ok {
			return name{first: fName, last: lName}, nil
		}
	}
	errorMsg := "Internal error retrieving name"
	http.Error(w, errorMsg, http.StatusInternalServerError)
	return name{}, errors.New(errorMsg)
}

func peopleSelectorHandler(w http.ResponseWriter, r *http.Request) {
	fName, lName, pNumber := r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name"), r.URL.Query().Get("phone_number")
	if fName != "" {
		if lName != "" {
			peopleJson, err := json.Marshal(models.FindPeopleByName(fName, lName))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(peopleJson)
			return
		}
	}
	if pNumber != "" {
		peopleJson, err := json.Marshal(models.FindPeopleByPhoneNumber(pNumber))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(peopleJson)
		return
	}
	http.Error(w, "Bad query string", http.StatusBadRequest)

}

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	r := mux.NewRouter()
	r.HandleFunc("/people", peopleHandler).Methods("GET")
	r.HandleFunc("/people/{id:[a-z0-9-]+}", peopleByIDHandler).Methods("GET")

	r.HandleFunc("/people/", peopleNameHandler).Methods("GET").
		Queries("first_name", "{first_name:[a-zA-Z]+}", "last_name", "{last_name:[a-zA-Z]+}")

	r.HandleFunc("/people/", peoplePhoneHandler).Methods("GET").
		Queries("phone_number", "")

	log.Println("Starting people server on :8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
