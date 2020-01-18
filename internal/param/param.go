package param

// Param for describing request's parameter
type Param struct {
	Select       string
	Order        string
	Limit        uint64
	Offset       uint64
	Search       string
	PreCondition string
	UserID       uint64
}
