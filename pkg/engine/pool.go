package engine

import "github.com/yoowhi/yarg/pkg/h"

type Pool[T GameObject] struct {
	list []T
}

func (p *Pool[T]) Add(actor T) {
	p.list = append(p.list, actor)
}

func (p *Pool[T]) Remove(actor *T) bool {
	for i, v := range p.list {
		if &v == actor {
			last := len(p.list) - 1
			p.list[i] = p.list[last]
			p.list = p.list[:last]
			return true
		}
	}
	return false
}

func (p *Pool[T]) Find(position h.Vector) *T {
	for _, v := range p.list {
		if v.GetPosition().Equals(position) {
			return &v
		}
	}
	return nil
}
