package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8090"`
		Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	} `yaml:"server"`
	Client struct {
		Books struct {
			Protocol string `yaml:"protocol" env:"PROTOCOL" env-default:"http://"`
			Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
			Port     string `yaml:"port" env:"PORT" env-default:"8091"`
			Endpoint string `yaml:"endpoint" env:"ENDPOINT" env-default:"/books"`
		} `yaml:"books"`
	} `yaml:"client"`
}

type ErrorResponse struct {
	Msg string `json:"message"`
}

type Book struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

var purchased []string

func main() {
	configLocation := "./config.yml"
	if len(os.Args) > 1 {
		configLocation = os.Args[1]
	}

	var cfg Config
	err := cleanenv.ReadConfig(configLocation, &cfg)

	if err != nil {
		log.Println("Error: ", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /purchase", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request made to POST /purchase")

		var titleBooks []string
		err := json.NewDecoder(r.Body).Decode(&titleBooks)

		if err != nil {
			log.Println("Error: ", err)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error reading request from user",
			})
			return
		}

		var titleBooksToBuy []string

		for _, titleBook := range titleBooks {
			bought := false
			for _, purchase := range purchased {
				if titleBook == purchase {
					bought = true
				}
			}
			if !bought {
				titleBooksToBuy = append(titleBooksToBuy, titleBook)
			}
		}

		url := cfg.Client.Books.Protocol + cfg.Client.Books.Host + cfg.Client.Books.Port + cfg.Client.Books.Endpoint
		req, err := http.NewRequest("GET", url, nil)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error creating request",
			})
			return
		}
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error doing the request",
			})
			return
		}

		defer res.Body.Close()
		body, readErr := io.ReadAll(res.Body)

		if readErr != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error reading response",
			})
			return
		}

		var books []Book

		if err := json.Unmarshal(body, &books); err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error parsing books json",
			})
			return
		}

		var purchasedBooks []Book = []Book{}

		for _, book := range books {
			for _, title := range titleBooksToBuy {
				if book.Title == title {
					purchasedBooks = append(purchasedBooks, book)
				}
			}
		}

		purchased = append(purchased, titleBooksToBuy...)

		json.NewEncoder(w).Encode(purchasedBooks)
	})

	mux.HandleFunc("GET /purchased", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request made to GET /purchased")

		url := cfg.Client.Books.Protocol + cfg.Client.Books.Host + cfg.Client.Books.Port + cfg.Client.Books.Endpoint
		req, err := http.NewRequest("GET", url, nil)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			log.Println("Error: ", err)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error creating request",
			})
			return
		}
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error doing the request",
			})
			return
		}

		defer res.Body.Close()
		body, readErr := io.ReadAll(res.Body)

		if readErr != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error reading response",
			})
		}

		var books []Book

		if err := json.Unmarshal(body, &books); err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Msg: "Error parsing books json",
			})
			return
		}

		var purchasedBooks []Book = []Book{}

		for _, book := range books {
			for _, title := range purchased {
				if book.Title == title {
					purchasedBooks = append(purchasedBooks, book)
				}
			}
		}

		json.NewEncoder(w).Encode(purchasedBooks)
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ok!")
	})

	server := http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: mux,
	}

	log.Printf("Server listening on %s and port %s\n", cfg.Server.Host, cfg.Server.Port)
	server.ListenAndServe()
}
