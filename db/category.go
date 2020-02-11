package db

import (
	"github.com/jmoiron/sqlx"
	"logger"
	"questions/model"
)

func GetCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select category_id,category_name from category"
	err = db.Select(&categoryList, sqlStr)
	// if err == sql.ErrNoRows {
	// 	err = nil
	// 	return
	// }
	return
}

func GetCategory(categoryIdList []int64) (categoryMap map[int64]*model.Category, err error) {

	if len(categoryIdList) == 0 {
		return
	}
	sqlStr := "select category_id,category_name from category where category_id in (?)"
	var tempIdList []interface{}

	for _, id := range categoryIdList {
		tempIdList = append(tempIdList, id)
	}
	query, args, err := sqlx.In(sqlStr, tempIdList)

	if err != nil {
		logger.LogError("sql has a problem,%v", err)
		return
	}
	var categoryList []*model.Category
	err = db.Select(&categoryList, query, args...)
	if err != nil {
		logger.LogError("query categoryList failed:%v ", err)
		return
	}
	categoryMap = make(map[int64]*model.Category)
	for _, category := range categoryList {
		categoryMap[category.CategoryId] = category
	}
	return
}
