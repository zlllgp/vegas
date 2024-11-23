package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/yoda/query"
	"github.com/zlllgp/vegas/kitex_gen/api"
	"gorm.io/gen"
)

type RightService struct {
	rightRep repository.RightRepository
}

func NewRightService(rightRep repository.RightRepository) *RightService {
	return &RightService{rightRep: rightRep}
}

func (s *RightService) Find(ctx context.Context, userId int64, planId int64, pageRequest *api.PaginationRequest, searchFields ...*api.SearchField) ([]*entity.Right, error) {
	page, size := pageRequest.GetPage(), pageRequest.GetSize()
	offset := (page - 1) * size
	limit := size
	conditions, err := s.parseSearchFields(searchFields...)
	if err != nil {
		klog.Errorf("parse search fields failed: %v", err)
		return nil, err
	}

	result, err := s.rightRep.Query(ctx, int(offset), int(limit), conditions...)
	if err != nil {
		klog.Errorf("query rights failed: %v", err)
	}

	rights := make([]*entity.Right, len(result))
	for i := range result {
		rights[i] = entity.NewWithModel(result[i])
	}
	return rights, nil
}

func (s *RightService) parseSearchFields(searchFields ...*api.SearchField) ([]gen.Condition, error) {
	result := make([]gen.Condition, len(searchFields))
	for i := 0; i < len(searchFields); i++ {
		fieldName := searchFields[i].Field
		condition, err := s.parse(fieldName, searchFields[i].Operator, searchFields[i].Value)
		if err != nil {
			return nil, err
		}
		result[i] = condition
	}
	return result, nil
}

func (p *RightService) parse(fieldName string, operator api.SearchOperator, value string) (gen.Condition, error) {
	switch fieldName {
	case "userName":
		return entity.NewSearchFiled(query.Q.UserRight.UserName, fieldName, operator).ToGormCondition(value)
	default:
		return nil, fmt.Errorf("unsupported field name %s", fieldName)
	}
}
