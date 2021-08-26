package main

import (
	"context"
	"io"
	"net"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-vacancy-api/internal/api"
	"github.com/ozoncp/ocp-vacancy-api/internal/config"
	"github.com/ozoncp/ocp-vacancy-api/internal/metrics"
	"github.com/ozoncp/ocp-vacancy-api/internal/producer"
	"github.com/ozoncp/ocp-vacancy-api/internal/repo"
	ocp_vacancy_api "github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {

	cfg := config.GetConfig()

	closer, err := initJaeger(cfg.Jaeger.Address)
	if err != nil {
		log.Fatal().
			Err(err)
	}
	defer closer.Close()

	go runMetrics(cfg.Prometheus.Address)

	go runJSON(cfg.GRPC.Address, cfg.REST.Address)

	if err := run(cfg.Database.ConnectionString, cfg.GRPC.Address, cfg.Kafka.Topic, cfg.Kafka.Brokers, cfg.BatchSize); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to serve GRPC...")
	}
}

func run(dbConnectionString, grpcPort, kafkaTopic string, kafkaBrokers []string, batchSize int) error {
	db, err := sqlx.Connect("pgx", dbConnectionString)
	if err != nil {
		return errors.Wrap(err, "failed to connect to DB")
	}

	kafka, err := producer.NewProducer(kafkaBrokers, kafkaTopic)
	if err != nil {
		return errors.Wrap(err, "failed to initialize Kafka producer")
	}
	server := grpc.NewServer()
	ocp_vacancy_api.RegisterOcpVacancyApiServer(server, api.NewOcpVacancyApi(repo.NewRepo(db, batchSize), kafka))

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return errors.Wrapf(err, "failed to start tcp listener on %s", grpcPort)
	}

	log.Info().
		Str("address", listener.Addr().String()).
		Msg("serving gRPC...")

	if err := server.Serve(listener); err != nil {
		return errors.Wrap(err, "failed to serve gRPC")
	}

	return nil
}

func runJSON(grpcServerEndpoint, httpPort string) {
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

func runMetrics(address string) {
	metrics.RegisterMetrics()

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	log.Info().
		Str("address", address).
		Msg("serving metrics...")

	if err := http.ListenAndServe(address, mux); err != nil {
		log.Panic().
			Msgf("failed to serve metrics endpoint: %v", err)
	}
}

func initJaeger(address string) (io.Closer, error) {
	cfgMetrics := &jaegerconfig.Configuration{
		ServiceName: "ocp-vacancy-api",
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans: true,
			//LocalAgentHostPort: address,
		},
	}
	tracer, closer, err := cfgMetrics.NewTracer(
		jaegerconfig.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Jaeger")
	}
	opentracing.SetGlobalTracer(tracer)

	log.Info().
		Str("address", address).
		Msgf("started tracing with Jaeger...")

	return closer, nil
}
