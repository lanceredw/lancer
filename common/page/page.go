package page

import (
	"gorm.io/gorm"
	"math"
)

type Request struct {
	PageSize  int64 `json:"page_size" form:"page_size"`   //page size
	PageIndex int64 `json:"page_index" form:"page_index"` //current page number
}

type Response struct {
	Count     int64       `json:"count"`      // total number of records
	TotalPage int64       `json:"total_page"` // total page
	PageSize  int64       `json:"page_size"`  // page size
	PageIndex int64       `json:"page_index"` // current page number
	List      interface{} `json:"list" `      //data list
}

// Paging Calculate paging data based on the current size count
func Paging(count int64, page Request) (pageResponse Response) {

	if page.PageSize <= 0 {
		pageResponse.PageSize = 10
	} else if page.PageSize > 10000 {
		pageResponse.PageSize = page.PageSize
	} else {
		pageResponse.PageSize = page.PageSize
	}
	if page.PageIndex <= 0 {
		pageResponse.PageIndex = 1
	} else {
		pageResponse.PageIndex = page.PageIndex
	}
	pageResponse.Count = count
	// count / pageSize
	p := float64(pageResponse.Count) / float64(pageResponse.PageSize)
	pageResponse.TotalPage = int64(math.Ceil(p))

	return pageResponse
}

// Paginate Pagination Query Scope
func Paginate(pageResponse Response) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int((pageResponse.PageIndex - 1) * pageResponse.PageSize)).Limit(int(pageResponse.PageSize))
	}
}
