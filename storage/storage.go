package storage

import (
	log "real_project/pkg/logger"
	"real_project/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageI interface {
	GetContentRepo() postgres.ContentRepoI
	GetOwnerRepo() postgres.OwnerRepoI
	GetCommonRepo() postgres.CommonRepoI
}

type storage struct {
	contentRepo postgres.ContentRepoI
	ownerRepo   postgres.OwnerRepoI
	commonRepo  postgres.CommonRepoI
}

func NewStorage(db *pgxpool.Pool, log log.Log) StorageI {
	return &storage{
		contentRepo: postgres.NewContentRepo(db, log),
		ownerRepo:   postgres.NewOwnerRepo(db, log),
	}
}

func (s *storage) GetContentRepo() postgres.ContentRepoI {
	return s.contentRepo
}
func (s *storage) GetOwnerRepo() postgres.OwnerRepoI {
	return s.ownerRepo
}

func (s *storage) GetCommonRepo() postgres.CommonRepoI {
	return s.commonRepo
}
