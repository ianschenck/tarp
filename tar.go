package tarp

import (
	"archive/tar"
	"io"
	"os"
	"path"
	"strings"
	"sync"
)

type TarProvider struct{}

var (
	Tar *TarProvider
)

func (t *TarProvider) Init(env map[string]string) error {
	return nil
}

func (t *TarProvider) Load(dsn string) (Package, error) {
	p := &tarPackage{}
	baseName := path.Base(dsn)
	fileBase := baseName[:strings.LastIndex(baseName, ".")]
	parts := strings.Split(fileBase, "_")
	if len(parts) == 0 {
		panic(fileBase)
	}
	p.name = parts[0]
	if len(parts) == 2 {
		p.version = parts[1]
	} else {
		p.version = "0.0.0"
	}
	if len(parts) > 2 {
		return nil, ErrBadName
	}
	if !SemverRE.MatchString(p.version) {
		return nil, ErrSemverViolation
	}

	f, err := os.Open(dsn)
	if err != nil {
		return nil, err
	}
	p.file = f
	return p, nil
}

func (t *tarPackage) Walk(walkFn WalkFunc) error {
	t.fileLock.Lock()
	defer t.fileLock.Unlock()
	_, err := t.file.Seek(0, 0)
	if err != nil {
		return err
	}
	tr := tar.NewReader(t.file)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		item := tarItem{}
		item.baseItem.name = hdr.Name
		item.baseItem.size = hdr.Size
		item.baseItem.mode = hdr.Mode
		item.baseItem.uname = hdr.Uname
		item.baseItem.gname = hdr.Gname
		switch hdr.Typeflag {
		case tar.TypeReg, tar.TypeRegA:
			item.baseItem.itemType = ItemReg
		case tar.TypeDir:
			item.baseItem.itemType = ItemDir
		case tar.TypeSymlink:
			item.baseItem.itemType = ItemSymlink
		case tar.TypeLink:
			item.baseItem.itemType = ItemLink
		default:
			item.baseItem.itemType = 0
		}
		item.baseItem.linkname = hdr.Linkname
		item.read = tr.Read
		err = walkFn(item, err)
		if err != nil {
			return err
		}
	}
}

func (t *tarPackage) Close() error {
	return t.file.Close()
}

type tarPackage struct {
	basePackage
	file     *os.File
	fileLock sync.Mutex
}

type tarItem struct {
	baseItem
	read func([]byte) (int, error)
}

func (i tarItem) Read(p []byte) (int, error) {
	return i.read(p)
}
