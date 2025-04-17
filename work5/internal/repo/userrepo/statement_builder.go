package userrepo

import (
	"github.com/Masterminds/squirrel"
)

var StatementBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) //nolint:gochecknoglobals
