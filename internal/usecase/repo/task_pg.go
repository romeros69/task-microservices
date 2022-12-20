package repo

import (
	"context"
	"task-microservices/internal/entity"
	"task-microservices/internal/usecase"
	"task-microservices/pkg/postgres"
)

type TaskRepo struct {
	pg *postgres.Postgres
}

func NewTaskRepo(pg *postgres.Postgres) *TaskRepo {
	return &TaskRepo{pg: pg}
}

var _ usecase.TaskRp = (*TaskRepo)(nil)

func (t TaskRepo) GetTasks(ctx context.Context) ([]entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) GetTaskByID(ctx context.Context, i int64) (entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) DeleteTaskByID(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) CreateTask(ctx context.Context, task entity.Task) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) ApproveTask(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
