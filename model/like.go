package model

const (
	CommentTypeForLike = iota
	AnswerTypeForLIke
)

type Like struct {
	Id     int64  `db:"like_id"`
	Type   int    `json:"type" db:"type"`
	UserId int64  `json:"user_id" db:"user_id"`
	IdStr  string `json:"id"`
}
