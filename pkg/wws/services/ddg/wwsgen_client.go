// Package ddg
//
// Web service for creating, editing and retrieving objects related to Document Templates, such as Text Blocks and Text Block Categories.
package ddg

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName    = "Dynamic_Document_Generation"
)

type Client struct {
	*wws.Client
}

// GetTextBlockCategories (Get_Text_Block_Categories)
//
// Returns text block categories.
func (c *Client) GetTextBlockCategories(ctx context.Context, input *GetTextBlockCategoriesInput) (output *GetTextBlockCategoriesOutput, err error) {
	output = &GetTextBlockCategoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetTextBlockCategoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Text_Block_Categories_Request"`
	GetTextBlockCategoriesRequestType
}

type GetTextBlockCategoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Text_Block_Categories_Response"`
	GetTextBlockCategoriesResponseType
}

// PutTextBlockCategory (Put_Text_Block_Category)
//
// Adds or updates a text block category.
func (c *Client) PutTextBlockCategory(ctx context.Context, input *PutTextBlockCategoryInput) (output *PutTextBlockCategoryOutput, err error) {
	output = &PutTextBlockCategoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutTextBlockCategoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Text_Block_Category_Request"`
	PutTextBlockCategoryRequestType
}

type PutTextBlockCategoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Text_Block_Category_Response"`
	PutTextBlockCategoryResponseType
}

// GetTextBlocks (Get_Text_Blocks)
//
// Returns text blocks.
func (c *Client) GetTextBlocks(ctx context.Context, input *GetTextBlocksInput) (output *GetTextBlocksOutput, err error) {
	output = &GetTextBlocksOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetTextBlocksInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Text_Blocks_Request"`
	GetTextBlocksRequestType
}

type GetTextBlocksOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Text_Blocks_Response"`
	GetTextBlocksResponseType
}

// PutTextBlock (Put_Text_Block)
//
// Adds or updates a text block.
func (c *Client) PutTextBlock(ctx context.Context, input *PutTextBlockInput) (output *PutTextBlockOutput, err error) {
	output = &PutTextBlockOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutTextBlockInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Text_Block_Request"`
	PutTextBlockRequestType
}

type PutTextBlockOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Text_Block_Response"`
	PutTextBlockResponseType
}
