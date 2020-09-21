package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// RowID is used for defining the type of id
type RowID uint64

// ToUint64 convert the value to uint, useful for gorm
func (r *RowID) ToUint64() uint64 {
	return uint64(*r)
}

// ToString convert the value to string
func (r *RowID) ToString() string {
	return fmt.Sprint(*r)
}

// RowIDPointer return a pointer to the RowID
func RowIDPointer(num uint64) *RowID {
	rowID := RowID(num)
	return &rowID
}

// StrToRowID convert string number to RowID
func StrToRowID(strNum string) (rowID RowID, err error) {
	tmpID, err := strconv.ParseUint(strNum, 10, 64)
	rowID = RowID(tmpID)
	return
}

// Value is used to save to the database
func (r RowID) Value() (driver.Value, error) {
	// TODO: instead of sprintf use number conversion
	return fmt.Sprint(r), nil
}

/*
// TODO: I don't know what is the usage of this
// https://github.com/jinzhu/gorm/issues/47

func (r *RowID) Scan(value interface{}) error {
	value, _ = value.(uint64)
	*r = RowID(value)
	return nil
}
*/

// func (r *RowID) Scan(value interface{}) error {
// 	result, _ := value.(RowID)
// 	*r = RowID(result)
// 	return nil
// }
