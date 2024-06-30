package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"my_gql_server/graph/model"
)

// Mutation createTodoが呼ばれた際に実行されるコード
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// TODO: ユーザーから受け取ったリクエスト情報inputを使って、TODOをと登録し
	// その登録されたTODOの情報をmodel.TODO型の戻り値に入れて返却
	return nil, nil
}

// Query Todosが呼ばれた際に実行されるコード
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// TODO: レスポンスに含めるTODO一覧を、戻り値[]*model.TODOに入れて返却
	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
