package types

func NewMsgDeleteTodo(creator string, id uint64) *MsgDeleteTodo {
	return &MsgDeleteTodo{
		Creator: creator,
		Id:      id,
	}
}
