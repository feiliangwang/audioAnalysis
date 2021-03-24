package errors

import (
	"errors"
	"testing"
)

/**
 * @Author: feiliang.wang
 * @Description: //TODO
 * @File:  errors_test
 * @Version: 1.0.0
 * @Date: 2020/8/7 7:26 下午
 */

func TestNewOperateError(t *testing.T) {
	err := NewOperateError(errors.New("123"))
	t.Log(err)
}

func TestNewParamError(t *testing.T) {
	err := NewParamError(errors.New("123"))
	t.Log(err)
}
