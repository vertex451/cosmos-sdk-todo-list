package keeper

import (
	"context"

	"cosmos-sdk-todo-list/x/cosmossdktodolist/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListTodos(ctx context.Context, req *types.QueryListTodosRequest) (*types.QueryListTodosResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryListTodosResponse{}, nil
}
