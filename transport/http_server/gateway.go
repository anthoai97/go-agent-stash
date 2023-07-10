package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"anquach.dev/go-agent-stash/serializer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/grpclog"
)

// Proxy to gRPC Server

func Run(dialAddr string) error {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	// conn, err := grpc.DialContext(
	// 	context.Background(),
	// 	dialAddr,
	// 	nil,
	// 	grpc.WithBlock(),
	// )

	// if err != nil {
	// 	return fmt.Errorf("failed to dial server: %w", err)
	// }

	gwmux := runtime.NewServeMux()
	// TODO: Register Service gRPC for gatewat
	port := serializer.GetEnvVar("PORT", "8080")
	gatewayAddr := "0.0.0.0:" + port
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RequestURI)
		w.Header().Set("Content-Type", "application/json")
		// time.Sleep(1 * time.Minute)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"alive": true}`)
	}))
	// gwServer := &http.Server{
	// 	Addr: gatewayAddr,
	// 	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		if strings.HasPrefix(r.URL.Path, "/") {
	// 			gwmux.ServeHTTP(w, r)
	// 			return
	// 		}
	// 		if r.Method == "GET" && r.URL.Path == "/health" {
	// 			log.Info(r.RequestURI)
	// 			w.Header().Set("Content-Type", "application/json")
	// 			time.Sleep(1 * time.Minute)
	// 			w.WriteHeader(http.StatusOK)

	// 			io.WriteString(w, `{"alive": true}`)
	// 			return
	// 		}
	// 		// oa.ServeHTTP(w, r)
	// 	}),
	// }

	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", http.ListenAndServe(gatewayAddr, mux))
}
