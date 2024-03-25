package main

import (
	"encoding/json"
	"fmt"
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
}

type Book struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

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
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		books := []Book{
			Book{
				Author:      "Shelley, Mary Wollstonecraft",
				Title:       "Frankenstein; Or, The Modern Prometheus",
				Description: "Frankenstein's monster (Fictitious character) -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/84/pg84.cover.medium.jpg",
			},
			Book{
				Author:      "Austen, Jane",
				Title:       "Pride and Prejudice",
				Description: "Courtship -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/1342/pg1342.cover.medium.jpg",
			},

			Book{
				Author:      "Melville, Herman",
				Title:       "Moby Dick; Or, The Whale",
				Description: "Adventure stories",
				Cover:       "https://www.gutenberg.org/cache/epub/2701/pg2701.cover.medium.jpg",
			},

			Book{
				Author:      "Shakespeare, William",
				Title:       "Romeo and Juliet",
				Description: "Conflict of generations -- Drama",
				Cover:       "https://www.gutenberg.org/cache/epub/1513/pg1513.cover.medium.jpg",
			},

			Book{
				Author:      "Eliot, George",
				Title:       "Middlemarch",
				Description: "City and town life -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/145/pg145.cover.medium.jpg",
			},

			Book{
				Author:      "Forster, E. M. (Edward Morgan)",
				Title:       "A Room with a View",
				Description: "British -- Italy -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/2641/pg2641.cover.medium.jpg",
			},

			Book{
				Author:      "Shakespeare, William",
				Title:       "The Complete Works of William Shakespeare",
				Description: "English drama -- Early modern and Elizabethan, 1500-1600",
				Cover:       "https://www.gutenberg.org/cache/epub/100/pg100.cover.medium.jpg",
			},

			Book{
				Author:      "Alcott, Louisa May",
				Title:       "Little Women; Or, Meg, Jo, Beth, and Amy",
				Description: "Autobiographical fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/37106/pg37106.cover.medium.jpg",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	})

	mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		books := []Book{
			Book{
				Author:      "Shelley, Mary Wollstonecraft",
				Title:       "Frankenstein; Or, The Modern Prometheus",
				Description: "Frankenstein's monster (Fictitious character) -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/84/pg84.cover.medium.jpg",
			},
			Book{
				Author:      "Austen, Jane",
				Title:       "Pride and Prejudice",
				Description: "Courtship -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/1342/pg1342.cover.medium.jpg",
			},

			Book{
				Author:      "Melville, Herman",
				Title:       "Moby Dick; Or, The Whale",
				Description: "Adventure stories",
				Cover:       "https://www.gutenberg.org/cache/epub/2701/pg2701.cover.medium.jpg",
			},

			Book{
				Author:      "Shakespeare, William",
				Title:       "Romeo and Juliet",
				Description: "Conflict of generations -- Drama",
				Cover:       "https://www.gutenberg.org/cache/epub/1513/pg1513.cover.medium.jpg",
			},

			Book{
				Author:      "Eliot, George",
				Title:       "Middlemarch",
				Description: "City and town life -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/145/pg145.cover.medium.jpg",
			},

			Book{
				Author:      "Forster, E. M. (Edward Morgan)",
				Title:       "A Room with a View",
				Description: "British -- Italy -- Fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/2641/pg2641.cover.medium.jpg",
			},

			Book{
				Author:      "Shakespeare, William",
				Title:       "The Complete Works of William Shakespeare",
				Description: "English drama -- Early modern and Elizabethan, 1500-1600",
				Cover:       "https://www.gutenberg.org/cache/epub/100/pg100.cover.medium.jpg",
			},

			Book{
				Author:      "Alcott, Louisa May",
				Title:       "Little Women; Or, Meg, Jo, Beth, and Amy",
				Description: "Autobiographical fiction",
				Cover:       "https://www.gutenberg.org/cache/epub/37106/pg37106.cover.medium.jpg",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
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
