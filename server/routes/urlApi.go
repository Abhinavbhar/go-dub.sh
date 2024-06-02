package routes

import (
	"Abhinavbhar/dub.sh/redis"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Url string `json:"url"`
}

func Url(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	uri := generateRandomString(3)
	finalUrl := "http://localhost:8080/url/" + uri

	client := redis.RedisDatabase()
	ctx := r.Context()

	// Store the actual URL with the `uri` as the key
	if err := client.Set(ctx, uri, data.Url, 0).Err(); err != nil {
		http.Error(w, "Failed to store URL in Redis", http.StatusInternalServerError)
		return
	}

	// Optionally, set an expiration time for the key
	if _, err := client.Expire(ctx, uri, time.Hour).Result(); err != nil {
		fmt.Println("Failed to set expiration time for key:", err)
	}

	response := map[string]string{"finalUrl": finalUrl}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
