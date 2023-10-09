package storage

import "simple-clinic/models"

type StorageI interface {
	Clinic() ClinicRepoI
	Branch() BranchRepoI
}

type ClinicRepoI interface {
	Create(req *models.CreateClinic) (*models.Clinic, error)
	GetByID(req *models.ClinicPrimaryKey) (*models.Clinic, error)
	GetList(req *models.GetListClinicRequest) (*models.GetListClinicResponse, error)
	Update(req *models.UpdateClinic) (*models.Clinic, error)
	Delete(req *models.ClinicPrimaryKey) error
}

type BranchRepoI interface {
	Create(req *models.CreateBranch) (*models.Branch, error)
	GetByID(req *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(req *models.UpdateBranch) (*models.Branch, error)
	Delete(req *models.BranchPrimaryKey) error
}
