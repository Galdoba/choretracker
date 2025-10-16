package dto

import (
	"encoding/json"
	"fmt"
)

const (
	Create = "create"
	Read   = "read"
	Update = "update"
	Delete = "delete"
)

type ToServiceRequest struct {
	Action   string `json:"action" validate:"required"`
	Identity ChoreIdentity
	Fields   ChoreContent
}

type ChoreIdentity struct {
	ID *int64 `json:"id" validate:"required"`
}

type ChoreContent struct {
	Title       *string `json:"title" validate:"omitempty,min=1,max=255"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Author      *string `json:"author" validate:"omitempty,min=1"`
	Schedule    *string `json:"schedule" validate:"omitempty"`
	Comment     *string `json:"comment" validate:"omitempty"`
}

func UnmarshalRequest(data []byte) (ToServiceRequest, error) {
	req := ToServiceRequest{}
	if err := json.Unmarshal(data, &req); err != nil {
		return req, fmt.Errorf("failed to unmarshal request: %v", err)
	}
	return req, nil
}

func (req *ToServiceRequest) Id() *int64 {
	return req.Identity.ID
}

func (req *ToServiceRequest) Content() ChoreContent {
	return req.Fields
}

///////////////////////depreacated

type CreateRequest struct {
	ChoreContent
}

func UnmarshalCreateRequest(data []byte) (CreateRequest, error) {
	r := CreateRequest{}
	if err := json.Unmarshal(data, &r); err != nil {
		return r, err
	}
	return r, nil
}

func (cr *CreateRequest) Content() ChoreContent {
	return cr.ChoreContent
}

type ReadRequest struct {
	ChoreIdentity
}

func UnmarshalReadRequest(data []byte) (ReadRequest, error) {
	r := ReadRequest{}
	if err := json.Unmarshal(data, &r); err != nil {
		return r, err
	}
	return r, nil
}

func (cr *ReadRequest) Id() *int64 {
	return cr.ID
}

type UpdateRequest struct {
	ChoreIdentity
	ChoreContent
}

func UnmarshalUpdateRequest(data []byte) (UpdateRequest, error) {
	r := UpdateRequest{}
	if err := json.Unmarshal(data, &r); err != nil {
		return r, err
	}
	return r, nil
}

func (ur *UpdateRequest) Id() *int64 {
	return ur.ID
}

func (ur *UpdateRequest) Content() ChoreContent {
	return ur.ChoreContent
}

type DeleteRequest struct {
	ChoreIdentity
}

func UnmarshalDeleteRequest(data []byte) (DeleteRequest, error) {
	r := DeleteRequest{}
	if err := json.Unmarshal(data, &r); err != nil {
		return r, err
	}
	return r, nil
}

func (dr *DeleteRequest) Id() *int64 {
	return dr.ID
}
