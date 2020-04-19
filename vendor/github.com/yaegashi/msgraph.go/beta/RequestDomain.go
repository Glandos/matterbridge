// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DomainRequestBuilder is request builder for Domain
type DomainRequestBuilder struct{ BaseRequestBuilder }

// Request returns DomainRequest
func (b *DomainRequestBuilder) Request() *DomainRequest {
	return &DomainRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DomainRequest is request for Domain
type DomainRequest struct{ BaseRequest }

// Get performs GET request for Domain
func (r *DomainRequest) Get(ctx context.Context) (resObj *Domain, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Domain
func (r *DomainRequest) Update(ctx context.Context, reqObj *Domain) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Domain
func (r *DomainRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DomainDNSRecordRequestBuilder is request builder for DomainDNSRecord
type DomainDNSRecordRequestBuilder struct{ BaseRequestBuilder }

// Request returns DomainDNSRecordRequest
func (b *DomainDNSRecordRequestBuilder) Request() *DomainDNSRecordRequest {
	return &DomainDNSRecordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DomainDNSRecordRequest is request for DomainDNSRecord
type DomainDNSRecordRequest struct{ BaseRequest }

// Get performs GET request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Get(ctx context.Context) (resObj *DomainDNSRecord, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Update(ctx context.Context, reqObj *DomainDNSRecord) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DomainSecurityProfileRequestBuilder is request builder for DomainSecurityProfile
type DomainSecurityProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns DomainSecurityProfileRequest
func (b *DomainSecurityProfileRequestBuilder) Request() *DomainSecurityProfileRequest {
	return &DomainSecurityProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DomainSecurityProfileRequest is request for DomainSecurityProfile
type DomainSecurityProfileRequest struct{ BaseRequest }

// Get performs GET request for DomainSecurityProfile
func (r *DomainSecurityProfileRequest) Get(ctx context.Context) (resObj *DomainSecurityProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DomainSecurityProfile
func (r *DomainSecurityProfileRequest) Update(ctx context.Context, reqObj *DomainSecurityProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DomainSecurityProfile
func (r *DomainSecurityProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type DomainForceDeleteRequestBuilder struct{ BaseRequestBuilder }

// ForceDelete action undocumented
func (b *DomainRequestBuilder) ForceDelete(reqObj *DomainForceDeleteRequestParameter) *DomainForceDeleteRequestBuilder {
	bb := &DomainForceDeleteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/forceDelete"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DomainForceDeleteRequest struct{ BaseRequest }

//
func (b *DomainForceDeleteRequestBuilder) Request() *DomainForceDeleteRequest {
	return &DomainForceDeleteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DomainForceDeleteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type DomainVerifyRequestBuilder struct{ BaseRequestBuilder }

// Verify action undocumented
func (b *DomainRequestBuilder) Verify(reqObj *DomainVerifyRequestParameter) *DomainVerifyRequestBuilder {
	bb := &DomainVerifyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/verify"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DomainVerifyRequest struct{ BaseRequest }

//
func (b *DomainVerifyRequestBuilder) Request() *DomainVerifyRequest {
	return &DomainVerifyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DomainVerifyRequest) Post(ctx context.Context) (resObj *Domain, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}