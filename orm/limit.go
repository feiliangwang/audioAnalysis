package orm

import "fmt"

/**
 * @Author: feiliang.wang
 * @Description: 分页查询处理
 * @File:  limit
 * @Version: 1.0.0
 * @Date: 2020/8/5 1:50 下午
 */

func Limit(num, size int) string {
	return fmt.Sprintf(" LIMIT %d,%d ", (num-1)*size, size)
}
