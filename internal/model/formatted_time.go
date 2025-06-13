package model

import (
	"fmt"
	"time"
)

type FormattedTime struct {
	time.Time
}

func (ft FormattedTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", ft.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}
