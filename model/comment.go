package model

import (
	"time"
)

const (
	QuestionType = iota
	CommentType
)

type Comment struct {
	CommentId         int64     `json:"-" db:"comment_id" `
	Content           string    `db:"content" json:"content"`
	AuthorId          int64     `db:"author_id" json:"author_id"`
	LikeCount         int       `db:"like_count" json:"like_count"`
	CommentCount      int       `db:"comment_count" json:"comment_count"`
	CreateTime        time.Time `db:"create_time" json:"create_time"`
	ParentId          int64     `db:"parent_id" `
	ParentIdStr       string    `json:"parent_id"`
	QuestionId        int64     `db:"question_id"`
	QuestionIdStr     string    `json:"question_id"`
	ReplyAuthorId     int64     `db:"reply_author_id" json:"reply_author_id"`
	ReplyCommentId    int64     `db:"reply_comment_id"`
	ReplyCommentIdStr string    `json:"reply_comment_id"`
	AuthorName        string    `json:"author_name"`
	ReplyAuthorName   string    `json:"reply_author_name"`
	CommentIdStr      string    `json:"comment_id"`
	Type              int       `json:"type" db:"type"`
}

type ApiCommentList struct {
	CommentList []*Comment `json:"comment_list"`
	Count       int64      `json:"count"`
}
