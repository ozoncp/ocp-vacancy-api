package api

import (
	"context"

	"github.com/ozoncp/ocp-vacancy-api/internal/models"
	"github.com/ozoncp/ocp-vacancy-api/internal/repo"
	ocp_vacancy_api "github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	ocp_vacancy_api.UnimplementedOcpVacancyApiServer
	repo repo.Repo
}

func (a *api) CreateVacancyV1(ctx context.Context, req *ocp_vacancy_api.CreateVacancyV1Request) (*ocp_vacancy_api.CreateVacancyV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancy := models.Vacancy{
		Link:   req.Link,
		Status: req.Status,
	}

	vacancyid, err := a.repo.CreateVacancy(ctx, vacancy)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vacancy")
	}

	log.Info().
		Str("Link", req.Link).
		Uint64("Status", req.Status).
		Msg("create vacancy")

	return &ocp_vacancy_api.CreateVacancyV1Response{Id: vacancyid}, nil
}

func (a *api) UpdateVacancyV1(ctx context.Context, req *ocp_vacancy_api.UpdateVacancyV1Request) (*ocp_vacancy_api.UpdateVacancyV1Response, error) {
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
		return nil, errors.Wrap(err, "failed to update vacancy")
	}

	log.Info().
		Uint64("ID", req.Id).
		Str("Link", req.Link).
		Uint64("Status", req.Status).
		Msg("update vacancy by id")

	return &ocp_vacancy_api.UpdateVacancyV1Response{
		Vacancy: &ocp_vacancy_api.VacancyV1{
			Id:     vacancy.ID,
			Link:   vacancy.Link,
			Status: vacancy.Status,
		},
	}, nil
}

func (a *api) DescribeVacancyV1(ctx context.Context, req *ocp_vacancy_api.DescribeVacancyV1Request) (*ocp_vacancy_api.DescribeVacancyV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancy, err := a.repo.DescribeVacancy(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get vacancy by ID")
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
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vacancies, err := a.repo.ListVacancies(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list vacancies")
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
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.repo.RemoveVacancy(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove vacancy")
	}

	log.Info().
		Uint64("ID", req.Id).
		Msg("remove vacancy by id")

	return &ocp_vacancy_api.RemoveVacancyV1Response{Found: true}, nil
}

func NewOcpVacancyApi(repo repo.Repo) ocp_vacancy_api.OcpVacancyApiServer {
	return &api{repo: repo}
}
