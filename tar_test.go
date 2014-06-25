package tarp

import (
	"testing"
)

func TestTarInit(t *testing.T) {
	err := Tar.Init(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTarLoad(t *testing.T) {
	dsn := "test_0.1.0-1.tar"
	p, err := Tar.Load(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	err = p.Walk(func(item Item, err error) error {
		t.Log(item)
		return err
	})
	if err != nil {
		t.Fatalf("walk failed '%s'", err)
	}
	err = p.Walk(func(item Item, err error) error {
		t.Log(item)
		return err
	})
	if err != nil {
		t.Fatalf("second walk failed '%s'", err)
	}
}
