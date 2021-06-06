// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/go-chi/chi/v5"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
)

type Server interface {
	// PetGet
	PetGet(ctx context.Context) (*Pet, error)
}