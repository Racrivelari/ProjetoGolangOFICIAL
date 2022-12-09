package main

import (
	"encoding/json"
	"net/http"
	"os"
	"log"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/config"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/handler"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/pkg/database"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/pkg/service"
	
	// lhttp "github.com/faelp22/tcs_curso/stoq/pkg/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	default_conf := &config.DBConfig{}    //aq nao ta puxando
	println("Driver Database: ",default_conf.DB_DRIVE)

	if file_config := os.Getenv("STOQ_CONFIG"); file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	dbpool := database.NewDB(default_conf)

	service := service.NewProdutoService(dbpool)

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	r.HandleFunc("/", redirect)

	handler.RegisterAPIHandlers(r, n, service)

	log.Fatal(http.ListenAndServe(":5000", r))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/webui/", http.StatusMovedPermanently)
}

