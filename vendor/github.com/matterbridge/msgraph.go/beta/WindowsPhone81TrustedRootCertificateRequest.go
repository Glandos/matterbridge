// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsPhone81TrustedRootCertificateRequestBuilder is request builder for WindowsPhone81TrustedRootCertificate
type WindowsPhone81TrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsPhone81TrustedRootCertificateRequest
func (b *WindowsPhone81TrustedRootCertificateRequestBuilder) Request() *WindowsPhone81TrustedRootCertificateRequest {
	return &WindowsPhone81TrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsPhone81TrustedRootCertificateRequest is request for WindowsPhone81TrustedRootCertificate
type WindowsPhone81TrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Get(ctx context.Context) (resObj *WindowsPhone81TrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Update(ctx context.Context, reqObj *WindowsPhone81TrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}