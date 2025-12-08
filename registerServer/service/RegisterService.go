package service

import (
	"encoding/json"
	"github.com/cxb116/ADX_ENGINE/registerServer/registry"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":80"
const ServicesUrl = "http://127.0.0.1:80"

type Register struct {
	registrations []registry.Registration
	mutex         *sync.RWMutex
}

func (this *Register) add(reg registry.Registration) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.registrations = append(this.registrations, reg)
	return nil
}

var reg = Register{
	registrations: make([]registry.Registration, 0),
	mutex:         new(sync.RWMutex),
}

type RegisterService struct{}

func (this *RegisterService) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var reg registry.Registration
		err := decode.Decode(&reg)
		if err != nil {
			log.Println("decode err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Adding Service: %s with URL: %s", reg.ServiceName, reg.ServiceURL)

	}
}
