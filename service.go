package browse

import (
	"github.com/go-resty/resty/v2"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	marketplace string
	appToken    string
	timeout     time.Duration
	limit       int
	isSandbox   bool
}

// NewService creates new Service.
// It requires already created Application Token.
// You can use this repo to create token https://github.com/hotafrika/ebay-common
// More info here https://developer.ebay.com/api-docs/static/oauth-client-credentials-grant.html
func NewService(appToken string) *Service {
	s := Service{appToken: appToken}
	s.WithTimeout(10 * time.Second)
	s.WithMarketplace(MarketplaceUSA)
	s.WithLimit(50)
	return &s
}

// WithTimeout sets timeout for HTTP client.
// The same timeout will be for all child requests.
func (s *Service) WithTimeout(timeout time.Duration) *Service {
	s.timeout = timeout
	return s
}

// WithMarketplace sets marketplace header for HTTP client.
// The same header will be for all child requests.
func (s *Service) WithMarketplace(marketplace EbayMarketplace) *Service {
	s.marketplace = string(marketplace)
	return s
}

// WithApplicationToken sets application token (Authorization) header for HTTP client.
// The same header will be for all child requests.
func (s *Service) WithApplicationToken(token string) *Service {
	s.appToken = token
	return s
}

// WithSandbox tells to Service to create Sandbox requests.
func (s *Service) WithSandbox(isSandbox bool) *Service {
	s.isSandbox = isSandbox
	return s
}

// WithLimit tells to Service to create requests with specified limit per page.
// Default: 50
func (s *Service) WithLimit(limit int) *Service {
	s.limit = limit
	return s
}

// NewItemSummarySearchRequest creates new ItemSummarySearchRequest
func (s *Service) NewItemSummarySearchRequest() *ItemSummarySearchRequest {
	r := &ItemSummarySearchRequest{}
	r.url = URLProductionItemSummarySearch
	if s.isSandbox {
		r.url = URLSandboxItemSummarySearch
	}
	r.client = resty.New().
		SetTimeout(s.timeout).
		SetHeader("X-EBAY-C-MARKETPLACE-ID", s.marketplace).
		SetHeader("Authorization", strings.Join([]string{"Bearer", s.appToken}, " "))
	r.Limit = strconv.Itoa(s.limit)
	r.FilterMap = make(map[string]Filterer)
	return r
}
