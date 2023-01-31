package hashy_test

import (
	"strings"
	"testing"

	"github.com/rnemeth90/hashy"
)

const (
	hashMe = "test"
)

func TestGenerateSHA224(t *testing.T) {
	want := "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809"
	got := hashy.GenerateSHA224(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestGenerateSHA256(t *testing.T) {
	want := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	got := hashy.GenerateSHA256(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestGenerateSHA384(t *testing.T) {
	want := "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9"
	got := hashy.GenerateSHA384(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestGenerateSHA512(t *testing.T) {
	want := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	got := hashy.GenerateSHA512(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestGenerateSHA512224(t *testing.T) {
	want := "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c"
	got := hashy.GenerateSHA512224(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestGenerateSHA512256(t *testing.T) {
	want := "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f"
	got := hashy.GenerateSHA512256(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGenerateMD5(t *testing.T) {
	want := "098f6bcd4621d373cade4e832627b4f6"
	got := hashy.GenerateMD5(hashMe)

	if strings.TrimSuffix(got, "\n") != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
