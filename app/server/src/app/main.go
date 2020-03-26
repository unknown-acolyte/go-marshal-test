package main

import (
	"app/mutations"
	"app/queries"
	"github.com/graphql-go/handler"
	"os"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	//"log"
	"net/http"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.RootQuery,
	Mutation: mutations.RootMutation,
})

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", disableCors(h))

	log.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err.Error)
	}
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}