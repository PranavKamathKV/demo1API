package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"regexp"

)
type keyPair struct{
	key interface{}
	value interface{}
}

var redis []keyPair

//Default page
func homePage(w http.ResponseWriter, r* http.Request){
	fmt.Fprint(w, "Did it just work!! Yusss")
	fmt.Println("Homepage")

}

// Method to handle put requests
func putPair(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	varKey:= vars["key"]
	varValue:= vars["value"]
	redis = append(redis,keyPair{key:varKey, value:varValue})
	fmt.Fprint(w, "Put operation:\n"+"Key:", varKey,"\tValue:", varValue)
	fmt.Println(redis)

}

// Method to handle get requests
func getPair(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var flag int
	varKey:=vars["id"]
	fmt.Println(varKey)
	for i:= range redis {
		if redis[i].key == varKey {
			fmt.Fprint(w, "Key:"+varKey, "Value:",redis[i].value)
			fmt.Println("We are here")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Fprint(w, "Key not found")
	}

}

// Method to handle delete requests
func deletePair(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	varKey:= vars["id"]
	var flag int
	for i:= range redis{
		if redis[i].key == varKey{
			redis[i] = redis[len(redis) -1]
			redis[len(redis) -1] = keyPair{nil, nil }
			redis = redis[:len(redis) -1]
			flag = 1
			fmt.Fprint(w, "Key to delete"+varKey+" has been deleted")
			break
		}
	}
	if flag == 0{
		fmt.Fprint(w, "Key not found to delete")
	}
}

// Method to count the pairs
func countPair(w http.ResponseWriter, r * http.Request){
	var count int
	for i:= range redis{
		fmt.Println(i)
		count= count + 1

	}
	fmt.Fprint(w, "Count:", count)


}

// Method to count the given pattern key
func countPairID(w http.ResponseWriter, r * http.Request){
    vars := mux.Vars(r)
    varKey := vars["id"]
    var count int
    for i:=range redis{
    	match, _:= regexp.MatchString(varKey, redis[i].key.(string))
    	if match ==true{
    		count = count + 1
		}
	}
    fmt.Fprint(w, "Count with id"+varKey+ " is ", count)
}

// Method to get all the keypairs
func getAll(w http.ResponseWriter, r *http.Request){
	for i:= range redis{
		fmt.Fprint(w, "Key:", redis[i].key, "Value:", redis[i].value, "\n")
	}
}

//Method to handle routes
func handleRequest(){
	myroute := mux.NewRouter().StrictSlash(true)
	myroute.HandleFunc("/", homePage)
	myroute.HandleFunc("/GET/{id}", getPair)
	myroute.HandleFunc("/PUT/{key}/{value}", putPair)
	myroute.HandleFunc("/DELETE/{id}", deletePair)
	myroute.HandleFunc("/COUNT", countPair)
	myroute.HandleFunc("/COUNT/{id}", countPairID)
	myroute.HandleFunc("/all", getAll)
	http.ListenAndServe(":8081", myroute)

}

// THE MAIN
func main(){

	handleRequest()

}