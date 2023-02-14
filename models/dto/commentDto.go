package dto

type CommentIdDto struct {
	Id uint `json:"id" binding:"required"`
}

type CommentDto struct {
	Vid     uint   `json:"vid" binding:"required"`
	Content string `json:"content" binding:"required,min=1"`
}
