package usecase

import (
	"context"
	"task-microservices/internal/entity"
)

type TaskUseCase struct {
	rp TaskRp
}

func NewTaskUseCase(rp TaskRp) *TaskUseCase {
	return &TaskUseCase{rp: rp}
}

var _ TaskContract = (*TaskUseCase)(nil)

func (t TaskUseCase) GetTasks(ctx context.Context) ([]entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) GetTaskByID(ctx context.Context, i int64) (entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) DeleteTaskByID(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) CreateTask(ctx context.Context, task entity.Task) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) ApproveTask(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
