package stream

import (
	"github.com/List412/streams_supervisor/pkg/domain/state"
	"github.com/jinzhu/gorm"
)

type Stream struct {
	gorm.Model
	state state.State
}
