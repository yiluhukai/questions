package answer

import (
	"fmt"
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

func AnswerListHandle(c *gin.Context) {
	questionId, err := util.GetQueryInt64(c, "question_id")
	if err != nil {
		logger.LogError("fetch answerList failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// offset和 limit用来分页
	offset, err := util.GetQueryInt64(c, "offset")
	if err != nil {
		logger.LogError("get offset failed, err:%v", err)
		offset = 0
	}

	limit, err := util.GetQueryInt64(c, "limit")
	if err != nil {
		logger.LogError("get limit failed, err:%v", err)
		limit = 10
	}

	// 通过question_id获取答案的id列表
	answerIdList, err := db.GetAnswerIdList(questionId, offset, limit)
	if err != nil {
		logger.LogError("get answers failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	if len(answerIdList) == 0 {
		logger.LogDebug("this question no answers")
		util.ResponseSuccess(c, model.ApiAnswerList{AnswerList: []*model.ApiAnswer{}, TotalCount: 0})
		return
	}

	//通过answerIdList 获取answerList
	answerList, err := db.GetAnswerList(answerIdList)

	if err != nil {
		logger.LogError("get answers by answerIdList failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	var userIdList []int64
	for _, v := range answerList {
		userIdList = append(userIdList, v.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.LogError("db.GetUserInfoList failed, user_ids:%v err:%v",
			userIdList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	apiAnswerList := &model.ApiAnswerList{}
	for _, v := range answerList {
		apiAnswer := &model.ApiAnswer{}
		apiAnswer.Answer = *v
		apiAnswer.Answer.AnswerIdStr = fmt.Sprintf("%v", v.AnswerId)

		for _, user := range userInfoList {
			if user.UserId == v.AuthorId {
				apiAnswer.AuthorName = user.Username
				break
			}
		}

		apiAnswerList.AnswerList = append(apiAnswerList.AnswerList, apiAnswer)
	}
	answerCount, err := db.GetAnswerCount(questionId)
	if err != nil {
		logger.LogError("get answer count failed,%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	apiAnswerList.TotalCount = int32(answerCount)
	util.ResponseSuccess(c, apiAnswerList)
}

func AnswerCommitHandle(c *gin.Context) {
	var aw model.Answer
	err := c.BindJSON(&aw)
	if err != nil {
		logger.LogError("bind answer failed:err", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// 将问题的id转化成int64
	questionId, err := strconv.ParseInt(aw.QuestionId, 10, 64)
	if err != nil {
		logger.LogError("parse question_id to int64 failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	// 验证提交的答案

	if len(aw.Content) == 0 {
		logger.LogError("length of answer's content is zero %v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// 对提交的内容转移，防止xss攻击
	aw.Content = html.EscapeString(aw.Content)

	// 从session中获取用户的id

	uid, err := account.GetUserId(c)

	if err != nil {
		logger.LogError("user does not found:%v", err)
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	aw.AuthorId = uid

	//生成答案的id
	answerId, err := gen_id.GetID()

	if err != nil {
		logger.LogError("create answer_id failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	aid := int64(answerId)

	aw.AnswerId = aid

	err = db.CreateAnswer(&aw, questionId)

	if err != nil {
		logger.LogError("save anser failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	util.ResponseSuccess(c, nil)
}
