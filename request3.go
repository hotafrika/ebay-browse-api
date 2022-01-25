package browse

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type GetItemsByItemGroupRequest struct {
	ItemGroupID string
	BasicRequest
}

// WithItemGroupID adds itemGroupID parameter to GetItemsByItemGroupRequest
func (r *GetItemsByItemGroupRequest) WithItemGroupID(itemGroupID string) *GetItemsByItemGroupRequest {
	r.ItemGroupID = itemGroupID
	return r
}

func (r *GetItemsByItemGroupRequest) Execute() (GetItemsByItemGroupResponse, error) {
	values := r.getParams()
	tRes, err := r.client.R().SetQueryParamsFromValues(values).Get(r.url)
	if err != nil {
		return GetItemsByItemGroupResponse{}, fmt.Errorf("error during http request %w", err)
	}
	if tRes.StatusCode() != 200 {
		return GetItemsByItemGroupResponse{}, fmt.Errorf("response code is %d, not 200", tRes.StatusCode())
	}
	//fmt.Println(tRes.String())
	res := GetItemsByItemGroupResponse{}
	err = json.Unmarshal(tRes.Body(), &res)
	if err != nil {
		return GetItemsByItemGroupResponse{}, fmt.Errorf("error during unmarshalling response %w", err)
	}
	return res, nil
}

func (r *GetItemsByItemGroupRequest) getParams() url.Values {
	values := url.Values{}
	if r.ItemGroupID != "" {
		values.Set("item_group_id", r.ItemGroupID)
	}
	return values
}

func (r *GetItemsByItemGroupRequest) ShowParams() url.Values {
	return r.getParams()
}
