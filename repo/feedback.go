package repo

import (
	"context"

	"github.com/mshsvnv/service/dto"
	"github.com/mshsvnv/service/model"
)

//go:generate mockery --name=IFeedbackrepo
type IFeedbackRepository interface {
	Create(ctx context.Context, feedback *model.Feedback) error
	Update(ctx context.Context, feedback *model.Feedback) error
	Remove(ctx context.Context, req *dto.RemoveFeedbackReq) error
	GetFeedback(ctx context.Context, req *dto.GetFeedbackReq) (*model.Feedback, error)
	GetFeedbacksByUserID(ctx context.Context, id int) ([]*model.Feedback, error)
	GetFeedbacksByRacketID(ctx context.Context, id int) ([]*model.Feedback, error)
}
