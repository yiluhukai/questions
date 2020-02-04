package filter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	//加载敏感词库
	err := Init("../data/filter.data.txt")
	if err != nil {
		t.Errorf("load filter data failed, err:%v", err)
		return
	}
	data := `傻逼玩意儿`
	result, hit := Replace(data, "***")
	fmt.Printf("hit:%#v, str:%v\n", hit, result)
}
