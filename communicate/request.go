package communicate

const (
	MinPageSize = 1
)

/**
 * @Author: feiliang.wang
 * @Description: 通用请求定义
 * @File:  request
 * @Version: 1.0.0
 * @Date: 2020/8/1 11:21 上午
 */

//分页请求
type PageRequest struct {
	SortAsc    uint32 `json:"sortAsc,omitempty"`
	SortField  string `json:"sortField,omitempty"`
	PageNumber int    `json:"pageNumber,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
}

func (r *PageRequest) Verification(count int) (maxPageNum int) {
	if r.PageSize < MinPageSize {
		r.PageSize = MinPageSize
	}
	maxPageNum = (count-1)/r.PageSize + 1
	minPageNum := 1
	if r.PageNumber < minPageNum {
		r.PageNumber = minPageNum
	}
	if r.PageNumber > maxPageNum {
		r.PageNumber = minPageNum
	}
	return maxPageNum
}
