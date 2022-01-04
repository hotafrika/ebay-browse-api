package browse

import (
	"github.com/go-resty/resty/v2"
	"strings"
)

type BasicRequest struct {
	url    string
	client *resty.Client
}

type Filterer interface {
	ToQueryParams() string
}

type OneValueFilter struct {
	Title string
	Value string
}

func (ovf OneValueFilter) ToQueryParams() string {
	builder := strings.Builder{}
	builder.WriteString(ovf.Title)
	builder.WriteString(":")
	builder.WriteString(ovf.Value)
	return builder.String()
}

type TwoValueFilter struct {
	Title  string
	Value1 string
	Value2 string
}

type ManyValueFilter struct {
	Title  string
	Values []string
}
