package communicate

/**
 * @Author: feiliang.wang
 * @Description: 带错误码错误定义
 * @File:  error
 * @Version: 1.0.0
 * @Date: 2020/8/7 4:27 下午
 */

type CodeError interface {
	error
	Code() int32
}
