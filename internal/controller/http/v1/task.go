package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-microservices/internal/entity"
	"task-microservices/internal/usecase"
)

type taskRoutes struct {
	auth *AuthMiddle
	t    usecase.TaskContract
}

func newTaskRoutes(handler *gin.RouterGroup, t usecase.TaskContract, auth *AuthMiddle) {
	tr := &taskRoutes{t: t, auth: auth}

	handler.GET("/tasks", tr.getTasks)
	handler.GET("/task/:id", tr.getTaskByID)
	handler.DELETE("/task/:id", tr.deleteTaskByID)
	handler.POST("/task", tr.createTask)
}

// @Summary GetAllTasks
// @Tags task
// @Security ApiKeyAuth
// @Description Get all tasks
// @ID get-all-tasks
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Task
// @Failure 500 {object} errResponse
// @Router /api/v1/tasks [get]
func (t *taskRoutes) getTasks(c *gin.Context) {
	err := t.auth.authMiddle(c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	listTasks, err := t.t.GetTasks(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, listTasks)
}

// @Summary GetTaskById
// @Tags task
// @Security ApiKeyAuth
// @Description Get task by id
// @ID get-task-by-id
// @Accept json
// @Produce json
// @Param id path string true "Enter id task"
// @Success 200 {object} entity.Task
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/task/{id} [get]
func (t *taskRoutes) getTaskByID(c *gin.Context) {
	err := t.auth.authMiddle(c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	task, err := t.t.GetTaskByID(c.Request.Context(), taskID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

// @Summary DeleteTask
// @Tags task
// @Security ApiKeyAuth
// @Description Delete task by id
// @ID delete-task
// @Accept json
// @Produce json
// @Param id path string true "Enter id task"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/task/{id} [delete]
func (t *taskRoutes) deleteTaskByID(c *gin.Context) {
	err := t.auth.authMiddle(c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = t.t.DeleteTaskByID(c.Request.Context(), taskID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary CreateTask
// @Tags task
// @Security ApiKeyAuth
// @Description Create new task
// @ID create-task
// @Accept json
// @Produce json
// @Param input body taskRequestDTO true "Enter author and status id of new task"
// @Success 201 {object} int
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/task [post]
func (t *taskRoutes) createTask(c *gin.Context) {
	err := t.auth.authMiddle(c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	req := new(taskRequestDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := t.t.CreateTask(c.Request.Context(), entity.Task{
		Author:   req.Author,
		StatusID: req.StatusID,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}
