package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func DecInt(r *http.Request, in interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(in)
}
func EncInt(w http.ResponseWriter, out interface{}) error {
	return json.NewEncoder(w).Encode(out)
}

func AdaptInt(inTp interface{}, f func(context.Context, interface{}) (interface{}, error)) http.HandlerFunc {
	tp := reflect.TypeOf(inTp)



	
	return func(w http.ResponseWriter, r *http.Request) {
		in := reflect.New(tp).Interface()
		err := DecInt(r, in)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("%s", err.Error())
			return
		}
		out, err := f(r.Context(), in)
		if err != nil {
			log.Printf("Error at AdaptInt: %s",err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = EncInt(w, out)
		if err != nil {
			log.Printf("Error Decoding: %s", err)
		}
	}
}
