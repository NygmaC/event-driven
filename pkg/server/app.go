package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var mapSessions = make(map[string]Session)

type Session struct {
	id     string
	client *websocket.Conn
}

type Message struct {
	From    string
	Target  string
	Message string
}

// Setup de endpoints
func setupRoutes() {
	http.HandleFunc("/ws/{session}", wsEndpoint)
	http.HandleFunc("/list", listAllSessions)
	http.HandleFunc("/send", sendMessage)
}

// Send message to client
func sendMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  Get target connection in memory
	sessionTarget := mapSessions[msg.Target]

	log.Println("Enviar mensagem")
	log.Println("Target Id: " + sessionTarget.id)

	b := []byte(msg.Message)

	// Write message in client connection
	if err := sessionTarget.client.WriteMessage(1, b); err != nil {
		log.Println(err)
		return
	}

}

// Read Message received by the client
func readMessage(ws *websocket.Conn) {
	for {

		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(message))
		b := []byte("texto de teste")

		if err := ws.WriteMessage(1, b); err != nil {
			log.Println(err)
			return
		}

	}
}

// List sessions in memory
func listAllSessions(w http.ResponseWriter, r *http.Request) {

	for key, element := range mapSessions {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

}

// Endpoint de acesso dos clients to websocket
func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("session")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Convert the http request to WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	// Add client session in memory
	mapSessions[id] = Session{id, ws}

	readMessage(ws)
}

func main() {
	fmt.Println("Server UP!!!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
