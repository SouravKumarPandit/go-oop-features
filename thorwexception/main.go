package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Its not exactly what it is just work around which should solve this issue of exception
// it will reduce code of error occurrences
func main() {
	http.HandleFunc("/", handleReq)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type item = struct {
	code    int
	message string
}

func handleReq(w http.ResponseWriter, req *http.Request) {
	var p = struct {
		Message string `json:"message"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item := processing(p.Message)

	_, _ = fmt.Fprintf(w, "hello \n--> "+p.Message, "\n\n Error : ", item)
}

func processing(message string) item {
	//you can buffer the channel
	//don't try to use same trick with panic and recover its not save
	throwable := make(chan item)

	go func() {
		switch message {
		case "hello":
			throwable <- item{
				code:    123,
				message: "You send Hello.We don't accept hello",
			}
			break
		case "welcome":
			throwable <- item{
				code:    124,
				message: "You send welcome.We don't accept welcome",
			}
			break
		default:
			throwable <- item{}
		}

	}()
	processResult := <-throwable
	return processResult
}
