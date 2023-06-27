package session

import (
	"time"
	"sync"
	"github.com/avenssi/"
)

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB(){

}

func GenerateNewSessionId(un string) string{

}

func IsSessionExpired(sid string) (string, bool){

}