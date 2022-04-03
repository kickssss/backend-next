package service

import (
	"context"

	"github.com/penguin-statistics/backend-next/internal/model"
	"github.com/penguin-statistics/backend-next/internal/repo"
)

type DropMatrixElementService struct {
	DropMatrixElementRepo *repo.DropMatrixElement
}

func NewDropMatrixElementService(dropMatrixElementRepo *repo.DropMatrixElement) *DropMatrixElementService {
	return &DropMatrixElementService{
		DropMatrixElementRepo: dropMatrixElementRepo,
	}
}

func (s *DropMatrixElementService) BatchSaveElements(ctx context.Context, elements []*model.DropMatrixElement, server string) error {
	return s.DropMatrixElementRepo.BatchSaveElements(ctx, elements, server)
}

func (s *DropMatrixElementService) DeleteByServer(ctx context.Context, server string) error {
	return s.DropMatrixElementRepo.DeleteByServer(ctx, server)
}

func (s *DropMatrixElementService) GetElementsByServer(ctx context.Context, server string) ([]*model.DropMatrixElement, error) {
	return s.DropMatrixElementRepo.GetElementsByServer(ctx, server)
}
