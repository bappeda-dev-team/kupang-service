package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type TargetPokinOpdStrategicRepositoryImpl struct{}

func NewTargetPokinOpdStrategicRepositoryImpl() *TargetPokinOpdStrategicRepositoryImpl {
	return &TargetPokinOpdStrategicRepositoryImpl{}
}

func (repository *TargetPokinOpdStrategicRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdStrategic) (domain.TargetPokinOpdStrategic, error) {
	query := "INSERT INTO target_pokin_opd_strategic (indikator_pokin_opd_strategic_id, nilai_target, satuan) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, target.IndikatorPokinOpdStrategicId, target.NilaiTarget, target.Satuan).Scan(&target.Id)
	if err != nil {
		return domain.TargetPokinOpdStrategic{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdStrategicRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdStrategic) (domain.TargetPokinOpdStrategic, error) {
	query := "UPDATE target_pokin_opd_strategic SET nilai_target = $1, satuan = $2, indikator_pokin_opd_strategic_id = $3, last_modified_date = NOW() WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, target.NilaiTarget, target.Satuan, target.IndikatorPokinOpdStrategicId, target.Id)
	if err != nil {
		return domain.TargetPokinOpdStrategic{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdStrategicRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM target_pokin_opd_strategic WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (repository *TargetPokinOpdStrategicRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdStrategic, error) {
	query := "SELECT id, indikator_pokin_opd_strategic_id, nilai_target, satuan FROM target_pokin_opd_strategic WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var target domain.TargetPokinOpdStrategic
	err := row.Scan(&target.Id, &target.IndikatorPokinOpdStrategicId, &target.NilaiTarget, &target.Satuan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TargetPokinOpdStrategic{}, errors.New("id tidak ditemukan")
		}
		return domain.TargetPokinOpdStrategic{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdStrategicRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdStrategic, error) {
	query := "SELECT id, indikator_pokin_opd_strategic_id, nilai_target, satuan FROM target_pokin_opd_strategic ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TargetPokinOpdStrategic{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdStrategic
	for rows.Next() {
		var target domain.TargetPokinOpdStrategic
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdStrategicId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdStrategic{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (repository *TargetPokinOpdStrategicRepositoryImpl) FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdStrategic, error) {
	query := "SELECT id, indikator_pokin_opd_strategic_id, nilai_target, satuan FROM target_pokin_opd_strategic WHERE indikator_pokin_opd_strategic_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, indikatorId)
	if err != nil {
		return []domain.TargetPokinOpdStrategic{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdStrategic
	for rows.Next() {
		var target domain.TargetPokinOpdStrategic
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdStrategicId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdStrategic{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}
