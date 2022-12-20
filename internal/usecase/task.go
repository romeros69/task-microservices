package usecase

import (
	"context"
	"fmt"
	"strconv"
	"task-microservices/internal/entity"
)

type TaskUseCase struct {
	rp           TaskRp
	taskStatusUC TaskStatusContract
}

func NewTaskUseCase(rp TaskRp, taskStatusUC TaskStatusContract) *TaskUseCase {
	return &TaskUseCase{rp: rp, taskStatusUC: taskStatusUC}
}

var _ TaskContract = (*TaskUseCase)(nil)

func (t *TaskUseCase) GetTasks(ctx context.Context) ([]entity.Task, error) {
	return t.rp.GetTasks(ctx)
}

func (t *TaskUseCase) GetTaskByID(ctx context.Context, id int64) (entity.Task, error) {
	return t.rp.GetTaskByID(ctx, id)
}

func (t *TaskUseCase) DeleteTaskByID(ctx context.Context, id int64) error {
	_, err := t.GetTaskByID(ctx, id)
	if err != nil {
		return fmt.Errorf("task with this id = %s does not exist", strconv.FormatInt(id, 10))
	}
	return t.rp.DeleteTaskByID(ctx, id)
}

func (t *TaskUseCase) CreateTask(ctx context.Context, task entity.Task) (int64, error) {
	existStatus := t.taskStatusUC.CheckExistByID(ctx, task.StatusID)
	if !existStatus {
		return -1, fmt.Errorf("task status with this id = %s does not exist", strconv.FormatInt(task.StatusID, 10))
	}
	return t.rp.CreateTask(ctx, task)
}

func (t *TaskUseCase) ApproveTask(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
