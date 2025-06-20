package model

type PageReq struct {
	DateRange []string `json:"dateRange" dc:"时间范围"`
	PageNum   int      `json:"pageNum" dc:"第几页"`
	PageSize  int      `json:"pageSize" dc:"每页多少数据"`
	OrderBy   string   `json:"orderBy" dc:"排序"`
}
