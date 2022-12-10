package http

import "net/http"

type httpHandlerWithError func(http.ResponseWriter, *http.Request) error

func (fn httpHandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        //  http.Error(w, err.Message, err.StatusCode)
        // handle error
    }
}