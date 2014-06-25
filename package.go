package tarp

import (
	"io"
)

// Package represents a package (collection of items).
type Package interface {
	Name() string
	Version() string
	Walk(walkFn WalkFunc) error
	Close() error
}

type WalkFunc func(item Item, err error) error

// Item is a linked list of package items.
type Item interface {
	Name() string
	Mode() int64
	Uname() string
	Gname() string
	Size() int64
	Type() ItemType
	Linkname() string
	io.Reader
}

// ItemType determines the type of item.
type ItemType byte

const (
	ItemReg     ItemType = 1
	ItemLink    ItemType = 2
	ItemSymlink ItemType = 3
	ItemDir     ItemType = 4
)

type basePackage struct {
	name    string
	version string
}

func (p basePackage) Name() string {
	return p.name
}

func (p basePackage) Version() string {
	return p.version
}

type baseItem struct {
	name     string
	mode     int64
	uname    string
	gname    string
	size     int64
	itemType ItemType
	linkname string
}

func (i baseItem) Name() string {
	return i.name
}

func (i baseItem) Mode() int64 {
	return i.mode
}

func (i baseItem) Uname() string {
	return i.uname
}

func (i baseItem) Gname() string {
	return i.gname
}

func (i baseItem) Size() int64 {
	return i.size
}

func (i baseItem) Type() ItemType {
	return i.itemType
}

func (i baseItem) Linkname() string {
	return i.linkname
}

func (t ItemType) String() string {
	switch t {
	case ItemReg:
		return "file"
	case ItemLink:
		return "link"
	case ItemSymlink:
		return "symlink"
	case ItemDir:
		return "dir"
	}
	return "unknown"
}
