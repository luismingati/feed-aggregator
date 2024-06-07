package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWitherror(w, http.StatusBadRequest, "something went wrong")
}
