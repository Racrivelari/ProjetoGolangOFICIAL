package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/config"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/handler"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/pkg/database"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/pkg/service"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	default_conf := &config.Config{}

	if file_config := os.Getenv("STOQ_CONFIG"); file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	conf := config.NewConfig(default_conf)

	dbpool := database.NewDB(conf)
	service := service.NewProdutoService(dbpool)

	println(conf.DB_DRIVE)
	println(conf.DB_DSN)

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	r.HandleFunc("/", redirect)
	handler.RegisterAPIHandlers(r, n, service)
	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/webui/", http.StatusMovedPermanently)
}

