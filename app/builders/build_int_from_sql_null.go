package builders

import "database/sql"

func GetIntValueFromSqlNull(prop sql.NullInt32) *int32 {
	if prop.Valid {
		return &prop.Int32
	}

	return nil
}
