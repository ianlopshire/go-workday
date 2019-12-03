
// Package extensibility
//
// Public Web Service for extensibility features across applications.
package extensibility

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Workday_Extensibility"
)

type Client struct {
	*wws.Client
}


// GetCustomLists (Get_Custom_Lists)
// 
// This service operation allows you to get all custom lists or custom lists for the specified criteria.
func (c *Client) GetCustomLists(ctx context.Context, input *GetCustomListsInput) (output *GetCustomListsOutput, err error) {
	output = &GetCustomListsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCustomListsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Custom_Lists_Request"`
	GetCustomListsRequestType
}

type GetCustomListsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Custom_Lists_Response"`
	GetCustomListsResponseType
}

// PutCustomList (Put_Custom_List)
// 
// This service operation allows you to add and update a custom list.
func (c *Client) PutCustomList(ctx context.Context, input *PutCustomListInput) (output *PutCustomListOutput, err error) {
	output = &PutCustomListOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCustomListInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Custom_List_Request"`
	PutCustomListRequestType
}

type PutCustomListOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Custom_List_Response"`
	PutCustomListResponseType
}

