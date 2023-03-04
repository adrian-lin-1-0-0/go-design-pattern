package cor

import (
	"testing"
)

func TestMiddleware(t *testing.T) {

	middleware := NewMiddleware()

	out := ""

	middleware.Use(func(msg string, next Next) {
		if msg == "1" {
			out = "1"
		} else {
			next(msg)
		}
	})

	middleware.Use(func(msg string, next Next) {
		if msg == "2" {
			out = "2"
		} else {
			next(msg)
		}
	})

	middleware.Use(func(msg string, next Next) {
		if msg == "3" {
			out = "3"
		} else {
			next(msg)
		}
	})

	middleware.Run("4")
	want := ""
	if out != want {
		t.Errorf("middleware.Run('4') got = %v, want %v", out, want)
	}
	middleware.Run("2")
	want = "2"
	if out != want {
		t.Errorf("middleware.Run('2') got = %v, want %v", out, want)
	}
	middleware.Run("2")
	if out != want {
		t.Errorf("middleware.Run('2') got = %v, want %v", out, want)
	}
}

func TestMiddlewareWithoutHandler(t *testing.T) {

	middleware := NewMiddleware()

	middleware.Run("4")

}
