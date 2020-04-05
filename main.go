package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go-hls-stream-player/router"
	"log"
	"net/http"
)

func init()  {
	
}

func main()  {
	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()
	playlistRouter := &router.PlaylistRouter{Router: api}
	playlistRouter.RegisterHandlers()

	r.NotFoundHandler = http.HandlerFunc(NotFound)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"X-Requested-With",
		},
	})

	log.Fatal(http.ListenAndServe(":" + "8006", corsOpts.Handler(r)))
}


func NotFound(w http.ResponseWriter, r *http.Request) {
	rsp := "route not found: " + r.URL.Path
	w.Write([]byte(rsp))
}