package param

// Param for describing request's parameter
type Param struct {
	Pagination
	Search       string
	PreCondition string
	UserID       uint64
}

// Pagination is a struct, contains the fields which affected the front-end pagination
type Pagination struct {
	Select string
	Order  string
	Limit  uint64
	Offset uint64
}
