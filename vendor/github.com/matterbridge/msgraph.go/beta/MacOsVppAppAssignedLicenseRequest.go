// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOsVppAppAssignedLicenseRequestBuilder is request builder for MacOsVppAppAssignedLicense
type MacOsVppAppAssignedLicenseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOsVppAppAssignedLicenseRequest
func (b *MacOsVppAppAssignedLicenseRequestBuilder) Request() *MacOsVppAppAssignedLicenseRequest {
	return &MacOsVppAppAssignedLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOsVppAppAssignedLicenseRequest is request for MacOsVppAppAssignedLicense
type MacOsVppAppAssignedLicenseRequest struct{ BaseRequest }

// Get performs GET request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Get(ctx context.Context) (resObj *MacOsVppAppAssignedLicense, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Update(ctx context.Context, reqObj *MacOsVppAppAssignedLicense) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}