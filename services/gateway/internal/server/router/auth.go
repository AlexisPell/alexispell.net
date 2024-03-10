package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexispell/creep/gateway/internal/config"
)

func HelloWorldAuth(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("%s", config.Cfg.AuthUrl))
    if err != nil {
        fmt.Println("Error making request:", err)
        fmt.Fprintf(w, "Error making request: %s", err)
    }
    defer resp.Body.Close()

    // Decode JSON response into an interface{} type
    var data interface{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Println("Error decoding JSON:", err)
				fmt.Fprintf(w, "Error decoding JSON: %s", err)
    }

    // Print the parsed data
    fmt.Println("Parsed Data:", data)
		fmt.Fprintf(w, "Parsed Data: %s", data)
}

func beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("id")
	resp, err := http.Get(fmt.Sprintf("%s/auth/%s", config.Cfg.AuthUrl, provider))
	if err != nil {
			fmt.Println("Error making request:", err)
			fmt.Fprintf(w, "Error making request: %s", err)
	}
	defer resp.Body.Close()

	// Decode JSON response into an interface{} type
	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			fmt.Println("Error decoding JSON:", err)
			fmt.Fprintf(w, "Error decoding JSON: %s", err)
	}

	// Print the parsed data
	fmt.Println("Parsed Data:", data)
	fmt.Fprintf(w, "Parsed Data: %s", data)
}