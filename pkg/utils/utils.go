package utils

import (
	"net/http"
	"io"
	"encoding/json"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err :=  io.ReadAll(r.Body); err != nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}