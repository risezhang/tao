{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package repo

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Repo interface {
	Query(ctx context.Context, q *QueryParams) ([]*{{.Name|title}}, error)
	Get(ctx context.Context, q *GetParams) (*{{.Name|title}}, error)
	Update(ctx context.Context, v Values, id int64) error
	Create(ctx context.Context, req *{{.Name|title}}) (int64, error)
	Delete(ctx context.Context, id int64) error
}