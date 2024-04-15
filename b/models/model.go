package models

import (
	a "test-gorm-gen/a/models"
)

type B struct {
	AID uint
	A   *a.A
	Six string
}
