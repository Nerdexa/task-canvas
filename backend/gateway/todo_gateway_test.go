package gateway

import (
	"context"
	"reflect"
	"task-canvas/domain"
	db_driver "task-canvas/driver"
	sqlc "task-canvas/driver/generated"
	mock_db_driver "task-canvas/mock/driver"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/mock/gomock"
)

func TestTodoGateway_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_db_driver.NewMockQuerier(ctrl)

	userId := domain.NewUserId()
	now := time.Now()

	mockDriver.EXPECT().FindTodo(context.Background(), uuid.UUID(userId)).Return([]sqlc.FindTodoRow{
		{TodoID: uuid.MustParse("56CD2629-3035-47EB-AA41-C8F25D5FC954"), Content: "title1", Completed: true, UserID: uuid.UUID(userId), Email: "test@test.com", PasswordHash: "password", CreatedAt: pgtype.Timestamp{Time: now.Add(-3 * time.Hour), Valid: true}},
		{TodoID: uuid.MustParse("97A46613-0E12-4A7F-B40E-57CF55EEFC84"), Content: "title2", Completed: true, UserID: uuid.UUID(userId), Email: "test@test.com", PasswordHash: "password", CreatedAt: pgtype.Timestamp{Time: now, Valid: true}},
		{TodoID: uuid.MustParse("10CE7F14-8B10-45C8-87E1-810008AE1ED7"), Content: "title3", Completed: true, UserID: uuid.UUID(userId), Email: "test@test.com", PasswordHash: "password", CreatedAt: pgtype.Timestamp{Time: now.Add(+10 * time.Hour), Valid: true}},
	}, nil)

	type fields struct {
		db_driver db_driver.Querier
	}
	type args struct {
		ctx    context.Context
		userId domain.UserId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Todo
		wantErr bool
	}{
		{
			name: "Todo一覧の取得",
			fields: fields{
				db_driver: mockDriver,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
			},
			want: []domain.Todo{
				{ID: domain.TodoId(uuid.MustParse("10CE7F14-8B10-45C8-87E1-810008AE1ED7")), Content: "title3", Completed: true, UserId: userId},
				{ID: domain.TodoId(uuid.MustParse("97A46613-0E12-4A7F-B40E-57CF55EEFC84")), Content: "title2", Completed: true, UserId: userId},
				{ID: domain.TodoId(uuid.MustParse("56CD2629-3035-47EB-AA41-C8F25D5FC954")), Content: "title1", Completed: true, UserId: userId},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &TodoGateway{
				db_driver: tt.fields.db_driver,
			}
			got, err := g.Get(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoGateway.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoGateway.Get() = %v, want %v", got, tt.want)
				diff := cmp.Diff(got, tt.want)
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestTodoGateway_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockTx := mock_db_driver.NewMockTx(ctrl)
	mockDriver := mock_db_driver.NewMockQuerier(ctrl)
	mockDriver.EXPECT().Begin(ctx).Return(mockTx, nil).Times(1)
	mockTx.EXPECT().Commit(ctx).Return(nil).Times(1)
	mockDriver.EXPECT().WithTx(mockTx).Return(mockDriver).Times(1)

	userId := domain.NewUserId()
	todoIdUuid := uuid.New()

	todo := domain.Todo{
		ID:        domain.TodoId(todoIdUuid),
		Content:   "title1",
		Completed: true,
		UserId:    userId,
	}
	mockDriver.EXPECT().InsertTodo(ctx, sqlc.InsertTodoParams{
		ID:        todoIdUuid,
		Content:   string(todo.Content),
		Completed: bool(todo.Completed),
	}).Times(1).Return(nil)

	mockDriver.EXPECT().InsertUserTodo(ctx, sqlc.InsertUserTodoParams{
		UserID: uuid.UUID(userId),
		TodoID: todoIdUuid,
	}).Times(1).Return(nil)

	type fields struct {
		db_driver db_driver.Querier
	}
	type args struct {
		ctx  context.Context
		todo domain.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Todoの作成",
			fields: fields{
				db_driver: mockDriver,
			},
			args: args{
				ctx:  ctx,
				todo: todo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &TodoGateway{
				db_driver: tt.fields.db_driver,
			}
			if err := g.Store(tt.args.ctx, tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("TodoGateway.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoGateway_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockTodoDriver := mock_db_driver.NewMockQuerier(ctrl)
	todoUuid := uuid.MustParse("56CD2629-3035-47EB-AA41-C8F25D5FC954")
	userId := domain.NewUserId()
	mockTodoDriver.EXPECT().UpdateTodo(ctx, sqlc.UpdateTodoParams{
		TodoID:    todoUuid,
		Content:   "title1",
		Completed: true,
		UserID:    uuid.UUID(userId),
	}).Times(1)

	type fields struct {
		db_driver db_driver.Querier
	}
	type args struct {
		ctx  context.Context
		todo domain.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Todoの更新",
			fields: fields{
				db_driver: mockTodoDriver,
			},
			args: args{
				ctx: ctx,
				todo: domain.Todo{
					ID:        domain.TodoId(todoUuid),
					Content:   "title1",
					Completed: true,
					UserId:    userId,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &TodoGateway{
				db_driver: tt.fields.db_driver,
			}
			if err := g.Update(tt.args.ctx, tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("TodoGateway.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoGateway_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userId := domain.NewUserId()

	mockTx := mock_db_driver.NewMockTx(ctrl)
	mockTx.EXPECT().Commit(ctx).Return(nil).Times(1)

	mockDriver := mock_db_driver.NewMockQuerier(ctrl)
	mockDriver.EXPECT().Begin(ctx).Return(mockTx, nil).Times(1)
	mockDriver.EXPECT().WithTx(mockTx).Return(mockDriver).Times(1)
	mockDriver.EXPECT().DeleteTodo(ctx, sqlc.DeleteTodoParams{
		TodoID: uuid.UUID(domain.TodoId(uuid.MustParse("56CD2629-3035-47EB-AA41-C8F25D5FC954"))),
		UserID: uuid.UUID(userId),
	}).Return(nil).Times(1)

	type fields struct {
		db_driver db_driver.Querier
	}
	type args struct {
		ctx    context.Context
		id     domain.TodoId
		userId domain.UserId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Todoの削除",
			fields: fields{
				db_driver: mockDriver,
			},
			args: args{
				ctx:    ctx,
				id:     domain.TodoId(uuid.MustParse("56CD2629-3035-47EB-AA41-C8F25D5FC954")),
				userId: userId,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &TodoGateway{
				db_driver: tt.fields.db_driver,
			}
			if err := g.Delete(tt.args.ctx, tt.args.id, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("TodoGateway.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
