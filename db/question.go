package db

import "questions/model"

func CreateQuestion(question *model.Question) (err error) {
	sqlStr := "insert into question(question_id,caption,content,author_id,category_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, question.QuestionId, question.Caption, question.Content, question.AuthorId, question.CategoryId)
	return
}
