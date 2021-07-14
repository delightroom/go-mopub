# 🔖 go-mopub
`go-mopub` is a library to make MoPub Api requests with golang


`Version` Info: v0.1.8<br/><br/>

## How to get
```go
    go get github.com/delightroom/go-mopub
```
<br/>

## MoPub Publishment Management Api
### How to initiate a client
```go
apiClient := client.NewClient(MOPUB_API_KEY, MOPUB_API_URL)
```

</br>To make client with given url
```go
apiClient := GenerateApiClient(MOPUB_API_KEY)
```

<br>

### Currently available methods
1. lineitem get request
```go
resp, err := apiClient.GetLineItem(lineItemID)
```

2. lineitem put (Bid) request

```go
lineItem := client.LineItemPutBodyData{
    Bid: 0.1 //new bid value
}
resp, err := apiClient.PutLineItem(lineItemID, lineItem)
```

<details><summary>This library supports below fields for lineitem put request</summary>
<p>

```go
type LineItemPutBodyData struct {
Name         string   `json:"name,omitempty"`
Bid          float64  `json:"bid,omitempty"`
//Possible values: non_video, all, video
AllowVideo   string   `json:"allowVideo,omitempty"`   
//Possible values: both, non_skippable, skippable
VideoSetting string   `json:"videoSetting,omitempty"` 
//Enabled doesn't exist on get&put response -> needs to be tested in the browser for now
Enabled      bool     `json:"enabled,omitempty"`      
//true -> status(archived) & false ->status(campaign-archived)
Archived     bool     `json:"archived,omitempty"`     
//to see the output of Archived field...! not to change status directly(not supported)
Status       string   `json:"status,omitempty"`       
//Must specify budget if budget type not unlimited
Budget       int64    `json:"budget,omitempty"`       
//if budgtType is limited, budget becomes null
BudgetType   string   `json:"budgetType,omitempty"`   
AdUnitKeys   []string `json:"adUnitKeys,omitempty"`
```

</p>
</details></br>


