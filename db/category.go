package db

import (
	"database/sql"
	"questions/model"
)

func GetCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select category_id,category_name from category"
	err = db.Select(categoryList, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	return
}
