package types

func NewMsgUpdateTodo(creator string, id uint64, title string, description string, priority uint64) *MsgUpdateTodo {
	return &MsgUpdateTodo{
		Creator:     creator,
		Id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	}
}
