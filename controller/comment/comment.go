package comment

import (
	"github.com/gin-gonic/gin"
	"html"
	"logger"
	"questions/db"
	"questions/gen_id"
	"questions/middleware/account"
	"questions/model"
	"questions/util"
	"strconv"
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
