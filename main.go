package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var t = flag.String("t", "Hello World", "Text to speek")
var l = flag.String("l", "us", "Language(es, us, uk, de, it, fr, jp, pt)")
var g = flag.String("g", "f", "Gender (m or f)")

func main() {
	flag.Parse()

	var voice string

	fmt.Printf("%s, %s, %s", *g, *l, *t)

	switch *l {
	case "pt":
		voice = "pt-BR_IsabelaVoice"
	case "jp":
		voice = "ja-JP_EmiVoice"
	case "fr":
		voice = "fr-FR_ReneeVoice"
	case "it":
		voice = "it-IT_FrancescaVoice"
	case "uk":
		voice = "en-GB_KateVoice"
	case "us":
		if *g == "m" {
			voice = "en-US_MichaelVoice"
		} else {
			voice = "en-US_LisaVoice"
		}
	case "es":
		if *g == "m" {
			voice = "es-ES_EnriqueVoice"
		} else {
			voice = "es-ES_LauraVoice"
		}
	}

	fmt.Println(voice)

	text := strings.Replace(*t, " ", "+", -1)
	url := fmt.Sprintf("https://stream.watsonplatform.net/text-to-speech/api/v1/synthesize?accept=audio/mp3&voice=%s&text=%s", voice, text)
	fmt.Println(url)

	// Create output file
	output, err := os.Create("speeky.mp3")
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
