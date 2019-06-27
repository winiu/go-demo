package core

import (
	"encoding/json"
	"io"
	"net/http"
)

var chain *Chain = GenerateChain()

func Run() {
	http.HandleFunc("/get", getBlockChain)
	http.HandleFunc("/write", writeBlockChain)
	http.ListenAndServe(":8888", nil)
}

func getBlockChain(w http.ResponseWriter, r *http.Request) {
	if bytes, error := json.Marshal(chain); error == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		io.WriteString(w, string(bytes))
	} else {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
}

func writeBlockChain(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	chain.SendData(blockData)
	getBlockChain(w, r)
}
