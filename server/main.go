package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	broker   = "tcp://localhost:1883"
	clientID = "mqtt_publisher"
	topic    = "cringecast"
	username = "admin"
	password = "pass"
)

var opts *mqtt.ClientOptions
var client mqtt.Client
var token mqtt.Token

type PostRequestBody struct {
	Query string `json:"query"`
	Lang  string `json:"lang"`
}

type UrlStrPostRequestBody struct {
	UrlStr string `json:"url"`
}

type Message struct {
	Command string `json:"command"`
	Payload string `json:"payload"`
}

type SayPayload struct {
	Query    string `json:"query"`
	Language string `json:"language"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	opts = mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID).SetUsername(username).SetPassword(password)
	client = mqtt.NewClient(opts)
	token = client.Connect()

	if token.Wait() && token.Error() != nil {
		log.Fatalf("Error during connection to MQTT: %s", token.Error())
	}

	http.HandleFunc("/say", handleSayRequest)
	http.HandleFunc("/play", handlePlayRequest)
	//http.HandleFunc("/stop", handleStopRequest)

	log.Println("Listening port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handleSayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := ""
	lang := ""

	if r.Method == http.MethodGet {
		query = r.URL.Query().Get("query")
		lang = r.URL.Query().Get("lang")
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error of reading request", http.StatusBadRequest)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(r.Body)

		var requestBody PostRequestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, "Error while decoding JSON", http.StatusBadRequest)
			return
		}

		query = requestBody.Query
		lang = requestBody.Lang
	}

	payloadObj := SayPayload{query, lang}
	payloadString, err := json.Marshal(payloadObj)
	if err != nil {
		log.Fatal(err)
	}
	obj := Message{"say", string(payloadString)}
	payload, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	token := client.Publish(topic, 0, false, payload)
	token.Wait()
}

func handlePlayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	urlStr := ""

	if r.Method == http.MethodGet {
		urlStr = r.URL.Query().Get("url")
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error of reading request", http.StatusBadRequest)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(r.Body)

		var requestBody UrlStrPostRequestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, "Error while decoding JSON", http.StatusBadRequest)
			return
		}

		urlStr = requestBody.UrlStr
	}

	obj := Message{"play", urlStr}
	payload, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	token := client.Publish(topic, 0, false, payload)
	token.Wait()
}

//func handleStopRequest(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	obj := Message{"stop", ""}
//	payload, err := json.Marshal(obj)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	token := client.Publish(topic, 0, false, payload)
//	token.Wait()
//}
