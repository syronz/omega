package types

// MinCol is a same as model.gorm, we use our name if in futer customize it don't face problem
type MinCol struct {
	ID RowID `gorm:"primary_key" json:"id,omitempty" `
}
