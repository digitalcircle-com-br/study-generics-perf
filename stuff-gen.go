package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func DecGen[T any](r *http.Request) (in T, err error) {
	ptrin := new(T)
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(ptrin)
	in = *ptrin
	return
}

func EncGen[T any](w http.ResponseWriter, out T) error {
	return json.NewEncoder(w).Encode(out)
}

func AdaptGen[TIN any, TOUT any](f func(context.Context, TIN) (TOUT, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in, err := DecGen[TIN](r)
		if err != nil {
			log.Printf("Error at AdaptGen: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		out, err := f(r.Context(), in)
		if err != nil {
			log.Printf("Error at AdaptGen: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = EncGen(w, out)
		if err != nil {
			log.Printf("Error Decoding: %s", err)
		}
	}
}
