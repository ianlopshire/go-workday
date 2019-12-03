

# extensibility
`import "github.com/ianlopshire/go-workday/pkg/wws/services/extensibility"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package extensibility

Public Web Service for extensibility features across applications.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type Client](#Client)
  * [func (c *Client) GetCustomLists(ctx context.Context, input *GetCustomListsInput) (output *GetCustomListsOutput, err error)](#Client.GetCustomLists)
  * [func (c *Client) PutCustomList(ctx context.Context, input *PutCustomListInput) (output *PutCustomListOutput, err error)](#Client.PutCustomList)
* [type CustomListDataType](#CustomListDataType)
* [type CustomListObjectIDType](#CustomListObjectIDType)
* [type CustomListObjectType](#CustomListObjectType)
* [type CustomListRequestCriteriaType](#CustomListRequestCriteriaType)
* [type CustomListRequestReferencesType](#CustomListRequestReferencesType)
* [type CustomListResponseDataType](#CustomListResponseDataType)
* [type CustomListResponseGroupType](#CustomListResponseGroupType)
* [type CustomListType](#CustomListType)
* [type CustomListValueType](#CustomListValueType)
* [type GetCustomListsInput](#GetCustomListsInput)
* [type GetCustomListsOutput](#GetCustomListsOutput)
* [type GetCustomListsRequestType](#GetCustomListsRequestType)
* [type GetCustomListsResponseType](#GetCustomListsResponseType)
* [type ProcessingFaultType](#ProcessingFaultType)
* [type PutCustomListInput](#PutCustomListInput)
* [type PutCustomListOutput](#PutCustomListOutput)
* [type PutCustomListRequestType](#PutCustomListRequestType)
* [type PutCustomListResponseType](#PutCustomListResponseType)
* [type ResponseFilterType](#ResponseFilterType)
  * [func (t *ResponseFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ResponseFilterType.MarshalXML)
  * [func (t *ResponseFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ResponseFilterType.UnmarshalXML)
* [type ResponseResultsType](#ResponseResultsType)
* [type ValidationErrorType](#ValidationErrorType)
* [type ValidationFaultType](#ValidationFaultType)
* [type WorkdayCommonHeaderType](#WorkdayCommonHeaderType)


#### <a name="pkg-files">Package files</a>
[wwsgen.go](/src/github.com/ianlopshire/go-workday/pkg/wws/services/extensibility/wwsgen.go) [wwsgen_client.go](/src/github.com/ianlopshire/go-workday/pkg/wws/services/extensibility/wwsgen_client.go) [wwsgen_types.go](/src/github.com/ianlopshire/go-workday/pkg/wws/services/extensibility/wwsgen_types.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    CurrentVersion = "v31.2"
    ServiceName    = "Workday_Extensibility"
)
```




## <a name="Client">type</a> [Client](/src/target/wwsgen_client.go?s=267:302#L18)
``` go
type Client struct {
    *wws.Client
}
```









### <a name="Client.GetCustomLists">func</a> (\*Client) [GetCustomLists](/src/target/wwsgen_client.go?s=451:573#L26)
``` go
func (c *Client) GetCustomLists(ctx context.Context, input *GetCustomListsInput) (output *GetCustomListsOutput, err error)
```
GetCustomLists (Get_Custom_Lists)

This service operation allows you to get all custom lists or custom lists for the specified criteria.




### <a name="Client.PutCustomList">func</a> (\*Client) [PutCustomList](/src/target/wwsgen_client.go?s=1112:1231#L48)
``` go
func (c *Client) PutCustomList(ctx context.Context, input *PutCustomListInput) (output *PutCustomListOutput, err error)
```
PutCustomList (Put_Custom_List)

This service operation allows you to add and update a custom list.




