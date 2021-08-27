package api

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-vacancy-api/internal/config"
	"github.com/ozoncp/ocp-vacancy-api/internal/metrics"
	"github.com/ozoncp/ocp-vacancy-api/internal/models"
	"github.com/ozoncp/ocp-vacancy-api/internal/producer"
	"github.com/ozoncp/ocp-vacancy-api/internal/repo"
	"github.com/ozoncp/ocp-vacancy-api/internal/utils"
	ocp_vacancy_api "github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	ocp_vacancy_api.UnimplementedOcpVacancyApiServer
	repo  repo.Repo
	kafka producer.Producer
}

func (a *api) CreateVacancyV1(ctx context.Context, req *ocp_vacancy_api.CreateVacancyV1Request) (*ocp_vacancy_api.CreateVacancyV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreateVacancyV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancy := models.Vacancy{
		Link:   req.Link,
		Status: req.Status,
	}

	vacancyid, err := a.repo.CreateVacancy(ctx, vacancy)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to create vacancy").Error())
	}

	log.Info().
		Str("Link", req.Link).
		Uint64("Status", req.Status).
		Msg("create vacancy")

	err = a.kafka.Send(producer.Message{
		Type: producer.Create,
		Body: map[string]interface{}{
			"id":        vacancyid,
			"timestamp": time.Now().Unix(),
		},
	})
	if err != nil {
		log.Error().
			Err(err).
			Msgf("vacancy was created, but failed to send kafka message")
		return nil, status.Error(codes.Internal, errors.Wrap(err, "vacancy was created, but failed to send kafka message").Error())
	}

	metrics.IncCreateCounter()

	return &ocp_vacancy_api.CreateVacancyV1Response{Id: vacancyid}, nil
}

func (a *api) UpdateVacancyV1(ctx context.Context, req *ocp_vacancy_api.UpdateVacancyV1Request) (*ocp_vacancy_api.UpdateVacancyV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdateVacancyV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancy := models.Vacancy{
		ID:     req.Id,
		Link:   req.Link,
		Status: req.Status,
	}

	vacancy, err := a.repo.UpdateVacancy(ctx, vacancy)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to update vacancy").Error())
	}

	log.Info().
		Uint64("ID", req.Id).
		Str("Link", req.Link).
		Uint64("Status", req.Status).
		Msg("update vacancy by id")

	err = a.kafka.Send(producer.Message{
		Type: producer.Update,
		Body: map[string]interface{}{
			"id":        vacancy.ID,
			"timestamp": time.Now().Unix(),
		},
	})
	if err != nil {
		log.Error().
			Err(err).
			Msgf("vacancy was updated, but failed to send kafka message")
		return nil, status.Error(codes.Internal, errors.Wrap(err, "vacancy was updated, but failed to send kafka message").Error())
	}

	metrics.IncUpdateCounter()

	return &ocp_vacancy_api.UpdateVacancyV1Response{
		Vacancy: &ocp_vacancy_api.VacancyV1{
			Id:     vacancy.ID,
			Link:   vacancy.Link,
			Status: vacancy.Status,
		},
	}, nil
}

func (a *api) DescribeVacancyV1(ctx context.Context, req *ocp_vacancy_api.DescribeVacancyV1Request) (*ocp_vacancy_api.DescribeVacancyV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("DescribeVacancyV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancy, err := a.repo.DescribeVacancy(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get vacancy by ID").Error())
	}

	log.Info().
		Uint64("ID", req.Id).
		Msg("get vacancy by id")

	return &ocp_vacancy_api.DescribeVacancyV1Response{
		Vacancy: &ocp_vacancy_api.VacancyV1{
			Id:     vacancy.ID,
			Link:   vacancy.Link,
			Status: vacancy.Status,
		},
	}, nil
}

func (a *api) ListVacanciesV1(ctx context.Context, req *ocp_vacancy_api.ListVacanciesV1Request) (*ocp_vacancy_api.ListVacanciesV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ListVacanciesV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancies, err := a.repo.ListVacancies(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to list vacancies").Error())
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("list vacancies")

	result := &ocp_vacancy_api.ListVacanciesV1Response{}
	result.Vacancies = make([]*ocp_vacancy_api.VacancyV1, len(vacancies))

	for i := range vacancies {
		result.Vacancies[i] = &ocp_vacancy_api.VacancyV1{
			Id:     vacancies[i].ID,
			Link:   vacancies[i].Link,
			Status: vacancies[i].Status,
		}
	}

	return result, nil
}

func (a *api) RemoveVacancyV1(ctx context.Context, req *ocp_vacancy_api.RemoveVacancyV1Request) (*ocp_vacancy_api.RemoveVacancyV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("RemoveVacancyV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.repo.RemoveVacancy(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to remove vacancy").Error())
	}

	log.Info().
		Uint64("ID", req.Id).
		Msg("remove vacancy by id")

	err = a.kafka.Send(producer.Message{
		Type: producer.Remove,
		Body: map[string]interface{}{
			"id":        req.Id,
			"timestamp": time.Now().Unix(),
		},
	})
	if err != nil {
		log.Error().
			Err(err).
			Msgf("vacancy was removed, but failed to send kafka message")
		return nil, status.Error(codes.Internal, errors.Wrap(err, "vacancy was removed, but failed to send kafka message").Error())
	}

	metrics.IncRemoveCounter()

	return &ocp_vacancy_api.RemoveVacancyV1Response{Found: true}, nil
}

func (a *api) MultiCreateVacanciesV1(ctx context.Context, req *ocp_vacancy_api.MultiCreateVacanciesV1Request) (*ocp_vacancy_api.MultiCreateVacanciesV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateVacanciesV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancies := make([]models.Vacancy, 0, len(req.Vacancies))
	for _, r := range req.Vacancies {
		vacancies = append(vacancies, models.Vacancy{
			Status: r.Status,
			Link:   r.Link,
		})
	}

	batchSize := config.GetConfig().BatchSize
	vacanciesByBatches := utils.SplitSliceVacancy(vacancies, batchSize)
	for i := 0; i < len(vacanciesByBatches); i++ {
		err := func() error {
			childSpan := tracer.StartSpan("MultiCreateVacanciesV1", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			return a.repo.AddVacancies(ctx, vacanciesByBatches[i])
		}()

		if err != nil {
			return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to add vacancies batch").Error())
		}
	}

	log.Info().
		Msg("multiple vacancies added")

	return &ocp_vacancy_api.MultiCreateVacanciesV1Response{Added: true}, nil
}

func NewOcpVacancyApi(repo repo.Repo, kafka producer.Producer) ocp_vacancy_api.OcpVacancyApiServer {
	return &api{repo: repo, kafka: kafka}
}
