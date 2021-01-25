package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
)

type typeAssertionError struct {
	expected reflect.Type
	found    reflect.Type
}

func NewTypeAssertionError(exp, found interface{}) error {
	return typeAssertionError{
		expected: reflect.TypeOf(exp),
		found:    reflect.TypeOf(found),
	}
}

func (e typeAssertionError) Error() string {
	return fmt.Sprintf("interface conversion: interface is '%v' not '%v'", e.found, e.expected)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assert-panic", TypeAssertingCatchMiddleware(AssertPanic))
	mux.HandleFunc("/random-panic", TypeAssertingCatchMiddleware(RandomPanic))

	srv := httptest.NewServer(mux)

	data := get(srv, "/assert-panic")
	fmt.Println("Client resp when the panic is a type assertion")
	fmt.Println(string(data))

	fmt.Println()

	empty := get(srv, "/random-panic")
	fmt.Println("Client resp when the panic is a random panic")
	fmt.Println(string(empty))
}

func get(srv *httptest.Server, endpoint string) []byte {
	resp, err := srv.Client().Get(srv.URL + endpoint)
	if err != nil {
		return []byte{}
	}
	data, _ := ioutil.ReadAll(resp.Body)
	return data
}

func TypeAssertingCatchMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					// typeAssertionError will be a middleware specific error type
					if errors.As(err, &typeAssertionError{}) {
						// Log this
						// TODO: xerrors keeps the stack trace right?
						fmt.Println("recovered:", err.Error())

						// Send an internal error to the client
						// Probably also include more information in the 'details' section
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("Internal Server error"))
						return
					}
				}
				// Just panic as you normally would, we don't know what this is
				// TODO: Still send an internal server error to client
				log.Panic(r)
			}
		}()

		next(w, r)
	}
}

func AssertPanic(w http.ResponseWriter, r *http.Request) {
	var x interface{}
	_, ok := x.(string)
	if !ok {
		panic(NewTypeAssertionError(string(""), x))
	}
	w.Write([]byte("Should never happen"))
}

func RandomPanic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CALLED")
	panic("Something else")
	w.Write([]byte("Should never happen"))
}

//// DevURLCtx gets the dev URL resource from the context.
//func DevURLCtx(ctx context.Context) *database.DevURL {
//	v, ok := ctx.Value(devURLCtxKey{}).(*database.DevURL)
//	if !ok {
//		panic(typeAssertionError{
//			expected: &database.DevURL{},
//			found:    ctx.Value(devURLCtxKey{}),
//		})
//	}
//	return v
//}
