package db

import (
	"github.com/jmoiron/sqlx"
	"yiluhuakai/logger"
	"yiluhuakai/questions/model"
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

func CreateAnswer(answer *model.Answer, questionId int64) (err error) {

	tx, err := db.Beginx()

	if err != nil {
		logger.LogError("created tx failed:%v", err)
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	sqlStr := "insert into answer(answer_id,content,author_id) values(?,?,?)"
	_, err = tx.Exec(sqlStr, answer.AnswerId, answer.Content, answer.AuthorId)
	if err != nil {
		logger.LogError("insert into answer failed:%v", err)
		_ = tx.Rollback()
		return
	}
	//  维护问题和关系列表
	sqlStr = "insert into question_answer_rel(question_id,answer_id) values(?,?)"
	_, err = tx.Exec(sqlStr, questionId, answer.AnswerId)

	if err != nil {
		logger.LogError("insert into question_answer_rel failed:%v", err)
		_ = tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}

// 更新count

func UpdateAnswerLike(like *model.Like, isInCreasement bool) (err error) {
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
		sqlStr = "update answer set voteup_count  =  voteup_count +1  where  answer_id =?"
	} else {
		sqlStr = "update answer set voteup_count  = voteup_count -1  where  answer_id =?"
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
