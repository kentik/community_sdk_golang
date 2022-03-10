package utils

import (
	"encoding/json"
	"log"
)

func LogPayload(l bool, msg string, payload interface{}) {
	if l {
		if payload == "" {
			log.Println(msg)
		} else {
			jsonData, err := json.Marshal(&payload)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(msg, string(jsonData))
		}
	}
}
