package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type TargetPokinOpdOperationalRepositoryImpl struct{}

func NewTargetPokinOpdOperationalRepositoryImpl() *TargetPokinOpdOperationalRepositoryImpl {
	return &TargetPokinOpdOperationalRepositoryImpl{}
}

func (repository *TargetPokinOpdOperationalRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperational) (domain.TargetPokinOpdOperational, error) {
	query := "INSERT INTO target_pokin_opd_operational (indikator_pokin_opd_operational_id, nilai_target, satuan) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, target.IndikatorPokinOpdOperationalId, target.NilaiTarget, target.Satuan).Scan(&target.Id)
	if err != nil {
		return domain.TargetPokinOpdOperational{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdOperationalRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperational) (domain.TargetPokinOpdOperational, error) {
	query := "UPDATE target_pokin_opd_operational SET nilai_target = $1, satuan = $2, indikator_pokin_opd_operational_id = $3, last_modified_date = NOW() WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, target.NilaiTarget, target.Satuan, target.IndikatorPokinOpdOperationalId, target.Id)
	if err != nil {
		return domain.TargetPokinOpdOperational{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdOperationalRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM target_pokin_opd_operational WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (repository *TargetPokinOpdOperationalRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdOperational, error) {
	query := "SELECT id, indikator_pokin_opd_operational_id, nilai_target, satuan FROM target_pokin_opd_operational WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var target domain.TargetPokinOpdOperational
	err := row.Scan(&target.Id, &target.IndikatorPokinOpdOperationalId, &target.NilaiTarget, &target.Satuan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TargetPokinOpdOperational{}, errors.New("id tidak ditemukan")
		}
		return domain.TargetPokinOpdOperational{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdOperationalRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdOperational, error) {
	query := "SELECT id, indikator_pokin_opd_operational_id, nilai_target, satuan FROM target_pokin_opd_operational ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TargetPokinOpdOperational{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdOperational
	for rows.Next() {
		var target domain.TargetPokinOpdOperational
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdOperationalId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdOperational{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (repository *TargetPokinOpdOperationalRepositoryImpl) FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdOperational, error) {
	query := "SELECT id, indikator_pokin_opd_operational_id, nilai_target, satuan FROM target_pokin_opd_operational WHERE indikator_pokin_opd_operational_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, indikatorId)
	if err != nil {
		return []domain.TargetPokinOpdOperational{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdOperational
	for rows.Next() {
		var target domain.TargetPokinOpdOperational
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdOperationalId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdOperational{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}
