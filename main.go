package main

import (
	"fmt"
	"github.com/heptiolabs/healthcheck"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go-hls-stream-player/Models"
	"go-hls-stream-player/router"
	"time"

	// "go-hls-stream-player/router"
	"log"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	Models.InitEnv()
}

func main()  {

	health := healthcheck.NewHandler()

	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	injector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
	extractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)

	zipkinSharedRPCSpan := jaeger.TracerOptions.ZipkinSharedRPCSpan(true)

	sender, err := jaeger.NewUDPTransport(Models.GetEnvStruct().TracingConnection, 0)


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

	go http.ListenAndServe("0.0.0.0:8888", health)

	if err == nil {
		fmt.Println("success: TRACING")
		tracer, closer := jaeger.NewTracer(
			"chunk-downloader",
			jaeger.NewConstSampler(true),
			jaeger.NewRemoteReporter(
				sender,
				jaeger.ReporterOptions.BufferFlushInterval(1*time.Second)),
			injector,
			extractor,
			zipkinSharedRPCSpan,
		)
		defer closer.Close()
		log.Fatal(http.ListenAndServe(":" + Models.GetEnvStruct().Port, nethttp.Middleware(tracer, corsOpts.Handler(r)))  )
	} else {
		fmt.Println( "err: ", err)
		log.Fatal(http.ListenAndServe(":" + Models.GetEnvStruct().Port, corsOpts.Handler(r)) )
	}
}


func NotFound(w http.ResponseWriter, r *http.Request) {
	rsp := "route not found: " + r.URL.Path
	w.Write([]byte(rsp))
}