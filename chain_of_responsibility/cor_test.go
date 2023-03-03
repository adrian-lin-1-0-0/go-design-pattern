package cor

import (
	"testing"
)

func TestMiddleware(t *testing.T) {

	middleware := NewMiddleware()

	middleware.Use(func(msg string, next Next) {
		if msg == "1" {
			print("1")
		} else {
			println("next")
			next(msg)
		}
	})

	middleware.Use(func(msg string, next Next) {
		if msg == "2" {
			print("2")
		} else {
			println("next")
			next(msg)
		}
	})

	middleware.Use(func(msg string, next Next) {
		if msg == "3" {
			print("3")
		} else {
			next(msg)
		}
	})

	middleware.Run("4")

}

func TestMiddlewareWithoutHandler(t *testing.T) {

	middleware := NewMiddleware()

	middleware.Run("4")

}
