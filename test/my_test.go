package test

import (
	bytes2 "bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"
)

type Message struct {
	Log       string `json:"log"`
	Timestamp int64  `json:"timestamp"`
}

func TestSendLogs(t *testing.T) {

	for {

		msg := Message{
			Log:       "[2024-05-24 10:30:15] [INFO] This is a sample log message.",
			Timestamp: time.Now().UnixMilli(),
		}

		bytes, err := json.Marshal(msg)
		if err != nil {
			log.Println(err)
			return
		}

		var buffer bytes2.Buffer
		buffer.Write(bytes)

		http.Post("http://localhost:9999/log/http", "application/json", &buffer)

		time.Sleep(time.Second * 3)

	}

}
