package db

import (
	"github.com/jmoiron/sqlx"
	"logger"
	"questions/model"
)

func GetAnswerIdList(qid int64, offset int64, limit int64) (answerIdList []int64, err error) {
	sqlStr := "select answer_id from question_answer_rel where question_id = ? order by  id desc limit ?,?"
	err = db.Select(&answerIdList, sqlStr, qid, offset, limit)
	if err != nil {
		logger.LogError("get answerList failed ；%v", err)
		return
	}
	return
}

func GetAnswerList(answerIdList []int64) (answerList []*model.Answer, err error) {
	sqlstr := `select
					answer_id, content, comment_count,
					voteup_count, author_id, status, can_comment,
					create_time, update_time
				 from
				 	answer where answer_id in(?)`
	var interfaceSlice []interface{}
	for _, c := range answerIdList {
		interfaceSlice = append(interfaceSlice, c)
	}

	insqlStr, params, err := sqlx.In(sqlstr, interfaceSlice)
	if err != nil {
		logger.LogError("sqlx.in failed, sqlstr:%v, err:%v", sqlstr, err)
		return
	}

	err = db.Select(&answerList, insqlStr, params...)
	if err != nil {
		logger.LogError("get answerList failed:%v", err)
		return
	}

	return
}

func GetAnswerCount(question_id int64) (count int, err error) {
	sqlstr := `select
							count(answer_id)
						from
							question_answer_rel
						where question_id=?`
	err = db.Get(&count, sqlstr, question_id)

	if err != nil {
		logger.LogError("fetech count of answer for quesrion =%v failed：%v", question_id, err)
		return
	}
	return
}
