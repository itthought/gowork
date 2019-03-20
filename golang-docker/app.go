package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/itthought/gowork/golang-docker/model"
)


func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var request model.SocketRoomStatusRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	var response = model.NewCreateSocketRoomResponse()
	response.SetAddress("10.150.236.140")
	response.SetGameStatus("connecting")
	response.SetPort(6555)
	response.SetProtocol("UDP")
	response.SetRoomId(request.RoomId)
	response.SetSocketStatus("created")
	json.NewEncoder(w).Encode(response)

	/*responseJson, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Could not find details %s\n",err)
	}
	w.Write(responseJson)*/

}


// Create Room
func createRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var request model.CreateSocketRoomRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	var response = model.NewCreateSocketRoomResponse()
	response.SetAddress("10.150.236.140")
	response.SetGameStatus("connecting")
	response.SetPort(6555)
	response.SetProtocol("UDP")
	response.SetRoomId(request.RoomId)
	response.SetSocketStatus("created")
	json.NewEncoder(w).Encode(response)

}

// Update a single book
func deleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	/*params:= mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)*/

}

// Create Room
func joinRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var request model.JoinSocketRoomRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	var response = model.NewCreateSocketRoomResponse()
	response.SetAddress("10.150.236.140")
	response.SetGameStatus("connecting")
	response.SetPort(6555)
	response.SetProtocol("UDP")
	response.SetRoomId(request.RoomId)
	response.SetSocketStatus("created")
	json.NewEncoder(w).Encode(response)

}

// Delete a single book
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Content-Type", "Application/json")
	var request model.CreateSocketRoomRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	var response = model.NewCreateSocketRoomResponse()
	response.SetAddress("10.150.236.140")
	response.SetGameStatus("connecting")
	response.SetPort(6555)
	response.SetProtocol("UDP")
	response.SetRoomId(request.RoomId)
	response.SetSocketStatus("created")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc("/api/create", createRoom).Methods("POST")
	r.HandleFunc("/api/status", getStatus).Methods("POST")
	r.HandleFunc("/api/delete", deleteRoom).Methods("POST")
	r.HandleFunc("/api/join", joinRoom).Methods("POST")
	r.HandleFunc("/api/update", update).Methods("POST")
/*	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")*/

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}