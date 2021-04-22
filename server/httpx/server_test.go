package httpx

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	router := chi.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("before1")
			next.ServeHTTP(writer, request)
			fmt.Println("after1")
			return
		})
	})
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("before2")
			next.ServeHTTP(writer, request)
			fmt.Println("after2")
			return
		})
	})
	router.Get("/get", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})
	http.ListenAndServe("127.0.0.1:8080", router)
}
