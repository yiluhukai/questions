package comment

import (
	"fmt"
	"html"
	"strconv"
	"yiluhuakai/logger"
	"yiluhuakai/questions/db"
	"yiluhuakai/questions/gen_id"
	"yiluhuakai/questions/middleware/account"
	"yiluhuakai/questions/model"
	"yiluhuakai/questions/util"

	"github.com/gin-gonic/gin"
)

const (
	MinCommentContentSize = 10
)

func PostCommentHandle(c *gin.Context) {

	// 获取提交的评论内容,question_id,author_id,reply_author_id
	var comment model.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		logger.LogError("post comment failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// question_id 类型转化

	comment.QuestionId, err = strconv.ParseInt(comment.QuestionIdStr, 10, 64)
	if err != nil {
		logger.LogError("parse to int failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	//  对评论的内容进行检验
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 {
		logger.LogError("len(comment.content) :%v, qid:%v， invalid param",
			len(comment.Content), comment.QuestionId)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// 从session中获取用户的id
	//接口测试
	uid, err := account.GetUserId(c)

	if err != nil {
		logger.LogError("fetch uid from session failed:%v", err)
		util.ResponseError(c, util.ErrCodeUserNotExist)
		return
	}
	//uid := int64(100)
	comment.AuthorId = uid
	//  检测评论的内容
	//1. 针对content做一个转义，防止xss漏洞
	comment.Content = html.EscapeString(comment.Content)

	// 生成一个comment_id
	c_id, err := gen_id.GetID()
	if err != nil {
		logger.LogError("get comment_id failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	comment_id := int64(c_id)

	comment.CommentId = comment_id

	err = db.CreatePostComment(&comment)
	if err != nil {
		logger.LogError("insert db failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, nil)
}

func CommentListHandle(c *gin.Context) {
	// 获取question_id、offset、limit
	question_id, err := util.GetQueryInt64(c, "question_id")
	if err != nil {
		logger.LogError("questin_id is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	offset, err := util.GetQueryInt64(c, "offset")
	if err != nil {
		logger.LogError("offset dosen't exist and user default value:%v", err)
		offset = 0
	}
	limit, err := util.GetQueryInt64(c, "limit")
	if err != nil {
		logger.LogError("limit dosen't exist and user default value:%v", err)
		limit = 10
	}
	// 获取评论信息

	commentList, count, err := db.GetCommentsList(question_id, limit, offset)

	if err != nil {
		logger.LogError("fetch commentList failed %v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	for _, comment := range commentList {
		comment.CommentIdStr = fmt.Sprintf("%v", comment.CommentId)
	}
	var commentApiList model.ApiCommentList

	commentApiList.CommentList = commentList
	commentApiList.Count = count

	util.ResponseSuccess(c, commentApiList)
}

// 回复一条评论
func ReplyCommentHandle(c *gin.Context) {
	var replyComment model.Comment
	err := c.BindJSON(&replyComment)
	if err != nil {
		logger.LogError("params is  invalid,%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("bind params succ:%#v", replyComment)

	// 将reply_coment_id转成int64

	r_id, err := strconv.ParseInt(replyComment.ReplyCommentIdStr, 10, 64)

	if err != nil {
		logger.LogError("replay_comment_id is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	replyComment.ReplyCommentId = r_id
	q_id, err := strconv.ParseInt(replyComment.QuestionIdStr, 10, 64)

	if err != nil {
		logger.LogError("replay_comment_id is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	replyComment.QuestionId = q_id

	p_id, err := strconv.ParseInt(replyComment.ParentIdStr, 10, 64)

	if err != nil {
		logger.LogError("replay_comment_id is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	replyComment.ParentId = p_id
	// 检验评论的内容
	if len(replyComment.Content) <= MinCommentContentSize || replyComment.QuestionId == 0 {
		logger.LogError("len(comment.content) :%v, qid:%v， invalid param",
			len(replyComment.Content), replyComment.QuestionId)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// 获取评论id
	uid, err := account.GetUserId(c)

	if err != nil {
		logger.LogError("fetch uid from session failed:%v", err)
		util.ResponseError(c, util.ErrCodeUserNotExist)
		return
	}

	replyComment.AuthorId = uid
	// 生成一个comment_id
	c_id, err := gen_id.GetID()
	if err != nil {
		logger.LogError("get comment_id failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	comment_id := int64(c_id)

	replyComment.CommentId = comment_id

	err = db.ReplyComment(&replyComment)

	if err != nil {
		logger.LogError("create replayComment failed :%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, nil)
}

func CommentReplyListHandle(c *gin.Context) {
	// 获取question_id、offset、limit
	parent_id, err := util.GetQueryInt64(c, "parent_id")
	if err != nil {
		logger.LogError("parent_id is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	offset, err := util.GetQueryInt64(c, "offset")
	if err != nil {
		logger.LogError("offset dosen't exist and user default value:%v", err)
		offset = 0
	}
	limit, err := util.GetQueryInt64(c, "limit")
	if err != nil {
		logger.LogError("limit dosen't exist and user default value:%v", err)
		limit = 10
	}
	// 获取评论信息

	commentList, count, err := db.GetReplyCommentList(parent_id, offset, limit)

	if err != nil {
		logger.LogError("fetch commentList failed %v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	for _, comment := range commentList {
		comment.CommentIdStr = fmt.Sprintf("%v", comment.CommentId)
	}
	var commentApiList model.ApiCommentList

	commentApiList.CommentList = commentList
	commentApiList.Count = count

	util.ResponseSuccess(c, commentApiList)
}

// 添加点赞
func LikeHandle(c *gin.Context) {
	var like model.Like

	err := c.BindJSON(&like)

	if err != nil {
		logger.LogError("bind paramters failed :%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	//将id有字符串转成int64
	like.Id, err = strconv.ParseInt(like.IdStr, 10, 64)

	if err != nil {
		logger.LogError("parseInt from string failed :%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	like.UserId, err = account.GetUserId(c)

	if err != nil {
		logger.LogError("user doesn't login :%v", err)
		util.ResponseError(c, util.ErrCodeQuestionNotExist)
		return
	}

	err = db.AddOrCancelLike(&like)

	if err != nil {
		logger.LogError("add or cancel like failed %v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, nil)
	return
}
