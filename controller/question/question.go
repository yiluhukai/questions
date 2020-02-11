package question

import (
	"logger"
	"questions/db"
	"questions/filter"
	"questions/gen_id"
	"questions/middleware/account"
	"questions/model"
	"questions/util"

	"github.com/gin-gonic/gin"
)

func QuestionSubmitHandle(c *gin.Context) {
	var question model.Question
	err := c.BindJSON(&question)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	logger.LogDebug("bind json succ, question:%#v", question)
	//  判断问题中是否包含非法字符
	result, hit := filter.Replace(question.Caption, "***")

	if hit {
		logger.LogDebug("hit filter data；%v", result)
		util.ResponseError(c, util.ErrCodeCaptionHit)
		return
	}

	result, hit = filter.Replace(question.Content, "***")

	if hit {
		logger.LogDebug("hit filter data；%v", result)
		util.ResponseError(c, util.ErrCodeContentHit)
		return
	}
	logger.LogDebug("filter data successfully")
	// 获取用户的id
	uid, err := account.GetUserId(c)
	if err != nil {
		logger.LogDebug("get userId failed:%v", err)
		util.ResponseError(c, util.ErrCodeUserNotExist)
		return
	}
	question.AuthorId = uid
	// 生成question_id
	qid, err := gen_id.GetID()
	if err != nil {
		logger.LogDebug("created question_id failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	question.QuestionId = int64(qid)

	err = db.CreateQuestion(&question)

	if err != nil {
		logger.LogError("create question failed：%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	logger.LogDebug("create question succ, question:%#v", question)
	util.ResponseSuccess(c, nil)
}

func QuestionDetailHandle(c *gin.Context) {
	question_id, err := util.GetQueryInt64(c, "question_id")
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	//使用问题id获取问题详情
	questionDetail, err := db.GetQuestion(question_id)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	//  通过类别id获取用户的类别信息
	categoryMap, err := db.GetCategory([]int64{questionDetail.CategoryId})

	if err != nil {
		logger.LogError("fetch categoryList failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	// 获取问题的作者信息
	userInfoList, err := db.GetUserInfoList([]int64{questionDetail.AuthorId})
	if err != nil {
		logger.LogError("fetch userInfoList failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	category, ok := categoryMap[questionDetail.CategoryId]
	if !ok {
		logger.LogError("fetch categoryList failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	var question model.ApiQuestionDetail
	question.CategoryName = category.CategoryName
	question.Question = *questionDetail
	for _, userInfo := range userInfoList {
		question.AuthorName = userInfo.Username
	}
	util.ResponseSuccess(c, question)
}
