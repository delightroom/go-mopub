package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OverrideFields is a subset field of LineItemGetResponseValue
type OverrideFields struct {
	Network_app_id     string `json:"network_app_id"`
	Network_adunit_id  string `json:"network_adunit_id"`
	Network_account_id string `json:"network_account_id"`
}

// LineItemGetResponseValue is a struct to get 'data' value of lineitem get request via MoPub Publisher Management API
type LineItemGetResponseValue struct {
	AdUnitKeys                   []string       `json:"adUnitKeys"`
	Advertiser                   string         `json:"advertiser"`
	AllocationPercentage         float64        `json:"allocationPercentage"`
	AutoCpm                      float64        `json:"autoCpm"`
	Bid                          float64        `json:"bid"`
	Budget                       string         `json:"budget"`
	BudgetStrategy               string         `json:"budgetStrategy"`
	BudgetType                   string         `json:"budgetType"`
	DayParts                     []string       `json:"dayParts"`
	DayPartTargeting             string         `json:"dayPartTargeting"`
	DeviceTargeting              bool           `json:"deviceTargeting"`
	DisallowAutoCpm              bool           `json:"disallowAutoCpm"`
	MaxAndroidVersion            int            `json:"maxAndroidVersion"`
	MinAndroidVersion            float64        `json:"minAndroidVersion"`
	MaxIosVersion                int            `json:"maxIosVersion"`
	MinIosVersion                float64        `json:"minIosVersion"`
	TargetAndroid                bool           `json:"targetAndroid"`
	TargetIos                    bool           `json:"targetIos"`
	TargetIphone                 bool           `json:"targetIphone"`
	TargetIpad                   bool           `json:"targetIpad"`
	TargetIpod                   bool           `json:"targetIpod"`
	Idfa_targeting               string         `json:"idfa_targeting"`
	End                          string         `json:"end"`
	FrequencyCaps                []string       `json:"frequencyCaps"`
	FrequencyCapsEnabled         bool           `json:"frequencyCapsEnabled"`
	IncludeConnectivityTargeting string         `json:"includeConnectivityTargeting"`
	TargetedCarriers             []string       `json:"targetedCarriers"`
	IncludeGeoTargeting          string         `json:"includeGeoTargeting"`
	Key                          string         `json:"key"`
	Keywords                     []string       `json:"keywords"`
	Name                         string         `json:"name"`
	NetworkType                  string         `json:"networkType"`
	OrderKey                     string         `json:"orderKey"`
	OrderName                    string         `json:"orderName"`
	Priority                     int            `json:"priority"`
	RefreshInterval              int            `json:"refreshInterval"`
	Start                        string         `json:"start"`
	Status                       string         `json:"status"`
	TargetedCountries            []string       `json:"targetedCountries"`
	TargetedRegions              []string       `json:"targetedRegions"`
	TargetedCities               []string       `json:"targetedCities"`
	TargetedZipCodes             []string       `json:"targetedZipCodes"`
	Type                         string         `json:"type"`
	UserAppsTargeting            string         `json:"userAppsTargeting"`
	UserAppsTargetingList        []string       `json:"userAppsTargetingList"`
	EnableOverrides              bool           `json:"enableOverrides"`
	OverrideFields               OverrideFields `json:"overrideFields"`
}

//LineItemGetResponse is a struct to get value of line item via MoPub Publisher Management API
type LineItemGetResponse struct {
	LineItemGetResponseValue LineItemGetResponseValue `json:"data"`
}

//LineItemPostBodyData is subset of LineItemUpdateBody to carry a bid number to be updated
type LineItemPostBodyData struct {
	Bid float64 `json:"bid"`
}

//LineItemPostBody is a struct for a body parameter of Mopub lineitem post API
type LineItemPostBody struct {
	Op   string               `json:"op"`
	Data LineItemPostBodyData `json:"data"`
}

var BaseUrl = "https://api.mopub.com/v2/line-items/"

//client is an interface to make get/post request to MoPub publisher management API
type client interface {
	GetLineItem(aa string) (string, error)
}

//apiClient is an implementation of a client interface with Apikey and Mopub Publisher management API url
type apiClient struct {
	ApiKey  string `json:"api_key"`
	BaseUrl string `json:"base_url"`
}

// MakeNewApiClient makes a new Api client for Mopub Publisher management API calls
func MakeNewApiClient(apiKey string) apiClient {
	return apiClient{ApiKey: apiKey, BaseUrl: BaseUrl}
}

func (a apiClient) GetLineItem(lineItemId string) (LineItemGetResponseValue, error) {
	mopubGetUrl := a.BaseUrl + lineItemId
	req, err := http.NewRequest("GET", mopubGetUrl, nil)

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("x-api-key", a.ApiKey)

	if err != nil {
		return LineItemGetResponseValue{}, err
	}

	// initialize http client
	client := &http.Client{}
	if err != nil {
		return LineItemGetResponseValue{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return LineItemGetResponseValue{}, err
	}

	defer resp.Body.Close()
	fmt.Println("📩 getApi resp.StatusCode:", resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println("str...", str)
	var lineItemGetResponse LineItemGetResponse

	err = json.Unmarshal([]byte(str), &lineItemGetResponse)
	if err != nil {
		return LineItemGetResponseValue{}, err
	}

	result := lineItemGetResponse.LineItemGetResponseValue
	return result, nil
}

// func (a apiClient) PostLineItem(lineItemId string){}
