package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	// healthzに関するルータを定義
	healthRouter(mux)
	todoRouter(mux, todoDB)
	return mux
}

// healthzに関するルータを定義
func healthRouter(mux *http.ServeMux) {
	healthz := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthz.ServeHTTP)

}

// todoに関するルータを定義
func todoRouter(mux *http.ServeMux, db *sql.DB) {
	todo := handler.NewTODOHandler(service.NewTODOService(db))//TODOHandlerを生成
	
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {//"/todos"に対するリクエストを処理する
		var err error
		switch r.Method {
		case http.MethodGet:
			//w は、HTTPレスポンスを書き込むための構造体
			//r は、HTTPリクエストを表す構造体
			err = responseJson(todo.ReadTodo(w, r))//ReadTodoを呼び出し、結果をJSONで返す
		case http.MethodPost:
			err = responseJson(todo.CreateTodo(w, r))
		case http.MethodPut:
			err = responseJson(todo.UpdateTodo(w, r))
		case http.MethodDelete:
			err = responseJson(todo.DeleteTodo(w, r))
		default:
			// TODO:エラーハンドリングする
		}
		//エラーがあった場合ログ出力して500を返す
		if err != nil {
			log.Println(err)
			responseJson(w, http.StatusInternalServerError, err)
		}
	})
}


func responseJson(w http.ResponseWriter, status int, response interface{}) error {
	//レスポンスを書き込む
	w.WriteHeader(status)
	//encodeして書き込む
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}
