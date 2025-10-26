package dto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Galdoba/choretracker/internal/constants"
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

func (res *ToServiceRequest) RequestType() string {
	return res.Action
}

type ChoreIdentity struct {
	ID *int64 `json:"id" validate:"required"`
}

func (ci *ChoreIdentity) GetID() (int64, bool) {
	if ci.ID == nil {
		return 0, false
	}
	if *ci.ID == 0 {
		return 0, false
	}
	return *ci.ID, true
}

type ChoreContent struct {
	Title       *string `json:"title" validate:"omitempty,min=1,max=255"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Author      *string `json:"author" validate:"omitempty,min=1"`
	Schedule    *string `json:"schedule" validate:"omitempty"`
	Comment     *string `json:"comment" validate:"omitempty"`
}

func (cnt *ChoreContent) Content() map[string]string {
	c := make(map[string]string)
	if cnt.Title != nil {
		c[constants.Fld_Title] = *cnt.Title
	}
	if cnt.Description != nil {
		c[constants.Fld_Descr] = *cnt.Description
	}
	if cnt.Author != nil {
		c[constants.Fld_Author] = *cnt.Author
	}
	if cnt.Schedule != nil {
		c[constants.Fld_Schedule] = *cnt.Schedule
	}
	if cnt.Comment != nil {
		c[constants.Fld_Comment] = *cnt.Comment
	}
	return c
}

func UnmarshalRequest(data []byte) (ToServiceRequest, error) {
	req := ToServiceRequest{}
	if err := json.Unmarshal(data, &req); err != nil {
		return req, fmt.Errorf("failed to unmarshal request: %v", err)
	}
	return req, nil
}

func UnURL(url string) *ChoreContent {
	// cc := ChoreContent{}
	// re := regexp.MustCompile(`/?()`)
}

func (req *ToServiceRequest) GetID() (int64, bool) {
	if req.Identity.ID == nil {
		return 0, false
	}
	return *req.Identity.ID, true
}

func (req *ToServiceRequest) InjectID(id int64) {
	if id == 0 {
		return
	}
	req.Identity.ID = &id
}

func (req *ToServiceRequest) InjectContent(content map[string]string) {
	if val, ok := content[constants.Fld_Title]; ok {
		req.Fields.Title = &val
	}
	if val, ok := content[constants.Fld_Descr]; ok {
		req.Fields.Description = &val
	}
	if val, ok := content[constants.Fld_Author]; ok {
		req.Fields.Author = &val
	}
	if val, ok := content[constants.Fld_Schedule]; ok {
		req.Fields.Schedule = &val
	}
	if val, ok := content[constants.Fld_Comment]; ok {
		req.Fields.Comment = &val
	}

}
