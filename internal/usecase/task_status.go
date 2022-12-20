package usecase

import "context"

type TaskStatusUseCase struct {
	rp TaskStatusRp
}

func NewTaskStatusUseCase(rp TaskStatusRp) *TaskStatusUseCase {
	return &TaskStatusUseCase{rp: rp}
}

var _ TaskStatusContract = (*TaskStatusUseCase)(nil)

func (t *TaskStatusUseCase) CheckExistByID(ctx context.Context, id int64) bool {
	_, err := t.rp.CheckExistByID(ctx, id)
	if err != nil {
		return false
	}
	return true
}
