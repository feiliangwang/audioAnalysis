package errors

/**
 * @Author: feiliang.wang
 * @Description: 错误定义
 * @File:  errors
 * @Version: 1.0.0
 * @Date: 2020/8/7 4:12 下午
 */

const ErrorParamCode = 1001
const ErrorOperateCode = 1002

type ParamError struct {
	err error
}

func NewParamError(err error) ParamError {
	return ParamError{err}
}

func (e ParamError) Error() string {
	return e.err.Error()
}

func (e ParamError) Code() int32 {
	return ErrorParamCode
}

type OperateError struct {
	err error
}

func NewOperateError(err error) OperateError {
	return OperateError{err}
}

func (e OperateError) Error() string {
	return e.err.Error()
}

func (e OperateError) Code() int32 {
	return ErrorOperateCode
}
