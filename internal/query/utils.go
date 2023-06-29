package query

type PaginationQuery struct {
	Size int `json:"size" default:"20"`
	Page int `json:"page" default:"1"`
}

type Pagination struct {
	TotalPages  int   `json:"totalPages"`
	TotalItems  int64 `json:"totalItems"`
	NextPage    int   `json:"nextPage"`
	CurrentPage int   `json:"currentPage"`
} //	@Name	Pagination

func Paginate(totalItems int64, totalData, page, size int) Pagination {
	if size == 0 {
		size = 20
	}
	nextPage := page + 1
	pages := float64(totalItems / int64(size))

	if nextPage > int(pages) {
		nextPage = int(pages)
	}

	if page-1 == 0 {
		page = 1
	}

	return Pagination{
		TotalPages:  int(pages),
		TotalItems:  int64(totalData),
		NextPage:    nextPage,
		CurrentPage: page,
	}
}
