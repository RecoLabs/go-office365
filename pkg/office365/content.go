package office365

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// ContentService .
type ContentService service

// List returns a list of content available for retrieval.
//
// Microsoft API Reference: https://docs.microsoft.com/en-us/office/office-365-management-api/office-365-management-activity-api-reference#list-available-content
//
// This operation lists the content currently available for retrieval for the specified content type.
// The content is an aggregation of actions and events harvested from multiple servers across multiple datacenters.
// The content will be listed in the order in which the aggregations become available, but the events and actions within
// the aggregations are not guaranteed to be sequential. An error is returned if the subscription status is disabled.
func (s *ContentService) List(ctx context.Context, ct *ContentType, startTime time.Time, endTime time.Time) ([]Content, error) {
	params := NewQueryParams()
	params.AddPubIdentifier(s.client.pubIdentifier)
	if err := params.AddContentType(ct); err != nil {
		return nil, err
	}
	if err := params.AddStartEndTime(startTime, endTime); err != nil {
		return nil, err
	}

	out := []Content{}
	var err error
	for {
		req, err := s.client.newRequest("GET", "subscriptions/content", params.Values, nil)
		if err != nil {
			return nil, err
		}

		var sub []Content
		response, err := s.client.do(ctx, req, &sub)
		if err != nil {
			return nil, err
		}
		out = append(out, sub...)

		nextPageURIStr := response.Header.Get("NextPageUri")
		if nextPageURIStr == "" {
			break
		}
		nextPageURI, err := url.ParseRequestURI(nextPageURIStr)
		if err != nil {
			return nil, err
		}
		nextPage := nextPageURI.Query().Get("nextpage")
		if nextPage == "" {
			return nil, fmt.Errorf("nextpage is not present as queryParam of NextPageUri header")
		}
		params.Set("nextpage", nextPage)
	}
	return out, err
}