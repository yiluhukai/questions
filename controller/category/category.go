package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"questions/db"
	"questions/util"
)

func GetCategoryList(c *gin.Context) {
	categoryList, err := db.GetCategoryList()
	if err != nil {
		fmt.Printf("fetch categoryList failed,%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
	}
	util.ResponseSuccess(c, categoryList)
}
