package browse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type ItemSummarySearchRequest struct {
	Q           string
	Limit       string
	Offset      string
	CategoryIDs []string
	AutoCorrect []string
	Sort        []string
	FilterMap   map[string]Filterer
	BasicRequest
}

// WithQuery adds query parameter to ItemSummarySearchRequest
func (r *ItemSummarySearchRequest) WithQuery(query string) *ItemSummarySearchRequest {
	r.Q = query
	return r
}

// WithLimit sets limit of items per page to ItemSummarySearchRequest
// Default : 50
func (r *ItemSummarySearchRequest) WithLimit(limit int) *ItemSummarySearchRequest {
	r.Limit = strconv.Itoa(limit)
	return r
}

// WithOffset adds offset to skip during the call. Field of ItemSummarySearchRequest
// Default : 0
func (r *ItemSummarySearchRequest) WithOffset(offset int) *ItemSummarySearchRequest {
	r.Offset = strconv.Itoa(offset)
	return r
}

// WithCategories adds categories for search to ItemSummarySearchRequest
// Leave empty to omit
//func (r *ItemSummarySearchRequest) WithCategories(categories ...string) *ItemSummarySearchRequest {
//	r.CategoryIDs = categories
//	return r
//}
func (r *ItemSummarySearchRequest) WithCategoryID(categoryID string) *ItemSummarySearchRequest {
	r.CategoryIDs = []string{categoryID}
	return r
}

// WithAutoCorrect adds autocorrecting option to ItemSummarySearchRequest
// Leave empty to omit
func (r *ItemSummarySearchRequest) WithAutoCorrect(values ...string) *ItemSummarySearchRequest {
	r.AutoCorrect = values
	return r
}

// WithSort adds sorting order to ItemSummarySearchRequest
// Default: BestMatch
func (r *ItemSummarySearchRequest) WithSort(sortBy ...string) *ItemSummarySearchRequest {
	r.Sort = sortBy
	return r
}

// WithFilterDeliveryCountry adds filter by deliveryCountry to ItemSummarySearchRequest
// Parameter for this function is Two-character ISO 3166 Country Code ("US", "GB" etc.)
func (r *ItemSummarySearchRequest) WithFilterDeliveryCountry(country string) *ItemSummarySearchRequest {
	title := string(FilterDeliveryCountry)
	f := OneValueFilter{
		Title: title,
		Value: country,
	}
	r.checkFilterMap()
	r.FilterMap[title] = f
	return r
}

// WithFilterDeliveryPostalCode adds filter by deliveryPostalCode to ItemSummarySearchRequest
// Must be used with deliveryCountry parameter.
func (r *ItemSummarySearchRequest) WithFilterDeliveryPostalCode(postalCode string) *ItemSummarySearchRequest {
	title := string(FilterDeliveryPostalCode)
	f := OneValueFilter{
		Title: title,
		Value: postalCode,
	}
	r.checkFilterMap()
	r.FilterMap[title] = f
	return r
}

func (r *ItemSummarySearchRequest) checkFilterMap() {
	if len(r.FilterMap) == 0 {
		r.FilterMap = make(map[string]Filterer)
	}
}

func (r *ItemSummarySearchRequest) Execute() (ItemSummarySearchResponse, error) {
	values := r.getParams()
	tRes, err := r.client.R().SetQueryParamsFromValues(values).Get(r.url)
	if err != nil {
		return ItemSummarySearchResponse{}, fmt.Errorf("error during http request %w", err)
	}
	if tRes.StatusCode() != 200 {
		return ItemSummarySearchResponse{}, fmt.Errorf("response code is %d, not 200", tRes.StatusCode())
	}
	//fmt.Println(tRes.String())
	res := ItemSummarySearchResponse{}
	err = json.Unmarshal(tRes.Body(), &res)
	if err != nil {
		return ItemSummarySearchResponse{}, fmt.Errorf("error during unmarshalling response %w", err)
	}
	return res, nil
}

func (r *ItemSummarySearchRequest) getParams() url.Values {
	values := url.Values{}
	if r.Q != "" {
		values.Set("q", r.Q)
	}
	if r.Limit != "" {
		values.Set("limit", r.Limit)
	}
	if r.Offset != "" {
		values.Set("offset", r.Offset)
	}
	if len(r.CategoryIDs) > 0 {
		value := strings.Join(r.CategoryIDs, ",")
		values.Set("category_ids", value)
	}
	if len(r.AutoCorrect) > 0 {
		value := strings.Join(r.AutoCorrect, ",")
		values.Set("auto_correct", value)
	}
	if len(r.Sort) > 0 {
		value := strings.Join(r.Sort, ",")
		values.Set("sort", value)
	}
	if len(r.FilterMap) > 0 {
		var valuesList []string
		for _, item := range r.FilterMap {
			valuesList = append(valuesList, item.ToQueryParams())
		}
		value := strings.Join(valuesList, ",")
		values.Set("filter", value)
	}
	return values
}

func (r *ItemSummarySearchRequest) ShowParams() url.Values {
	return r.getParams()
}
