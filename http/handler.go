package http

import (
	"encoding/json"
	"net/http"

	"github.com/anshumancodes/ShardCache/cache"
)

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func HelloFrom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := &Response{Message: "http server is running sucessfully!"}
	json.NewEncoder(w).Encode(response)
}

// using this basically we can access the cache from the handler and manage it
type Handler struct {
	cache *cache.Cache // pointer to the cache
}

// this is a constructor function for the Handler struct
func NewHandler(c *cache.Cache) *Handler {
	return &Handler{
		cache: c,
	}
}

// get a Handler struct
// then read for the path key from request path key
// then using the handler access the cache and use cache.get to retrive value by using the key
// if the key is not found, return a 404 error
// if the key is found, return the value

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// read the key from the request path
	key := r.PathValue("key")
	// then using the handler access the cache and use cache.get to retrive value by using the key
	value, ok := h.cache.Get(key)

	// returning 404
	if !ok {
		http.Error(w, "key not found", http.StatusNotFound)
	}

	// set json header
	w.Header().Set("Content-Type", "application/json")

	// return key value pair

	json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})

}

// will pass the Handler struct
// get key from request
// we will decode the value from request body
// if error return a 400 error
// if not a error , will use cache.set to set the value in the cache
func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	var body struct {
		Value string `json:"value"`
	}

	// decode the request body into the body struct
	err := json.NewDecoder(r.Body).Decode(&body)

	// check for errors
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// set the value in the cache using cache package's set method
	h.cache.Set(key, body.Value)

	w.Header().Set("Content-Type", "application/json")

	// return the key value pair
	json.NewEncoder(w).Encode(map[string]string{"key": key, "value": body.Value})

}
