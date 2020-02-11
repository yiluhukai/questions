package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logger"
	"time"

	"questions/db"
	"questions/model"
	"questions/util"
	"strconv"
)

func GetCategoryListHandle(c *gin.Context) {
	categoryList, err := db.GetCategoryList()
	if err != nil {
		fmt.Printf("fetch categoryList failed,%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, categoryList)
}

func GetQuestionListHandle(c *gin.Context) {
	//获取种类的id
	categoryIdStr, ok := c.GetQuery("category_id")
	if !ok {
		logger.LogDebug("get category_id failed")
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		logger.LogError("parese int failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	// 通过id获取questions

	questionList, err := db.GetQuestionList(categoryId)
	if err != nil {
		logger.LogError("get questionList by id failed :%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	if len(questionList) == 0 {
		logger.LogDebug("questionList of this category is empty")
		questionList = make([]*model.Question, 0)
		util.ResponseSuccess(c, questionList)
		return
	}
	// 获取问题的作者信息
	var userIdList []int64
	userIdMap := make(map[int64]bool, 20)
	for _, question := range questionList {
		_, ok := userIdMap[question.AuthorId]
		if ok {
			continue
		}
		userIdMap[question.AuthorId] = true
		userIdList = append(userIdList, question.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.LogError("get userinfoList by userIdList failed；%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	var apiQuestionList []model.ApiQuestion
	for _, question := range questionList {
		var apiQuestion model.ApiQuestion
		apiQuestion.Question = *question
		apiQuestion.CreateTimeStr = question.CreateTime.Format(time.RFC822)
		for _, userInfo := range userInfoList {
			if question.AuthorId == userInfo.UserId {
				apiQuestion.AuthorName = userInfo.Username
				break
			}
		}
		apiQuestionList = append(apiQuestionList, apiQuestion)
	}
	util.ResponseSuccess(c, apiQuestionList)
}
