package types

func NewMsgCompleteTodo(creator string, id uint64) *MsgCompleteTodo {
	return &MsgCompleteTodo{
		Creator: creator,
		Id:      id,
	}
}
