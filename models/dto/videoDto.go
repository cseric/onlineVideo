package dto

type VideoIdDto struct {
	Id uint `json:"id" binding:"required"`
}

type AuditVideoVo struct {
	Id     uint   `json:"id" binding:"required"`
	Remark string `json:"remark" binding:"required"`
}

type UploadVideoDto struct {
	Title string `json:"title" binding:"required,min=1"`
	Cover string `json:"cover" binding:"required"`
	Path  string `json:"path" binding:"required"`
	Brief  string `json:"brief" binding:"required,min=1"`
}

type UpdateVideoDto struct {
	Id    uint   `json:"id" binding:"required"`
	Title string `json:"title" binding:"required,min=1"`
	Cover string `json:"cover" binding:"required"`
	Path  string `json:"path" binding:"required"`
	Brief  string `json:"brief" binding:"required,min=1"`
}
