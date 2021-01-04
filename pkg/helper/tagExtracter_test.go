package helper

import (
	"lab/internal/types"
	"reflect"
	"testing"
	"time"
)

// Company model
type Company struct {
	Name          string     `gorm:"not null" json:"name,omitempty"`
	LegalName     string     `gorm:"not null;unique" json:"legal_name,omitempty"`
	Key           string     `gorm:"type:text" json:"key,omitempty"`
	ServerAddress string     `json:"server_address,omitempty"`
	Expiration    *time.Time `json:"expiration,omitempty"`
	License       string     `gorm:"unique" json:"license,omitempty"`
	Plan          string     `json:"plan,omitempty"`
	Detail        string     `json:"detail,omitempty"`
	Phone         string     `gorm:"not null" json:"phone,omitempty"`
	Email         string     `gorm:"not null" json:"email,omitempty"`
	Website       string     `gorm:"not null" json:"website,omitempty"`
	Type          string     `gorm:"not null" json:"type,omitempty"`
	Code          string     `gorm:"not null" json:"code,omitempty"`
	Logo          string     `json:"logo"`
	Banner        string     `json:"banner"`
	Footer        string     `json:"footer" table:"-"`
}

// Result model
type Result struct {
	CreatedBy   types.RowID `json:"created_by,omitempty"`
	UpdatedBy   types.RowID `json:"updated_by,omitempty"`
	PatientID   types.RowID `json:"patient_id,omitempty"`
	Gender      types.Enum  `json:"gender,omitempty"`
	Doctor      string      `sql:"-" json:"doctor,omitempty" table:"doctors.name as doctor"`
	Medications string      `json:"medications,omitempty"`
	Total       float64     `json:"total,omitempty"`
	Discount    float64     `json:"discount,omitempty"`
	Notes       string      `json:"notes"`
	Patient     string      `sql:"-" json:"patient,omitempty" table:"bas_accounts.name as patient"`
	DOB         *time.Time  `sql:"-" json:"dob,omitempty" table:"sam_patients.dob as dob"`
	Company     Company     `sql:"-" json:"company,omitempty" table:"-"`
}

func TestTagExtractor(t *testing.T) {
	t.Log("this is a test for tag-extractor")

	r := TagExtracter(reflect.TypeOf(Result{}), "sam_results")
	t.Log(r)

}
