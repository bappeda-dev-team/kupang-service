package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type TargetPokinOpdOperationalNRepositoryImpl struct{}

func NewTargetPokinOpdOperationalNRepositoryImpl() *TargetPokinOpdOperationalNRepositoryImpl {
	return &TargetPokinOpdOperationalNRepositoryImpl{}
}

func (repository *TargetPokinOpdOperationalNRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperationalN) (domain.TargetPokinOpdOperationalN, error) {
	query := "INSERT INTO target_pokin_opd_operationalN (indikator_pokin_opd_operationalN_id, nilai_target, satuan) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, target.IndikatorPokinOpdOperationalNId, target.NilaiTarget, target.Satuan).Scan(&target.Id)
	if err != nil {
		return domain.TargetPokinOpdOperationalN{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdOperationalNRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperationalN) (domain.TargetPokinOpdOperationalN, error) {
	query := "UPDATE target_pokin_opd_operationalN SET nilai_target = $1, satuan = $2, indikator_pokin_opd_operationalN_id = $3, last_modified_date = NOW() WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, target.NilaiTarget, target.Satuan, target.IndikatorPokinOpdOperationalNId, target.Id)
	if err != nil {
		return domain.TargetPokinOpdOperationalN{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdOperationalNRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM target_pokin_opd_operationalN WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (repository *TargetPokinOpdOperationalNRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdOperationalN, error) {
	query := "SELECT id, indikator_pokin_opd_operationalN_id, nilai_target, satuan FROM target_pokin_opd_operationalN WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var target domain.TargetPokinOpdOperationalN
	err := row.Scan(&target.Id, &target.IndikatorPokinOpdOperationalNId, &target.NilaiTarget, &target.Satuan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TargetPokinOpdOperationalN{}, errors.New("id tidak ditemukan")
		}
		return domain.TargetPokinOpdOperationalN{}, err
	}

	return target, nil
}

func (repository *TargetPokinOpdOperationalNRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdOperationalN, error) {
	query := "SELECT id, indikator_pokin_opd_operationalN_id, nilai_target, satuan FROM target_pokin_opd_operationalN ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TargetPokinOpdOperationalN{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdOperationalN
	for rows.Next() {
		var target domain.TargetPokinOpdOperationalN
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdOperationalNId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdOperationalN{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (repository *TargetPokinOpdOperationalNRepositoryImpl) FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdOperationalN, error) {
	query := "SELECT id, indikator_pokin_opd_operationalN_id, nilai_target, satuan FROM target_pokin_opd_operationalN WHERE indikator_pokin_opd_operationalN_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, indikatorId)
	if err != nil {
		return []domain.TargetPokinOpdOperationalN{}, err
	}
	defer rows.Close()

	var targets []domain.TargetPokinOpdOperationalN
	for rows.Next() {
		var target domain.TargetPokinOpdOperationalN
		err := rows.Scan(&target.Id, &target.IndikatorPokinOpdOperationalNId, &target.NilaiTarget, &target.Satuan)
		if err != nil {
			return []domain.TargetPokinOpdOperationalN{}, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}
