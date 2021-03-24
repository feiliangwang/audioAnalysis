package orm

import (
	"fmt"
)

/**
 * @Author: feiliang.wang
 * @Description: 排序处理
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2020/8/5 1:48 下午
 */

func Order(sortAsc int, sortField string) string {
	if sortField != "" {
		sortKind := "DESC"
		if sortAsc == 1 {
			sortKind = "ASC"
		}
		return fmt.Sprintf(" order by %s %s ", SnakeField(sortField), sortKind)
	} else {
		return ""
	}
}
