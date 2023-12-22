package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// ここで、HealthzHandler構造体を定義している
type HealthzHandler struct{}

// ここで、NewHealthzHandler関数を定義している
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// w http.ResponseWriterは、HTTPレスポンスを書き込むための構造体
// r *http.Requestは、HTTPリクエストを表す構造体
// ServeHTTPメソッドは、HTTPリクエストを受け取り、HTTPレスポンスを返す
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := &model.HealthzResponse{
		Message: "OK",
	}
	//json.NewEncoder(w)は、HTTPレスポンスを書き込むための構造体を引数に取り、JSONエンコーダを返す
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}
