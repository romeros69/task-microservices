package usecase

import (
	"context"
	"task-microservices/internal/entity"
)

type (
	TaskContract interface {
		GetTasks(context.Context) ([]entity.Task, error)
		GetTaskByID(context.Context, int64) (entity.Task, error)
		DeleteTaskByID(context.Context, int64) error
		CreateTask(context.Context, entity.Task) (int64, error)
		ApproveTask(context.Context) error // TODO - по какому признаку?
	}

	TaskRp interface {
		GetTasks(context.Context) ([]entity.Task, error)
		GetTaskByID(context.Context, int64) (entity.Task, error)
		DeleteTaskByID(context.Context, int64) error
		CreateTask(context.Context, entity.Task) (int64, error)
		ApproveTask(context.Context) error // TODO - по какому признаку?
	}

	TaskStatusContract interface {
		CheckExistByID(context.Context, int64) bool
	}

	TaskStatusRp interface {
		CheckExistByID(context.Context, int64) (entity.TaskStatus, error)
	}
)
