package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"yiluhuakai/logger"
	"yiluhuakai/questions/model"
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
	sqlStr = `insert into comment_rel(comment_id,question_id,parent_id,level,reply_author_id,type) 
				values(?,?,?,?,?,?)`
	_, err = tx.Exec(sqlStr, comment.CommentId, comment.QuestionId, comment.ParentId, 1, comment.ReplyAuthorId, comment.Type)
	if err != nil {
		logger.LogError("insert into  comment_rel count failed, comment:%#v err:%v", comment, err)
		_ = tx.Rollback()
		return
	}

	//更新问题或者答案的评论数量
	if comment.Type == model.CommentType {
		//更新答案
		sqlStr = "update  answer  set comment_count = comment_count + 1 where author_id =?"
	} else {
		//问题
		sqlStr = "update  question  set comment_count = comment_count + 1 where question_id =?"
	}

	_, err = tx.Exec(sqlStr, comment.QuestionId)

	if err != nil {
		logger.LogError("update comment_count failde:%v", err)
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
			_ = tx.Rollback()
		}
	}()
	sqlStr := "insert into comment(comment_id,content,author_id) values(?,?,?)"
	_, err = tx.Exec(sqlStr, replyComment.CommentId, replyComment.Content, replyComment.AuthorId)

	if err != nil {
		logger.LogError("insert into comment failed:%v", err)
		_ = tx.Rollback()
		return
	}

	// 维护comment_rel表
	sqlStr = "insert into comment_rel(question_id,comment_id,reply_author_id,level,parent_id) values(?,?,?,?,?)"

	_, err = tx.Exec(sqlStr, replyComment.QuestionId, replyComment.CommentId, replyComment.ReplyAuthorId, 2, replyComment.ParentId)

	if err != nil {
		logger.LogError("insert into comment_rel failed:%v", err)
		_ = tx.Rollback()
		return
	}
	// 更新评论被评论的数量

	sqlStr = "update comment set comment_count =comment_count + 1 where comment_id = ?"
	_, err = tx.Exec(sqlStr, replyComment.ReplyCommentId)
	if err != nil {
		logger.LogError("update comment_count failed:%v", err)
		_ = tx.Rollback()
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

// 点赞

func AddOrCancelLike(like *model.Like) (err error) {
	var id int
	sqlStr := "select id from like_owner_rel where like_id =? and user_id =?"
	err = db.Get(&id, sqlStr, like.Id, like.UserId)

	if err != nil && err != sql.ErrNoRows {
		logger.LogError("query like record failed:%v", err)
		return
	}
	if err == sql.ErrNoRows {
		logger.LogDebug("doesn't add like for this answer or comment")

		//判断是对答案还是对评论的点赞
		if like.Type == model.AnswerTypeForLIke {
			err = UpdateAnswerLike(like, true)
		} else {
			err = UpdateCommentLike(like, true)
		}
		return
	}

	//记录存在，取消点赞
	if like.Type == model.AnswerTypeForLIke {
		err = UpdateAnswerLike(like, false)
	} else {
		err = UpdateCommentLike(like, false)
	}
	return
}

func UpdateCommentLike(like *model.Like, isInCreasement bool) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		logger.LogError("start tx failed :%v", err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	var sqlStr string
	if isInCreasement {
		sqlStr = "update comment set comment_count  =  comment_count +1  where  comment_id =?"
	} else {
		sqlStr = "update comment set comment_count  =  comment_count -1  where  comment_id =?"
	}
	_, err = tx.Exec(sqlStr, like.Id)

	if err != nil {
		logger.LogError("update comment' comment_count failed:%v", err)
		_ = tx.Rollback()
		return
	}

	// 维护关系表
	if isInCreasement {
		sqlStr = "insert into like_owner_rel(type,like_id,user_id)  values(?,?,?)"

	} else {
		sqlStr = "delete from like_owner_rel where type=? and like_id  = ? and  user_id = ?"
	}
	_, err = tx.Exec(sqlStr, 0, like.Id, like.UserId)
	if err != nil {
		logger.LogDebug("insert into like_owner_rel failed:%v", err)
		_ = tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}
