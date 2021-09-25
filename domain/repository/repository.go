package repository

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/entity"
)

type FindAllObjUsecasesRepo interface {
  FindAllObjUsecases(ctx context.Context, objController *entity.ObjController) ([]*entity.ObjUsecase, error)
}
