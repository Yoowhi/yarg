package engine

import "github.com/yoowhi/yarg/pkg/h"

type Pool[T GameObject] struct {
	List []T
}

func (p *Pool[T]) Add(actor T) {
	p.List = append(p.List, actor)
}

func (p *Pool[T]) Remove(actor *T) bool {
	for i, v := range p.List {
		if &v == actor {
			last := len(p.List) - 1
			p.List[i] = p.List[last]
			p.List = p.List[:last]
			return true
		}
	}
	return false
}

func (p *Pool[T]) Find(position h.Vector) *T {
	for _, v := range p.List {
		if v.GetPosition().Equals(position) {
			return &v
		}
	}
	return nil
}
