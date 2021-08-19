package main

import (
	"context"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-vacancy-api/internal/api"
	ocp_vacancy_api "github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	grpcPort           = ":8081"
	grpcServerEndpoint = ":8081"
	httpPort           = ":8083"
)

func main() {
	go runJSON()

	if err := run(); err != nil {
		log.Fatal().
			Msgf(err.Error())
	}
}

func run() error {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return errors.Wrapf(err, "failed to start tcp listener on %s", grpcPort)
	}

	server := grpc.NewServer()
	ocp_vacancy_api.RegisterOcpVacancyApiServer(server, api.NewOcpVacancyApi())

	log.Info().
		Str("address", listener.Addr().String()).
		Msg("serving gRPC...")

	if err := server.Serve(listener); err != nil {
		return errors.Wrap(err, "failed to serve gRPC")
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := ocp_vacancy_api.RegisterOcpVacancyApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Panic().
			Msgf("failed to register API handler from endpoint '%s': %v", grpcServerEndpoint, err)
	}

	log.Info().
		Str("address", httpPort).
		Msg("serving REST...")

	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Panic().
			Msgf("failed to serve JSON endpoint: %v", err)
	}
}
