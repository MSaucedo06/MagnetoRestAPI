package main

import (
	"encoding/json"
	"fmt"
	"log"
	
	"net/http"
	"regexp"
	"github.com/gorilla/mux"
)



type test_struct struct {
	Dna []string
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")	
}

func newEvent(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t)
	var result bool = isMutant(t.Dna)
	
	if result {
		rw.WriteHeader(http.StatusOK)
	} else { 
		rw.WriteHeader(http.StatusForbidden)
	}	

	
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hola")
}

func validDNA(dna []string) bool {
	var valid int = 0	
	re := regexp.MustCompile(`[BDEFHIJKLMNOPQRSUVWXYZ]`)   

	for _, chain := range dna {		
		x := re.Match([]byte(chain))		
		if x {
			fmt.Println("ERROR: It is not a nitrogenous base of DNA:", chain)
			valid++
		}
		
	}

	if valid > 0 {
		return false
	} else { 
		return true
	}
}

func searchMutanChain(dna []string) int {
	var mutant int = 0 

	for _, chain := range dna {
		x := charRepeat(chain)
		if x {
			mutant++
			fmt.Println("Is Mutant:", chain)			
		}
		
	}
	return mutant
}

func isMutant(dna []string) bool {
	var result int = 0
	if validDNA(dna){
		h:= searchMutanChain(dna)
		result = h
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

func charRepeat(chain string) bool {
	repeatCount := 1
	thresh := 4
	lastChar := ""
	mutantDNA := false

	for _, r := range chain {
		c := string(r)
		if c == lastChar {
			repeatCount++
			if repeatCount == thresh {
				mutantDNA = true
			}
		} else {
			repeatCount = 1
		}
		lastChar = c
	}

	return mutantDNA
}

func toMatrix() {

	a := make([][]uint8, 6)
	for i := range a {
   		a[i] = make([]uint8, 6)
}
}



func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/mutant", newEvent).Methods("POST")
	router.HandleFunc("/stats", getAllEvents).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}


