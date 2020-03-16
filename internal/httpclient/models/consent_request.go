// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsentRequest ConsentRequest Contains information on an ongoing consent request.
// swagger:model consentRequest
type ConsentRequest struct {

	// ACR represents the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it
	// to express that, for example, a user authenticated using two factor authentication.
	Acr string `json:"acr,omitempty"`

	// Challenge is the identifier ("authorization challenge") of the consent authorization request. It is used to
	// identify the session.
	Challenge string `json:"challenge,omitempty"`

	// client
	Client *OAuth2Client `json:"client,omitempty"`

	// context
	Context JSONRawMessage `json:"context,omitempty"`

	// LoginChallenge is the login challenge this consent challenge belongs to. It can be used to associate
	// a login and consent request in the login & consent app.
	LoginChallenge string `json:"login_challenge,omitempty"`

	// LoginSessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag)
	// this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false)
	// this will be a new random value. This value is used as the "sid" parameter in the ID Token and in OIDC Front-/Back-
	// channel logout. It's value can generally be used to associate consecutive login requests by a certain user.
	LoginSessionID string `json:"login_session_id,omitempty"`

	// oidc context
	OidcContext *OpenIDConnectContext `json:"oidc_context,omitempty"`

	// RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which
	// initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but
	// might come in handy if you want to deal with additional request parameters.
	RequestURL string `json:"request_url,omitempty"`

	// requested access token audience
	RequestedAccessTokenAudience []string `json:"requested_access_token_audience,omitempty"`

	// requested scope
	RequestedScope []string `json:"requested_scope,omitempty"`

	// Skip, if true, implies that the client has requested the same scopes from the same user previously.
	// If true, you must not ask the user to grant the requested scopes. You must however either allow or deny the
	// consent request using the usual API call.
	Skip bool `json:"skip,omitempty"`

	// Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope
	// requested by the OAuth 2.0 client.
	Subject string `json:"subject,omitempty"`
}

// Validate validates this consent request
func (m *ConsentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidcContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentRequest) validateClient(formats strfmt.Registry) error {

	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) validateOidcContext(formats strfmt.Registry) error {

	if swag.IsZero(m.OidcContext) { // not required
		return nil
	}

	if m.OidcContext != nil {
		if err := m.OidcContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc_context")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentRequest) UnmarshalBinary(b []byte) error {
	var res ConsentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
