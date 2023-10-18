package types

import (
	"github.com/mattismoel/icalendar/types"
)

func New(productName, cooperation, filepath string) *types.ICalendar {
	return &types.ICalendar{
		ProductName:  productName,
		Coorperation: cooperation,
		FilePath:     filepath,
		Events: make([]*types.ICalEvent, 0),
	}
}