## <a name="CustomListDataType">type</a> [CustomListDataType](/src/target/wwsgen_types.go?s=88:525#L10)
``` go
type CustomListDataType struct {
    WebServiceAlias     string                `xml:"urn:com.workday/bsvc Web_Service_Alias,omitempty"`
    CustomFieldTypeName string                `xml:"urn:com.workday/bsvc Custom_Field_Type_Name,omitempty"`
    DoNotUse            *bool                 `xml:"urn:com.workday/bsvc Do_Not_Use,omitempty"`
    CustomListValueData []CustomListValueType `xml:"urn:com.workday/bsvc Custom_List_Value_Data,omitempty"`
}
```
Custom List Data










## <a name="CustomListObjectIDType">type</a> [CustomListObjectIDType](/src/target/wwsgen_types.go?s=589:712#L18)
``` go
type CustomListObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CustomListObjectType">type</a> [CustomListObjectType](/src/target/wwsgen_types.go?s=714:921#L23)
``` go
type CustomListObjectType struct {
    ID         []CustomListObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CustomListRequestCriteriaType">type</a> [CustomListRequestCriteriaType](/src/target/wwsgen_types.go?s=973:1196#L29)
``` go
type CustomListRequestCriteriaType struct {
    CustomListAlias      string `xml:"urn:com.workday/bsvc Custom_List_Alias,omitempty"`
    CustomListValueAlias string `xml:"urn:com.workday/bsvc Custom_List_Value_Alias,omitempty"`
}
```
Request Criteria for Bringing back Custom List










## <a name="CustomListRequestReferencesType">type</a> [CustomListRequestReferencesType](/src/target/wwsgen_types.go?s=1223:1365#L35)
``` go
type CustomListRequestReferencesType struct {
    CustomListReference []CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference"`
}
```
Custom List Reference










## <a name="CustomListResponseDataType">type</a> [CustomListResponseDataType](/src/target/wwsgen_types.go?s=1392:1514#L40)
``` go
type CustomListResponseDataType struct {
    CustomList []CustomListType `xml:"urn:com.workday/bsvc Custom_List,omitempty"`
}
```
Custom List and  Data










## <a name="CustomListResponseGroupType">type</a> [CustomListResponseGroupType](/src/target/wwsgen_types.go?s=1546:1670#L45)
``` go
type CustomListResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Custom List Response Group










