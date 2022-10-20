package custom

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func (v *CustomValidator) Uniquedb(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), ".")

	var count int8
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s = ?", params[0], params[1])
	v.db.Raw(query, fl.Field().String()).Scan(&count)

	return count == 0
}
