package db

import (
	"logger"
	"questions/model"
)

func CreateQuestion(question *model.Question) (err error) {
	sqlStr := "insert into question(question_id,caption,content,author_id,category_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, question.QuestionId, question.Caption, question.Content, question.AuthorId, question.CategoryId)
	return
}

// type Question struct {
// 	QuestionId    int64     `json:"question_id_number" db:"question_id"`
// 	Caption       string    `json:"caption" db:"caption"`
// 	Content       string    `json:"content" db:"content"`
// 	AuthorId      int64     `json:"author_id_number" db:"author_id"`
// 	CategoryId    int64     `json:"category_id" db:"category_id"`
// 	Status        int32     `json:"status" db:"status"`
// 	CreateTime    time.Time `json:"-" db:"create_time"`
// 	CreateTimeStr string    `json:"create_time"`
// 	QuestionIdStr string    `json:"question_id"`
// 	AuthorIdStr   string    `json:"author_id"`
// }
func GetQuestionList(category_id int64) (questestionList []*model.Question, err error) {
	sqlStr := "select question_id,caption,content,author_id,category_id,create_time from question where category_id = ?"
	err = db.Select(&questestionList, sqlStr, category_id)
	return
}

func GetQuestion(question_id int64) (questionDetail *model.Question, err error) {
	questionDetail = &model.Question{}
	sqlStr := "select question_id,caption,content,author_id,category_id,create_time from question where question_id = ?"
	err = db.Get(questionDetail, sqlStr, question_id)
	if err != nil {
		logger.LogError("fetch question deatil via question_id  failed:%v", err)
		return
	}
	return
}
