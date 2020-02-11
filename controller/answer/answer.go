package answer

import (
	"github.com/gin-gonic/gin"
	"logger"
	"questions/db"
	"questions/model"
	"questions/util"
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
