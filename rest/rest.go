package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abc7468/roycoin/blockchain"
	"github.com/gorilla/mux"
)

var port string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See A block",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A block",
			Payload:     "data:string",
		},
	}
	fmt.Println(data)
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
		// json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())

	case "POST":

		//{"data":"my block data"}
		return
		// var addBlockBody addBlockBody
		// utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		// blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		// rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	// fmt.Println(URLDescription{
	// 	URL:         "/",
	// 	Method:      "GET",
	// 	Description: "See Documentation",
	// })
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
