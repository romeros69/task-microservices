package repo

import (
	"context"
	"fmt"
	"strconv"
	"task-microservices/internal/entity"
	"task-microservices/internal/usecase"
	"task-microservices/pkg/postgres"
	"time"
)

type TaskRepo struct {
	pg *postgres.Postgres
}

func NewTaskRepo(pg *postgres.Postgres) *TaskRepo {
	return &TaskRepo{pg: pg}
}

var _ usecase.TaskRp = (*TaskRepo)(nil)

func (t *TaskRepo) GetTasks(ctx context.Context) ([]entity.Task, error) {
	query := `SELECT * FROM task`
	rows, err := t.pg.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err = rows.Scan(
			&task.ID,
			&task.CreationDate,
			&task.Author,
			&task.StatusID,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing task: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TaskRepo) GetTaskByID(ctx context.Context, id int64) (entity.Task, error) {
	query := "SELECT * FROM task WHERE id = $1"
	rows, err := t.pg.Pool.Query(ctx, query, id)
	if err != nil {
		return entity.Task{}, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return entity.Task{}, fmt.Errorf("there is no task with this id: %s", strconv.FormatInt(id, 10))
	}
	var task entity.Task
	err = rows.Scan(
		&task.ID,
		&task.CreationDate,
		&task.Author,
		&task.StatusID,
	)
	if err != nil {
		return entity.Task{}, fmt.Errorf("error parsing task: %w", err)
	}
	return task, nil
}

func (t *TaskRepo) DeleteTaskByID(ctx context.Context, id int64) error {
	query := `DELETE FROM task WHERE id = $1`
	rows, err := t.pg.Pool.Query(ctx, query, id)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (t *TaskRepo) CreateTask(ctx context.Context, task entity.Task) (int64, error) {
	query := "INSERT INTO task (creation_date, author, status_id) VALUES ($1, $2, $3) RETURNING id"
	rows, err := t.pg.Pool.Query(ctx, query, time.Now(), task.Author, task.StatusID)
	if err != nil {
		return -1, fmt.Errorf("cannot execute query: %w", err)
	}
	var id int64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return -1, fmt.Errorf("error parsing id task: %w", err)
		}
	}
	defer rows.Close()
	return id, nil
}

func (t *TaskRepo) ApproveTask(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
