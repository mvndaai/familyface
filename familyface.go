package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var people = []*person{}
var idCount = 0
var port = 9090

type person struct {
	ID          int               `json:"id"`
	MomID       int               `json:"mom_id"`
	DadID       int               `json:"dad_id"`
	ChildrenIDs []int             `json:"children_ids"`
	Marriages   []*marriage       `json:"marriages"`
	Details     map[string]string `json:"details"`
}

type marriage struct {
	SpouseID int  `json:"spouse_id"`
	Current  bool `json:"current"`
}

func main() {
	http.HandleFunc("/face", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/dump", dumpPeople)
	http.HandleFunc("/people/get", getPeopleJSON)
	http.HandleFunc("/people/search", searchPeopleJSON)
	http.HandleFunc("/person/get", getPersonJSON)
	http.HandleFunc("/person/add", addPerson)
	http.HandleFunc("/person/add/mom", addMom)
	http.HandleFunc("/person/add/dad", addDad)
	http.HandleFunc("/person/add/marriage", addMarriage)
	http.HandleFunc("/status", status)
	http.HandleFunc("/", handler)

	fmt.Printf("Starting familyface server on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func dumpPeople(w http.ResponseWriter, r *http.Request) {
	people = []*person{}
	idCount = 0
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 handler"))
}

func status(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("People: %v", len(people))
	w.Write([]byte(str))
}

func nextID() int {
	idCount++
	return idCount
}

func connectMom(mom, child *person) bool {
	if mom == nil || child == nil {
		return false
	}
	child.MomID = mom.ID
	mom.ChildrenIDs = append(mom.ChildrenIDs, child.ID)
	return true
}

func connectDad(dad, child *person) bool {
	if dad == nil || child == nil {
		return false
	}
	child.DadID = dad.ID
	dad.ChildrenIDs = append(dad.ChildrenIDs, child.ID)
	return true
}

func getPersonByID(ID int) *person {
	if ID == 0 {
		return nil
	}
	for _, p := range people {
		if ID == p.ID {
			return p
		}
	}
	return nil
}

func addMom(w http.ResponseWriter, r *http.Request) {
	params, ok := getBodyMap(w, r)
	if !ok {
		return
	}
	momID, _ := strconv.Atoi(params["mom_id"])
	childID, _ := strconv.Atoi(params["child_id"])
	mom := getPersonByID(momID)
	child := getPersonByID(childID)
	connectMom(mom, child)
}

func addDad(w http.ResponseWriter, r *http.Request) {
	params, ok := getBodyMap(w, r)
	if !ok {
		return
	}
	dadID, _ := strconv.Atoi(params["dad_id"])
	childID, _ := strconv.Atoi(params["child_id"])
	dad := getPersonByID(dadID)
	child := getPersonByID(childID)
	connectDad(dad, child)
}

func getPeopleJSON(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(people)
	if err != nil {
		log.Printf("Failed to marshal json: %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(b))
	w.Write(b)
}

func searchPeopleJSON(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query()["term"][0]
	term = strings.ToLower(term)
	hasTerm := []*person{}

	if len(term) > 0 {
		for _, p := range people {
			for _, v := range p.Details {
				if strings.Contains(strings.ToLower(v), term) {
					hasTerm = append(hasTerm, p)
				}
			}
		}
	}

	b, err := json.Marshal(hasTerm)
	if err != nil {
		log.Printf("Failed to marshal json: %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(b))
	w.Write(b)

}

func getPersonJSON(w http.ResponseWriter, r *http.Request) {
	personID, err := strconv.Atoi(r.URL.Query()["person_id"][0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.URL.Query())
	fmt.Println("getPersonJSON id", personID)
	p := getPersonByID(personID)
	b, err := json.Marshal(p)
	if err != nil {
		log.Printf("Failed to marshal json: %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(b))
	w.Write(b)
}

func addMarriage(w http.ResponseWriter, r *http.Request) {
	params, ok := getBodyMap(w, r)
	if !ok {
		return
	}
	personID, _ := strconv.Atoi(params["person_id"])
	spouseID, _ := strconv.Atoi(params["spouse_id"])
	person := getPersonByID(personID)
	spouse := getPersonByID(spouseID)
	createMarriage(person, spouse)
	fmt.Printf("%#v\n", person)
}

func createMarriage(person, spouse *person) {
	for i := range person.Marriages {
		person.Marriages[i].Current = false
	}
	for i := range spouse.Marriages {
		spouse.Marriages[i].Current = false
	}
	mp := marriage{SpouseID: spouse.ID, Current: true}
	ms := marriage{SpouseID: person.ID, Current: true}
	person.Marriages = append(person.Marriages, &mp)
	spouse.Marriages = append(spouse.Marriages, &ms)
}

func getBodyMap(w http.ResponseWriter, r *http.Request) (map[string]string, bool) {
	tmp := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		log.Printf("Could not unmarshal json - %q", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return nil, false
	}
	return tmp, true
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	params, ok := getBodyMap(w, r)
	if !ok {
		return
	}
	p := person{ID: nextID()}
	p.Details = map[string]string{"surname": fmt.Sprintf("Name%v", idCount), "givennames": "No"}
	for k, v := range params {
		fmt.Println(k, v)
		if len(v) > 0 {
			p.Details[k] = v
		}
	}
	people = append(people, &p)
	//fmt.Printf("%#v\n", p)
	w.Write([]byte(fmt.Sprintf("%v", p.ID)))
}
