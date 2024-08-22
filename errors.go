// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "github.com/OneLoop-HQ/oneloop-go/core"
)

// Bad Request
type BadRequestError struct {
	*core.APIError
	Body *ErrorResponse
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// Forbidden
type ForbiddenError struct {
	*core.APIError
	Body *ErrorResponse
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

// Unauthorized
type UnauthorizedError struct {
	*core.APIError
	Body *ErrorResponse
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}
