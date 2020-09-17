package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func server() {
	router := mux.NewRouter()

	//api server sisaldab ühte endpointi nimega */endpoint, mis tagastab
	//JSON kujul legaalse vastuse.
	//JSON vastus on lahti seletatud üleval
	router.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"võti_set\":3, \"võti_hset\":{\"väli_hset_1\":5, \"väli_hset_2\":3}}"))
	})

	if err := http.ListenAndServe("127.0.0.1:7000", router); err != nil {
		panic(err)
	}
}

func main() {
	server()
}
