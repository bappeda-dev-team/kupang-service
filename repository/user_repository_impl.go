package repository

import (
	"context"
	"database/sql"
	"errors"

	"kupang-service/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := `INSERT INTO "user" (nama, nip, email, status, role, kode_opd, opd_id, nama_opd, pegawai_id, nama_pegawai, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	err := tx.QueryRowContext(ctx, query, user.Nama, user.Nip, user.Email, user.Status, user.Role, user.KodeOpd, user.OpdId, user.NamaOpd, user.PegawaiId, user.NamaPegawai, user.RoleId).Scan(&user.Id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := `UPDATE "user" SET nama = $1, nip = $2, email = $3, status = $4, role = $5, kode_opd = $6, opd_id = $7, nama_opd = $8, pegawai_id = $9, nama_pegawai = $10, role_id = $11, last_modified_date = NOW() WHERE id = $12`
	_, err := tx.ExecContext(ctx, query, user.Nama, user.Nip, user.Email, user.Status, user.Role, user.KodeOpd, user.OpdId, user.NamaOpd, user.PegawaiId, user.NamaPegawai, user.RoleId, user.Id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM "user" WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	query := `SELECT id, nama, nip, email, status, role, kode_opd, opd_id, nama_opd, pegawai_id, nama_pegawai, role_id FROM "user" WHERE id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	var user domain.User
	err := row.Scan(&user.Id, &user.Nama, &user.Nip, &user.Email, &user.Status, &user.Role, &user.KodeOpd, &user.OpdId, &user.NamaOpd, &user.PegawaiId, &user.NamaPegawai, &user.RoleId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("id tidak ditemukan")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error) {
	query := `SELECT id, nama, nip, email, status, role, kode_opd, opd_id, nama_opd, pegawai_id, nama_pegawai, role_id FROM "user" ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.User{}, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Nama, &user.Nip, &user.Email, &user.Status, &user.Role, &user.KodeOpd, &user.OpdId, &user.NamaOpd, &user.PegawaiId, &user.NamaPegawai, &user.RoleId)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
