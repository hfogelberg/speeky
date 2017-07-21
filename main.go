package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// voice := "en-GB_KateVoice"
	voice := "en-US_MichaelVoice"
	// voice := "en-US_AllisonVoice"
	// voice := "es-ES_LauraVoice"
	str := "Hello! How are you doing?"
	// str := "Hola! Buenos dias! Qúe tal?"

	text := strings.Replace(str, " ", "+", -1)
	url := fmt.Sprintf("https://stream.watsonplatform.net/text-to-speech/api/v1/synthesize?accept=audio/mp3&voice=%s&text=%s", voice, text)

	// Create output file
	output, err := os.Create("talk2.mp3")
	if err != nil {
		log.Printf("Error creating wav file %s\n", err.Error())
		return
	}

	// Build request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error building request %s \n", err.Error())
		return
	}

	user := os.Getenv("WATSON_USER_ID")
	pwd := os.Getenv("WATSON_PASSWORD")

	req.SetBasicAuth(user, pwd)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making http request %s\n", err.Error())
		return
	}

	defer resp.Body.Close()

	if _, err := io.Copy(output, resp.Body); err != nil {
		log.Printf("Error copying output %s\n", err.Error())
		return
	}

	fmt.Println("File talk.wav downloaded!")

}
