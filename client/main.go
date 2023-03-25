package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var stopPlaying = false

type Message struct {
	Command string `json:"command"`
	Payload string `json:"payload"`
}

type SayPayload struct {
	Query    string `json:"query"`
	Language string `json:"language"`
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var obj Message
	if err := json.Unmarshal(msg.Payload(), &obj); err != nil {
		log.Println(err)
		return
	}
	switch obj.Command {
	case "say":
		var sayPayload SayPayload
		if err := json.Unmarshal([]byte(obj.Payload), &sayPayload); err != nil {
			log.Println(err)
			return
		}
		say(sayPayload)
	case "play":
		playAudio(obj.Payload)
		//case "stop":
		//	stopPlaying = true
	}
}

func main() {
	username := flag.String("mqtt-username", "admin", "MQTT Username")
	password := flag.String("mqtt-password", "pass", "MQTT Password")
	broker := flag.String("mqtt-url", "tcp://localhost:1883", "MQTT Broker URL")
	topic := flag.String("mqtt-topic", "cringecast", "MQTT Topic")

	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	rand.Seed(time.Now().UnixNano())
	clientID := fmt.Sprintf("mqtt_subscriber_%d", rand.Int())
	opts := mqtt.NewClientOptions().AddBroker(*broker).SetClientID(clientID).SetUsername(*username).SetPassword(*password)
	opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		log.Fatalf("Error during connection to MQTT: %s", token.Error())
	}

	token = client.Subscribe(*topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", *topic)

	select {}
}

func say(sayPayload SayPayload) {
	sentences := splitToSentences(sayPayload.Query, 100)

	for _, sentence := range sentences {
		urlStr := fmt.Sprintf(
			"https://translate.google.com.vn/translate_tts?ie=UTF-8&client=tw-ob&q=%s&tl=%s",
			url.QueryEscape(sentence),
			sayPayload.Language)
		playAudio(urlStr)
	}
}

func playAudio(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error loading audio file:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	decoder, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		fmt.Println("Error decoding audio file:", err)
		return
	}

	sampleRate := decoder.SampleRate() * 2

	context, err := oto.NewContext(int(sampleRate), 1, 2, 8192)
	if err != nil {
		fmt.Println("Error creating audio player:", err)
		return
	}
	defer func(context *oto.Context) {
		err := context.Close()
		if err != nil {

		}
	}(context)

	player := context.NewPlayer()

	buf := make([]byte, 8192)
	for {
		if stopPlaying {
			stopPlaying = false
			break
		}
		n, err := decoder.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading audio data:", err)
			return
		}

		// Записываем блоки аудио данных в плеер
		if _, err := player.Write(buf[:n]); err != nil {
			fmt.Println("Error playing audio data:", err)
			return
		}
	}
}

func splitToSentences(text string, maxLength int) []string {
	var sentences []string

	sentences = strings.Split(text, ".")

	for i := 0; i < len(sentences); i++ {
		sentences[i] = strings.TrimSpace(sentences[i])
	}

	sentences = removeEmpty(sentences)

	var merged []string
	var current string
	for i := 0; i < len(sentences); i++ {
		if len(current)+len(sentences[i])+2 <= maxLength {
			if current != "" {
				current += " "
			}
			current += sentences[i]
		} else {
			merged = append(merged, current)
			current = sentences[i]
		}
	}
	if current != "" {
		merged = append(merged, current)
	}

	return merged
}

func removeEmpty(s []string) []string {
	var result []string
	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
