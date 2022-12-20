package repo

import (
	"context"
	"fmt"
	"strconv"
	"task-microservices/internal/entity"
	"task-microservices/internal/usecase"
	"task-microservices/pkg/postgres"
)

type TaskStatusRepo struct {
	pg *postgres.Postgres
}

func NewTaskStatusRepo(pg *postgres.Postgres) *TaskStatusRepo {
	return &TaskStatusRepo{pg: pg}
}

var _ usecase.TaskStatusRp = (*TaskStatusRepo)(nil)

func (t *TaskStatusRepo) CheckExistByID(ctx context.Context, id int64) (entity.TaskStatus, error) {
	query := "SELECT * FROM task_status WHERE id = $1"
	rows, err := t.pg.Pool.Query(ctx, query, id)
	if err != nil {
		return entity.TaskStatus{}, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return entity.TaskStatus{}, fmt.Errorf("there is no task status with this id: %s", strconv.FormatInt(id, 10))
	}
	var taskStatus entity.TaskStatus
	err = rows.Scan(
		&taskStatus.ID,
		&taskStatus.Status,
	)
	if err != nil {
		return entity.TaskStatus{}, fmt.Errorf("error parsing task status: %w", err)
	}
	return taskStatus, nil
}
