package routes

import (
	"context"
	"fmt"
	"net/http"

	"Abhinavbhar/dub.sh/redis"

	"github.com/gorilla/mux"
)

func BaseUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["value"]

	client := redis.RedisDatabase()

	// Fetch the URL from Redis
	url, err := client.Get(context.Background(), value).Result()

	fmt.Print(err)
	if err != nil {
		// Handle the error, e.g., URL not found
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
