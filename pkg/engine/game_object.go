package engine

import "github.com/yoowhi/yarg/pkg/h"

type GameObject interface {
	GetPosition() h.Vector
	SetPosition(value h.Vector)
}
