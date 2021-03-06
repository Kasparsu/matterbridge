// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ScopedRoleMembershipRequestBuilder is request builder for ScopedRoleMembership
type ScopedRoleMembershipRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScopedRoleMembershipRequest
func (b *ScopedRoleMembershipRequestBuilder) Request() *ScopedRoleMembershipRequest {
	return &ScopedRoleMembershipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScopedRoleMembershipRequest is request for ScopedRoleMembership
type ScopedRoleMembershipRequest struct{ BaseRequest }

// Get performs GET request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Get(ctx context.Context) (resObj *ScopedRoleMembership, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Update(ctx context.Context, reqObj *ScopedRoleMembership) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
