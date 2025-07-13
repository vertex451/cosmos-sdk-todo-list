package types

func NewMsgCreateTodo(creator string, title string, description string, priority uint64) *MsgCreateTodo {
	return &MsgCreateTodo{
		Creator:     creator,
		Title:       title,
		Description: description,
		Priority:    priority,
	}
}
