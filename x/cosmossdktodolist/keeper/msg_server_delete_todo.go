package keeper

import (
	"context"

	"cosmos-sdk-todo-list/x/cosmossdktodolist/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) DeleteTodo(ctx context.Context, msg *types.MsgDeleteTodo) (*types.MsgDeleteTodoResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgDeleteTodoResponse{}, nil
}
