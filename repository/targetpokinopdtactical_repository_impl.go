package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type TargetPokinOpdTacticalRepositoryImpl struct{}

func NewTargetPokinOpdTacticalRepositoryImpl() *TargetPokinOpdTacticalRepositoryImpl {
	return &TargetPokinOpdTacticalRepositoryImpl{}
}

func (repository *TargetPokinOpdTacticalRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdTactical) (domain.TargetPokinOpdTactical, error) {
	query := "INSERT INTO target_pokin_opd_tactical (indikator_pokin_opd_tactical_id, nilai_target, satuan) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, target.IndikatorPokinOpdTacticalId, target.NilaiTarget, target.Satuan).Scan(&target.Id)
	if err != nil {
		return domain.TargetPokinOpdTactical{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdTacticalRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdTactical) (domain.TargetPokinOpdTactical, error) {
	query := "UPDATE target_pokin_opd_tactical SET nilai_target = $1, satuan = $2, indikator_pokin_opd_tactical_id = $3, last_modified_date = NOW() WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, target.NilaiTarget, target.Satuan, target.IndikatorPokinOpdTacticalId, target.Id)
	if err != nil {
		return domain.TargetPokinOpdTactical{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdTacticalRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM target_pokin_opd_tactical WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (repository *TargetPokinOpdTacticalRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdTactical, error) {
	query := "SELECT id, indikator_pokin_opd_tactical_id, nilai_target, satuan FROM target_pokin_opd_tactical WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var target domain.TargetPokinOpdTactical
	err := row.Scan(&target.Id, &target.IndikatorPokinOpdTacticalId, &target.NilaiTarget, &target.Satuan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TargetPokinOpdTactical{}, errors.New("id tidak ditemukan")
		}
		return domain.TargetPokinOpdTactical{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdTacticalRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdTactical, error) {
	query := "SELECT id, indikator_pokin_opd_tactical_id, nilai_target, satuan FROM target_pokin_opd_tactical ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TargetPokinOpdTactical{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdTactical
	for rows.Next() {
		var target domain.TargetPokinOpdTactical
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdTacticalId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdTactical{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (repository *TargetPokinOpdTacticalRepositoryImpl) FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdTactical, error) {
	query := "SELECT id, indikator_pokin_opd_tactical_id, nilai_target, satuan FROM target_pokin_opd_tactical WHERE indikator_pokin_opd_tactical_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, indikatorId)
	if err != nil {
		return []domain.TargetPokinOpdTactical{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdTactical
	for rows.Next() {
		var target domain.TargetPokinOpdTactical
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdTacticalId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdTactical{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}
