package middleware

import "net/http"

func Recovery(h http.Handler) http.Handler {
fn := func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			// エラーが発生した場合は、500 Internal Server Errorを返す
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()
// ServeHTTPメソッドを呼び出す
	h.ServeHTTP(w, r)
}
return http.HandlerFunc(fn)
}


