package browse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type GetItemByLegacyIdRequest struct {
	LegacyItemID       string
	LegacyVariationID  string
	LegacyVariationSKU string
	FieldGroups        []string
	BasicRequest
}

// WithItemID adds legacyItemID parameter to GetItemByLegacyIdRequest
func (r *GetItemByLegacyIdRequest) WithItemID(legacyItemID string) *GetItemByLegacyIdRequest {
	r.LegacyItemID = legacyItemID
	return r
}

// WithLegacyVariationID adds legacyVariationID parameter to GetItemByLegacyIdRequest
func (r *GetItemByLegacyIdRequest) WithLegacyVariationID(legacyVariationID string) *GetItemByLegacyIdRequest {
	r.LegacyItemID = legacyVariationID
	return r
}

// WithLegacyVariationSKY adds legacyVariationSKY parameter to GetItemByLegacyIdRequest
func (r *GetItemByLegacyIdRequest) WithLegacyVariationSKY(legacyVariationSKY string) *GetItemByLegacyIdRequest {
	r.LegacyVariationSKU = legacyVariationSKY
	return r
}

// WithFieldGroups adds fieldsGroups parameter to GetItemByLegacyIdRequest
func (r *GetItemByLegacyIdRequest) WithFieldGroups(fieldGroups ...FieldGroupsOptions) *GetItemByLegacyIdRequest {
	tfg := make([]string, 0)
	tfgm := make(map[string]struct{})
	for _, fg := range fieldGroups {
		tfgm[string(fg)] = struct{}{}
	}
	for k := range tfgm {
		tfg = append(tfg, k)
	}
	r.FieldGroups = tfg
	return r
}

func (r *GetItemByLegacyIdRequest) Execute() (GetItemByLegacyIdResponse, error) {
	values := r.getParams()
	tRes, err := r.client.R().SetQueryParamsFromValues(values).Get(r.url)
	if err != nil {
		return GetItemByLegacyIdResponse{}, fmt.Errorf("error during http request %w", err)
	}
	if tRes.StatusCode() != 200 {
		return GetItemByLegacyIdResponse{}, fmt.Errorf("response code is %d, not 200", tRes.StatusCode())
	}
	//fmt.Println(tRes.String())
	res := GetItemByLegacyIdResponse{}
	err = json.Unmarshal(tRes.Body(), &res)
	if err != nil {
		return GetItemByLegacyIdResponse{}, fmt.Errorf("error during unmarshalling response %w", err)
	}
	return res, nil
}

func (r *GetItemByLegacyIdRequest) getParams() url.Values {
	values := url.Values{}
	if r.LegacyItemID != "" {
		values.Set("legacy_item_id", r.LegacyItemID)
	}
	if r.LegacyVariationID != "" {
		values.Set("legacy_variation_id", r.LegacyVariationID)
	}
	if r.LegacyVariationSKU != "" {
		values.Set("legacy_variation_sku", r.LegacyVariationSKU)
	}
	if len(r.FieldGroups) > 0 {
		value := strings.Join(r.FieldGroups, ",")
		values.Set("fieldgroups", value)
	}
	return values
}

func (r *GetItemByLegacyIdRequest) ShowParams() url.Values {
	return r.getParams()
}
