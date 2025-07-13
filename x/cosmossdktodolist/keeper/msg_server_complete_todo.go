package keeper

import (
	"context"

	"cosmos-sdk-todo-list/x/cosmossdktodolist/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) CompleteTodo(ctx context.Context, msg *types.MsgCompleteTodo) (*types.MsgCompleteTodoResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgCompleteTodoResponse{}, nil
}
