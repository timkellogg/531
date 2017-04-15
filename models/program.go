package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Program - model
type Program struct {
	gorm.Model
	Variation      string
	StartingWeight uint
	UserID         uint
}

// AcceptableVariations - the possible 531 workout types. Currently, supports just "BBB"
var AcceptableVariations = [1]string{"BBB"}

// BeforeSave - validations
func (p *Program) BeforeSave() (err []error) {
	messages := make([]error, 0)

	// Presence Validations
	includesVariation := false
	for _, variation := range AcceptableVariations {
		if p.Variation == variation {
			includesVariation = true
		}
	}

	if includesVariation == false {
		messages = append(messages, errors.New("Program must include an acceptable variation"))
	}

	return messages
}
