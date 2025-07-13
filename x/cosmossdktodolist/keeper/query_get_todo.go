package keeper

import (
	"context"

	"cosmos-sdk-todo-list/x/cosmossdktodolist/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) GetTodo(ctx context.Context, req *types.QueryGetTodoRequest) (*types.QueryGetTodoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryGetTodoResponse{}, nil
}
