package etcdtree

import (
	"strings"
)

type Node struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	IsKey   bool   `json:"isKey,omitempty"`
	SubTree Tree   `json:"subTree,omitempty"`
}

type Tree map[string]*Node

type Root struct {
	separator string
	Tree      Tree `json:"tree"`
}

func New(separator string) *Root {
	return &Root{
		separator: separator,
		Tree:      Tree{},
	}
}

func (r *Root) Add(key string) {
	r.Tree.add(key, r.separator, 0)
}

func (t Tree) add(key, separator string, depth int) {
	// key can not be blank
	if key == "" {
		return
	}

	// if there is no separator then we are gonna just to a flat map
	if separator == "" {
		t[key] = &Node{
			Name:    key,
			Path:    key,
			IsKey:   true,
			SubTree: nil,
		}
		return
	}

	parts := strings.Split(key, separator)
	sub := parts[depth]
	path := strings.Join(parts[0:depth+1], separator)

	if sub == "" {
		sub = separator
	}

	if path == "" {
		path = separator
	}

	if t[sub] == nil {
		t[sub] = &Node{
			Name: sub,
			Path: path,
		}
	}

	if key == path {
		t[sub].IsKey = true
		return
	}

	if t[sub].SubTree == nil {
		t[sub].SubTree = Tree{}
	}

	depth++
	t[sub].SubTree.add(key, separator, depth)
}
