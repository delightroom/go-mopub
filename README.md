# 🔖 go-mopub
`go-mopub` is a library to make MoPub Api requests with golang


`Version` Info: v0.1.4<br/><br/>

## How to get
```go
    go get github.com/delightroom/go-mopub
```
<br/>

## How to initiate a client
```go
apiClient := client.GenerateApiClient(MOPUB_API_KEY)
```
<br/>

## Currently available methods

1. MoPub Publishment Management Api
    - lineitem get request
        ```go
        resp, err := apiClient.GetLineItem(lineItemID)
        ```
    - lineitem put (Bid) request
        ```go
        resp, err := apiClient.PutLineItemBid(lineItemID, newBidValue)
        ```
