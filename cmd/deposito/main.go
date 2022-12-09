package main

import (
	"encoding/json"
	"net/http"
	"os"
	"log"
	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/handler"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
	// lhttp "github.com/faelp22/tcs_curso/stoq/pkg/http"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	default_conf := &config.DBConfig{}    //aq nao ta puxando
	println("A",default_conf.DB_DRIVE)

	if file_config := os.Getenv("STOQ_CONFIG"); file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	// conf := config.NewConfig(default_conf)
	// println("b",default_conf.DB_DRIVE)
	dbpool := database.NewDB(default_conf)

	// dbpool := database.NewDB(conf)
	println("c",default_conf.DB_DRIVE)  //n chega aq

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

