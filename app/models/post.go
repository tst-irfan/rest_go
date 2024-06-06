
package models
import (
	"github.com/jinzhu/gorm"
	"rest_go/db"
	"rest_go/app/helpers"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`

}

var PostQuery = db.QueryHelper[Post]{}
var PostValidation = helpers.ValidationHelper{
	RequiredFields:          []string{},
	ShouldGreaterThanFields: []helpers.Field{},
	ShouldLessThanFields:    []helpers.Field{},
}

func (m *Post) BeforeSave() error {
	_, err := PostValidation.Validate(m)
	if err != nil {
		return err
	}
	return nil
}
