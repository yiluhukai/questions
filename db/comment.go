package db

import (
	"github.com/jmoiron/sqlx"
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

// 获取一级评论

func GetCommentsList(question_id, limit, offset int64) (commentList []*model.Comment, count int64, err error) {
	var commentIdList []int64
	sqlStr := "select comment_id from comment_rel  where question_id = ? and level = ? order by create_time desc limit ?,?"
	logger.LogDebug("question_id =%v ,limit =%v ,offset =%v", question_id, limit, offset)
	err = db.Select(&commentIdList, sqlStr, question_id, 1, offset, limit)
	if err != nil {
		logger.LogError("fetch commentIdList failed,%v", err)
		return
	}
	if len(commentIdList) == 0 {
		commentList = make([]*model.Comment, 0)
		return
	}
	// 根据id获取coment的信息
	sqlStr = "select comment_id, content, author_id, like_count, comment_count,create_time from comment where comment_id in (?)"
	var commentIds []interface{}
	for _, cid := range commentIdList {
		commentIds = append(commentIds, cid)
	}

	query, args, err := sqlx.In(sqlStr, commentIds)

	err = db.Select(&commentList, query, args...)
	if err != nil {
		logger.LogError("query commentList from comment failed:%v", err)
		return
	}
	// 获取数量
	sqlStr = "select count(comment_id) from comment_rel where question_id = ? and level =?"
	err = db.Get(&count, sqlStr, question_id, 1)
	if err != nil {
		logger.LogDebug("get comment count failed:%v", err)
		return
	}
	return
}

// 回复一条评论

func ReplyComment(replyComment *model.Comment) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		logger.LogError("start tx failed:%v", err)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	sqlStr := "insert into comment(comment_id,content,author_id) values(?,?,?)"
	_, err = tx.Exec(sqlStr, replyComment.CommentId, replyComment.Content, replyComment.AuthorId)

	if err != nil {
		logger.LogError("insert into comment failed:%v", err)
		err = tx.Rollback()
		return
	}

	// 维护comment_rel表
	sqlStr = "insert into comment_rel(question_id,comment_id,reply_author_id,level,parent_id) values(?,?,?,?,?)"

	_, err = tx.Exec(sqlStr, replyComment.QuestionId, replyComment.CommentId, replyComment.ReplyAuthorId, 2, replyComment.ParentId)

	if err != nil {
		logger.LogError("insert into comment_rel failed:%v", err)
		err = tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}

//返回回复的评论列表
func GetReplyCommentList(parentId, offset, limit int64) (replyComments []*model.Comment, count int64, err error) {

	var commentIdList []int64
	sqlStr := "select comment_id from comment_rel  where parent_id = ? and level = ? order by create_time desc limit ?,?"

	err = db.Select(&commentIdList, sqlStr, parentId, 2, offset, limit)
	if err != nil {
		logger.LogError("fetch commentIdList failed,%v", err)
		return
	}
	if len(commentIdList) == 0 {
		replyComments = make([]*model.Comment, 0)
		return
	}
	// 根据id获取coment的信息
	sqlStr = "select comment_id, content, author_id, like_count, comment_count,create_time from comment where comment_id in (?)"
	var commentIds []interface{}
	for _, cid := range commentIdList {
		commentIds = append(commentIds, cid)
	}

	query, args, err := sqlx.In(sqlStr, commentIds)

	err = db.Select(&replyComments, query, args...)
	if err != nil {
		logger.LogError("query commentList from comment failed:%v", err)
		return
	}
	// 获取数量
	sqlStr = "select count(comment_id) from comment_rel where parent_id = ? and level =?"
	err = db.Get(&count, sqlStr, parentId, 2)
	if err != nil {
		logger.LogDebug("get comment count failed:%v", err)
		return
	}
	return
}
