package event

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type EventServiceHandler struct{}

func (eh *EventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FindEventHandler: called")
	fmt.Println(r)
	json.NewEncoder(w).Encode(r)
}

func (eh *EventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("AllEventHandler called\n")
	//fmt.Println(r)
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(map[string]interface{}{"request": r.Accept-Encoding})

}

func (eh *EventServiceHandler) AddNewEventHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("AddNewEventHandler called\n")
}
