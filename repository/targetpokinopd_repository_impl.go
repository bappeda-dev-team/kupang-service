package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type TargetPokinOpdRepositoryImpl struct{}

func NewTargetPokinOpdRepositoryImpl() *TargetPokinOpdRepositoryImpl {
	return &TargetPokinOpdRepositoryImpl{}
}

func (repository *TargetPokinOpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, targetPokinOpd domain.TargetPokinOpd) (domain.TargetPokinOpd, error) {
	query := "INSERT INTO target_pokin_opd (indikator_pokin_opd_id, nilai_target, satuan) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, targetPokinOpd.IndikatorPokinOpdId, targetPokinOpd.NilaiTarget, targetPokinOpd.Satuan).Scan(&targetPokinOpd.Id)
	if err != nil {
		return domain.TargetPokinOpd{}, err
	}

	return targetPokinOpd, nil
}

func (repository *TargetPokinOpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, targetPokinOpd domain.TargetPokinOpd) (domain.TargetPokinOpd, error) {
	query := "UPDATE target_pokin_opd SET nilai_target = $1, satuan = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, targetPokinOpd.NilaiTarget, targetPokinOpd.Satuan, targetPokinOpd.Id)
	if err != nil {
		return domain.TargetPokinOpd{}, err
	}

	return targetPokinOpd, nil
}

func (repository *TargetPokinOpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM target_pokin_opd WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *TargetPokinOpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpd, error) {
	query := "SELECT id, indikator_pokin_opd_id, nilai_target, satuan FROM target_pokin_opd WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var targetPokinOpd domain.TargetPokinOpd
	err := row.Scan(&targetPokinOpd.Id, &targetPokinOpd.IndikatorPokinOpdId, &targetPokinOpd.NilaiTarget, &targetPokinOpd.Satuan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TargetPokinOpd{}, errors.New("id tidak ditemukan")
		}
		return domain.TargetPokinOpd{}, err
	}

	return targetPokinOpd, nil
}

func (repository *TargetPokinOpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpd, error) {
	query := "SELECT id, indikator_pokin_opd_id, nilai_target, satuan FROM target_pokin_opd ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TargetPokinOpd{}, err
	}
	defer rows.Close()

	var targetPokinOpdList []domain.TargetPokinOpd
	for rows.Next() {
		var targetPokinOpd domain.TargetPokinOpd
		err := rows.Scan(&targetPokinOpd.Id, &targetPokinOpd.IndikatorPokinOpdId, &targetPokinOpd.NilaiTarget, &targetPokinOpd.Satuan)
		if err != nil {
			return []domain.TargetPokinOpd{}, err
		}

		targetPokinOpdList = append(targetPokinOpdList, targetPokinOpd)
	}

	return targetPokinOpdList, nil
}

func (repository *TargetPokinOpdRepositoryImpl) FindByIndikatorPokinOpdId(ctx context.Context, tx *sql.Tx, indikatorPokinOpdId int) ([]domain.TargetPokinOpd, error) {
	query := "SELECT id, indikator_pokin_opd_id, nilai_target, satuan FROM target_pokin_opd WHERE indikator_pokin_opd_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, indikatorPokinOpdId)
	if err != nil {
		return []domain.TargetPokinOpd{}, err
	}
	defer rows.Close()

	var targetPokinOpdList []domain.TargetPokinOpd
	for rows.Next() {
		var targetPokinOpd domain.TargetPokinOpd
		err := rows.Scan(&targetPokinOpd.Id, &targetPokinOpd.IndikatorPokinOpdId, &targetPokinOpd.NilaiTarget, &targetPokinOpd.Satuan)
		if err != nil {
			return []domain.TargetPokinOpd{}, err
		}

		targetPokinOpdList = append(targetPokinOpdList, targetPokinOpd)
	}

	return targetPokinOpdList, nil
}
