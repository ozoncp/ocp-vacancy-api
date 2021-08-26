package repo

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-vacancy-api/internal/models"
)

// Repo is a repository interface for Vacancy entity
type Repo interface {
	AddVacancies(ctx context.Context, vacancies []models.Vacancy) error
	CreateVacancy(ctx context.Context, vacancy models.Vacancy) (uint64, error)
	UpdateVacancy(ctx context.Context, vacancy models.Vacancy) (models.Vacancy, error)
	ListVacancies(ctx context.Context, limit, offset uint64) ([]models.Vacancy, error)
	DescribeVacancy(ctx context.Context, vacancyID uint64) (models.Vacancy, error)
	RemoveVacancy(ctx context.Context, vacancyID uint64) error
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

const vacanciesTable string = "vacancies"

func (r *repo) AddVacancies(ctx context.Context, vacancies []models.Vacancy) error {
	return errors.New("not implemented")
}

func (r *repo) CreateVacancy(ctx context.Context, vacancy models.Vacancy) (uint64, error) {
	qry := squirrel.Insert(vacanciesTable).
		SetMap(map[string]interface{}{
			"link":   vacancy.Link,
			"status": vacancy.Status,
		}).
		Suffix("RETURNING id")

	err := qry.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db).
		QueryRowContext(ctx).
		Scan(&vacancy.ID)

	return vacancy.ID, err
}

func (r *repo) UpdateVacancy(ctx context.Context, vacancy models.Vacancy) (models.Vacancy, error) {
	res, err := squirrel.Update(vacanciesTable).
		SetMap(map[string]interface{}{
			"link":   vacancy.Link,
			"status": vacancy.Status,
		}).
		Where("id = ?", vacancy.ID).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar).
		ExecContext(ctx)
	if err != nil {
		return models.Vacancy{}, err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return models.Vacancy{}, errors.Wrap(err, "failed to get affected rows")
	}

	if ra <= 0 {
		return models.Vacancy{}, errors.New("no rows were affected: probably, there is no row with given ID")
	}

	return vacancy, nil
}

func (r *repo) ListVacancies(ctx context.Context, limit, offset uint64) ([]models.Vacancy, error) {
	qry, _, err := squirrel.Select("id", "link", "status").
		From(vacanciesTable).
		Limit(limit).
		Offset(offset).
		OrderBy("id").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var result []models.Vacancy

	err = r.db.SelectContext(ctx, &result, qry)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list vacanices")
	}
	return result, nil
}

func (r *repo) DescribeVacancy(ctx context.Context, vacancyID uint64) (models.Vacancy, error) {
	qry, args, err := squirrel.Select("id", "link", "status").
		From(vacanciesTable).
		Where("id = ?").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Vacancy{}, errors.Wrap(err, "failed to build query")
	}

	var result models.Vacancy

	err = r.db.GetContext(ctx, &result, qry, args)
	if err != nil {
		return models.Vacancy{}, errors.Wrap(err, "failed to get vacancy by id")
	}
	return result, nil
}

func (r *repo) RemoveVacancy(ctx context.Context, vacancyID uint64) error {
	res, err := squirrel.Delete(vacanciesTable).
		Where("id = ?", vacancyID).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get affected rows")
	}

	if ra <= 0 {
		return errors.New("no rows were affected: probably, there is no row with given ID")
	}

	return nil
}
