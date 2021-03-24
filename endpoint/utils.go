package endpoint

import (
	"fmt"
	"github.com/feiliangwang/audioAnalysis/communicate/errors"
)

/**
 * @Author: feiliang.wang
 * @Description: 帮助工具
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 2021/3/23 23:32
 */

var (
	ErrorBadRequest = errors.NewParamError(fmt.Errorf("invalid request parameter"))
)
