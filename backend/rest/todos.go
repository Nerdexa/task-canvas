package rest

import (
	"net/http"
	"task-canvas/config"
	"task-canvas/domain"
	db_driver "task-canvas/driver"
	"task-canvas/gateway"
	"task-canvas/logger"
	"task-canvas/useCase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type PostTodosRequest struct {
	Content   string `json:"content" validate:"required"`
	Completed bool   `json:"completed"`
}

type PutTodoRequest struct {
	Id        string `json:"id" param:"id" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Completed bool   `json:"completed"`
}

type PostTodosRequestResponse struct {
	Id string `json:"id"`
}

type DeleteTodoRequest struct {
	Id string `param:"id" validate:"required"`
}

func GetTodos(c echo.Context) error {
	todoDriver := db_driver.NewQuerier(config.PgPool)
	todoGateway := gateway.NewTodoGateway(todoDriver)
	todoUseCase := useCase.NewGetTodoUseCase(todoGateway)

	userIdStr := c.Get("userId").(string)
	userIdUuid, err := uuid.Parse(userIdStr)
	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todos, err := todoUseCase.Get(c.Request().Context(), domain.UserId(userIdUuid))
	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	restods := make([]Todo, 0, len(todos))
	for _, todo := range todos {
		restods = append(restods, Todo{
			Id:        uuid.UUID(todo.ID).String(),
			Content:   string(todo.Content),
			Completed: bool(todo.Completed),
		})
	}

	res := GetTodosResponse{
		Todos: restods,
	}

	return c.JSON(http.StatusOK, res)
}

func PostTodos(c echo.Context) error {
	reqTodo := new(PostTodosRequest)

	if err := c.Bind(reqTodo); err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(reqTodo); err != nil {
		return err
	}

	userIdStr := c.Get("userId").(string)
	userIdUuid, err := uuid.Parse(userIdStr)
	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	dbDriver := db_driver.NewQuerier(config.PgPool)
	todoIdGateway := gateway.NewTodoIdGateway()
	todoGateway := gateway.NewTodoGateway(dbDriver)
	todoUseCase := useCase.NewStoreTodoUseCase(todoGateway, todoIdGateway)

	todoId, err := todoUseCase.Store(ctx, domain.TodoContent(reqTodo.Content), domain.TodoCompleted(reqTodo.Completed), domain.UserId(userIdUuid))
	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := PostTodosRequestResponse{
		Id: uuid.UUID(todoId).String(),
	}

	return c.JSON(http.StatusOK, res)
}

func PutTodo(c echo.Context) error {
	reqTodo := new(PutTodoRequest)

	if err := c.Bind(reqTodo); err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(reqTodo); err != nil {
		return err
	}

	userIdStr := c.Get("userId").(string)
	userIdUuid, err := uuid.Parse(userIdStr)
	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todoDriver := db_driver.NewQuerier(config.PgPool)
	todoGateway := gateway.NewTodoGateway(todoDriver)
	todoUseCase := useCase.NewUpdateTodoUseCase(todoGateway)

	err = todoUseCase.UpdateTodoUseCase(c.Request().Context(), domain.Todo{
		ID:        domain.TodoId(uuid.MustParse(reqTodo.Id)),
		Content:   domain.TodoContent(reqTodo.Content),
		Completed: domain.TodoCompleted(reqTodo.Completed),
		UserId:    domain.UserId(userIdUuid),
	})

	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func DeleteTodo(c echo.Context) error {
	reqTodo := new(DeleteTodoRequest)

	if err := c.Bind(reqTodo); err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(reqTodo); err != nil {
		return err
	}

	userIdStr := c.Get("userId").(string)
	userIdUuid, err := uuid.Parse(userIdStr)
	if err != nil {
		logger.Logger.Error("Failed to bind release: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todoDriver := db_driver.NewQuerier(config.PgPool)
	todoGateway := gateway.NewTodoGateway(todoDriver)
	todoUseCase := useCase.NewDeleteTodoUseCase(todoGateway)

	err = todoUseCase.Delete(c.Request().Context(), domain.TodoId(uuid.MustParse(reqTodo.Id)), domain.UserId(userIdUuid))
	if err != nil {
		logger.Logger.Error("Failed to useCase: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