## <a name="CustomListType">type</a> [CustomListType](/src/target/wwsgen_types.go?s=1706:1939#L50)
``` go
type CustomListType struct {
    CustomListReference *CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference,omitempty"`
    CustomListData      []CustomListDataType  `xml:"urn:com.workday/bsvc Custom_List_Data,omitempty"`
}
```
Custom List Reference and Data










## <a name="CustomListValueType">type</a> [CustomListValueType](/src/target/wwsgen_types.go?s=1967:2295#L56)
``` go
type CustomListValueType struct {
    Order           string `xml:"urn:com.workday/bsvc Order,omitempty"`
    ListValueName   string `xml:"urn:com.workday/bsvc List_Value_Name,omitempty"`
    WebServiceAlias string `xml:"urn:com.workday/bsvc Web_Service_Alias"`
    Deprecated      *bool  `xml:"urn:com.workday/bsvc Deprecated,omitempty"`
}
```
Custom List Value Data










## <a name="GetCustomListsInput">type</a> [GetCustomListsInput](/src/target/wwsgen_client.go?s=732:864#L35)
``` go
type GetCustomListsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Custom_Lists_Request"`
    GetCustomListsRequestType
}
```









## <a name="GetCustomListsOutput">type</a> [GetCustomListsOutput](/src/target/wwsgen_client.go?s=866:1001#L40)
``` go
type GetCustomListsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Custom_Lists_Response"`
    GetCustomListsResponseType
}
```









## <a name="GetCustomListsRequestType">type</a> [GetCustomListsRequestType](/src/target/wwsgen_types.go?s=2355:2931#L64)
``` go
type GetCustomListsRequestType struct {
    RequestReferences *CustomListRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CustomListRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CustomListResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request for a Specific Custom List or All Custom Lists










## <a name="GetCustomListsResponseType">type</a> [GetCustomListsResponseType](/src/target/wwsgen_types.go?s=2961:3751#L73)
``` go
type GetCustomListsResponseType struct {
    RequestReferences *CustomListRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CustomListRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CustomListResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *CustomListResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Custom List Response










## <a name="ProcessingFaultType">type</a> [ProcessingFaultType](/src/target/wwsgen_types.go?s=3753:3864#L83)
``` go
type ProcessingFaultType struct {
    DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}
```









## <a name="PutCustomListInput">type</a> [PutCustomListInput](/src/target/wwsgen_client.go?s=1389:1518#L57)
``` go
type PutCustomListInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Custom_List_Request"`
    PutCustomListRequestType
}
```









## <a name="PutCustomListOutput">type</a> [PutCustomListOutput](/src/target/wwsgen_client.go?s=1520:1652#L62)
``` go
type PutCustomListOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Custom_List_Response"`
    PutCustomListResponseType
}
```









## <a name="PutCustomListRequestType">type</a> [PutCustomListRequestType](/src/target/wwsgen_types.go?s=3963:4397#L88)
``` go
type PutCustomListRequestType struct {
    CustomListReference *CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference,omitempty"`
    CustomListData      *CustomListDataType   `xml:"urn:com.workday/bsvc Custom_List_Data,omitempty"`
    AddOnly             bool                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request Element . A custom list can be edited by specifying a reference or a webservice alias










## <a name="PutCustomListResponseType">type</a> [PutCustomListResponseType](/src/target/wwsgen_types.go?s=4435:4675#L96)
``` go
type PutCustomListResponseType struct {
    CustomListReference *CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference,omitempty"`
    Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response Element for Custom List










## <a name="ResponseFilterType">type</a> [ResponseFilterType](/src/target/wwsgen_types.go?s=4805:5168#L102)
``` go
type ResponseFilterType struct {
    AsOfEffectiveDate *time.Time `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
    AsOfEntryDateTime *time.Time `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
    Page              float64    `xml:"urn:com.workday/bsvc Page,omitempty"`
    Count             float64    `xml:"urn:com.workday/bsvc Count,omitempty"`
}
```
Parameters that let you filter the data returned in the response. You can filter returned data by dates and page attributes.










### <a name="ResponseFilterType.MarshalXML">func</a> (\*ResponseFilterType) [MarshalXML](/src/target/wwsgen_types.go?s=5170:5255#L109)
``` go
func (t *ResponseFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ResponseFilterType.UnmarshalXML">func</a> (\*ResponseFilterType) [UnmarshalXML](/src/target/wwsgen_types.go?s=5699:5786#L121)
``` go
func (t *ResponseFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ResponseResultsType">type</a> [ResponseResultsType](/src/target/wwsgen_types.go?s=6434:6757#L135)
``` go
type ResponseResultsType struct {
    TotalResults float64 `xml:"urn:com.workday/bsvc Total_Results,omitempty"`
    TotalPages   float64 `xml:"urn:com.workday/bsvc Total_Pages,omitempty"`
    PageResults  float64 `xml:"urn:com.workday/bsvc Page_Results,omitempty"`
    Page         float64 `xml:"urn:com.workday/bsvc Page,omitempty"`
}
```
The "Response_Results" element contains summary information about the data that has been returned from your request including "Total_Results", "Total_Pages", and the current "Page" returned.










## <a name="ValidationErrorType">type</a> [ValidationErrorType](/src/target/wwsgen_types.go?s=6759:7006#L142)
``` go
type ValidationErrorType struct {
    Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
    DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
    Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}
```









## <a name="ValidationFaultType">type</a> [ValidationFaultType](/src/target/wwsgen_types.go?s=7008:7138#L148)
``` go
type ValidationFaultType struct {
    ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}
```









## <a name="WorkdayCommonHeaderType">type</a> [WorkdayCommonHeaderType](/src/target/wwsgen_types.go?s=7140:7305#L152)
``` go
type WorkdayCommonHeaderType struct {
    IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}
```













- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
