package db

import (
	"logger"
	"questions/model"
)

func CreatePostComment(comment *model.Comment) (err error) {
	tx, err := db.Beginx()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if err != nil {
		logger.LogError("start tx failed:%v", err)
		return
	}
	sqlStr := "insert into comment(comment_id,author_id,content) values(?,?,?)"
	_, err = tx.Exec(sqlStr, comment.CommentId, comment.AuthorId, comment.Content)
	if err != nil {
		logger.LogError("create  comment for question %v failed:%v", comment.QuestionId, err)
		_ = tx.Rollback()
		return
	}
	// 维护coment_rel表
	sqlStr = "insert into comment_rel(comment_id,question_id,parent_id,level,reply_author_id) values(?,?,?,?,?)"
	_, err = tx.Exec(sqlStr, comment.CommentId, comment.QuestionId, comment.ParentId, 1, comment.ReplyAuthorId)
	if err != nil {
		logger.LogError("insert into  comment_rel count failed, comment:%#v err:%v", comment, err)
		_ = tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}
