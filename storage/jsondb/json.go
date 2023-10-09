package jsondb

import (
	"os"
	"simple-clinic/config"
	"simple-clinic/storage"
)

type Store struct {
	clinic storage.ClinicRepoI
	branch storage.BranchRepoI
}

func NewJsonDbConnection(cfg *config.Config) (storage.StorageI, error) {

	clinicFile, err := os.Open(cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	branchFile, err := os.Open(cfg.BranchPath)
	if err != nil {
		return nil, err
	}

	return &Store{
		clinic: NewClinicRepo(cfg, clinicFile),
		branch: NewBranchRepo(cfg, branchFile),
	}, nil
}

func (s Store) Clinic() storage.ClinicRepoI {
	return s.clinic
}

func (s Store) Branch() storage.BranchRepoI {
	return s.branch
}
