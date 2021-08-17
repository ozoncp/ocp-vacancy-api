package api

import (
	"context"

	ocp_vacancy_api "github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	ocp_vacancy_api.UnimplementedOcpVacancyApiServer
}

func (a *api) CreateVacancyV1(ctx context.Context, req *ocp_vacancy_api.CreateVacancyV1Request) (*ocp_vacancy_api.CreateVacancyV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Str("Link", req.Link).
		Uint64("Status", req.Status).
		Msg("create vacancy")

	return &ocp_vacancy_api.CreateVacancyV1Response{ID: 1}, nil
}

func (a *api) DescribeVacancyV1(ctx context.Context, req *ocp_vacancy_api.DescribeVacancyV1Request) (*ocp_vacancy_api.DescribeVacancyV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("ID", req.ID).
		Msg("get vacancy by id")

	return &ocp_vacancy_api.DescribeVacancyV1Response{
		Vacancy: &ocp_vacancy_api.VacancyV1{
			ID:     req.ID,
			Link:   "https://example.com",
			Status: 0,
		},
	}, nil
}

func (a *api) ListVacanciesV1(ctx context.Context, req *ocp_vacancy_api.ListVacanciesV1Request) (*ocp_vacancy_api.ListVacanciesV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("list vacancies")

	return &ocp_vacancy_api.ListVacanciesV1Response{
		Vacancies: []*ocp_vacancy_api.VacancyV1{
			{
				ID:     1,
				Link:   "https://example.com",
				Status: 0,
			},
		},
	}, nil
}

func (a *api) RemoveVacancyV1(ctx context.Context, req *ocp_vacancy_api.RemoveVacancyV1Request) (*ocp_vacancy_api.RemoveVacancyV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("ID", req.ID).
		Msg("remove vacancy by id")

	return &ocp_vacancy_api.RemoveVacancyV1Response{Found: false}, nil
}

func NewOcpVacancyApi() ocp_vacancy_api.OcpVacancyApiServer {
	return &api{}
}
