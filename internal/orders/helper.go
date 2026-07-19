package implorder

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func requiredText(value string) pgtype.Text {
	return pgtype.Text{String: value, Valid: true}
}

func optionalText(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}
	return requiredText(*value)
}
