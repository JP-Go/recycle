package database

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func fromTimestampzToTime(t pgtype.Timestamptz) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func fromTimeToTimestampz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{
			Time:  time.Time{},
			Valid: false,
		}
	}
	return pgtype.Timestamptz{
		Time:  *t,
		Valid: true,
	}
}
