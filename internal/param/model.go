package param

// Param for describing request's parameter
type Param struct {
	Columns       string
	Order         string
	Limit         uint64
	Offset        uint64
	Search        string
	PreConditions string
	UserID        uint64
}
