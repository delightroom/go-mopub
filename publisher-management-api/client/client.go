package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// LineItemOverrideFields is a subset field of LineItemResponseValue
type LineItemOverrideFields struct {
	Network_app_id     string `json:"network_app_id"`
	Network_adunit_id  string `json:"network_adunit_id"`
	Network_account_id string `json:"network_account_id"`
}

// LineItemResponseValue is a struct to get 'data' value of lineitem get request via MoPub Publisher Management API
type LineItemResponseValue struct {
	AdUnitKeys []string `json:"adUnitKeys"`
	Advertiser string   `json:"advertiser"`
	//docs-int
	AllocationPercentage float64  `json:"allocationPercentage"`
	AutoCpm              float64  `json:"autoCpm"`
	Bid                  float64  `json:"bid"`
	Budget               int64    `json:"budget"`
	BudgetStrategy       string   `json:"budgetStrategy"`
	BudgetType           string   `json:"budgetType"`
	DayParts             []string `json:"dayParts"`
	DayPartTargeting     string   `json:"dayPartTargeting"`
	DeviceTargeting      bool     `json:"deviceTargeting"`
	DisallowAutoCpm      bool     `json:"disallowAutoCpm"`
	//docs-int
	MaxAndroidVersion string `json:"maxAndroidVersion"`
	//docs-double
	MinAndroidVersion string `json:"minAndroidVersion"`
	//docs-int
	MaxIosVersion string `json:"maxIosVersion"`
	//docs-double
	MinIosVersion string `json:"minIosVersion"`
	TargetAndroid bool   `json:"targetAndroid"`
	TargetIos     bool   `json:"targetIos"`
	TargetIphone  bool   `json:"targetIphone"`
	TargetIpad    bool   `json:"targetIpad"`
	TargetIpod    bool   `json:"targetIpod"`
	//order might be different
	Idfa_targeting               string   `json:"idfa_targeting"`
	End                          string   `json:"end"`
	FrequencyCaps                []string `json:"frequencyCaps"`
	FrequencyCapsEnabled         bool     `json:"frequencyCapsEnabled"`
	IncludeConnectivityTargeting string   `json:"includeConnectivityTargeting"`
	TargetedCarriers             []string `json:"targetedCarriers"`
	IncludeGeoTargeting          string   `json:"includeGeoTargeting"`
	Key                          string   `json:"key"`
	Keywords                     []string `json:"keywords"`
	Name                         string   `json:"name"`
	NetworkType                  string   `json:"networkType"`
	OrderKey                     string   `json:"orderKey"`
	OrderName                    string   `json:"orderName"`
	Priority                     int      `json:"priority"`
	RefreshInterval              int      `json:"refreshInterval"`
	Start                        string   `json:"start"`
	Status                       string   `json:"status"`
	TargetedCountries            []string `json:"targetedCountries"`
	TargetedRegions              []string `json:"targetedRegions"`
	TargetedCities               []string `json:"targetedCities"`
	TargetedZipCodes             []string `json:"targetedZipCodes"`
	Type                         string   `json:"type"`
	UserAppsTargeting            string   `json:"userAppsTargeting"`
	UserAppsTargetingList        []string `json:"userAppsTargetingList"`
	//optinal? doesn't exist on our test requests
	EnableOverrides bool `json:"enableOverrides"`
	//optinal? doesn't exist on our test requests
	LineItemOverrideFields LineItemOverrideFields `json:"LineItemOverrideFields"`
	//optinal? doesn't exist on docs but our test requests
	AllowVideo string `json:"allowVideo"`
	//optinal? doesn't exist on docs but our test requests
	VideoSetting string `json:"videoSetting"`
}

//LineItemResponse is a struct to get a response of MoPub Publisher Management lineitem get API request
type LineItemResponse struct {
	LineItemResponseValue LineItemResponseValue `json:"data"`
}

//LineItemPutBodyData is subset of LineItemUpdateBody to carry a bid number to be updated
type LineItemPutBodyData struct {
	Name         string  `json:"name,omitempty"`
	Bid          float64 `json:"bid,omitempty"`
	AllowVideo   string  `json:"allowVideo,omitempty"`   //Possible values: non_video, all, video
	VideoSetting string  `json:"videoSetting,omitempty"` //Possible values: both, non_skippable, skippable
	Enabled      bool    `json:"enabled,omitempty"`      //note : Enabled doesn't exist on get&put response -> needs to be tested in the browser for now
	Archived     bool    `json:"archived,omitempty"`     //true -> status(archived) & false ->status(campaign-archived)
	Status       string  `json:"status,omitempty"`       //to see the output of Archived field...! not to change status directly(not supported)
	Budget       int64   `json:"budget,omitempty"`       //Must specify budget if budget type not unlimited
	BudgetType   string  `json:"budgetType,omitempty"`   //if budgtType is limited, budget becomes null
}

//LineItemPutBody is a struct for a body parameter of Mopub lineitem post API
type LineItemPutBody struct {
	Op   string              `json:"op"`
	Data LineItemPutBodyData `json:"data"`
}

var DefaultBaseUrl = "https://api.mopub.com/v2/line-items/"

//QQQ
//Client is an interface to make get/post request to MoPub publisher management API
type Client interface {
	GetLineItem(lineItemId string) (string, error)
	PutLineItem(lineItemId string, lineItem LineItemPutBodyData) (string, error)
}

//ApiClient is an implementation of a client interface with Apikey and Mopub Publisher management API url
type ApiClient struct {
	ApiKey  string `json:"api_key"`
	BaseUrl string `json:"base_url"`
}

// NewClient makes a new Api client for Mopub Publisher management API calls
func NewClient(apiKey, baseUrl string) ApiClient {
	return ApiClient{ApiKey: apiKey, BaseUrl: baseUrl}
}

// GenerateApiClient makes a new Api client for Mopub Publisher management API calls
func GenerateApiClient(apiKey string) ApiClient {
	return ApiClient{ApiKey: apiKey, BaseUrl: DefaultBaseUrl}
}

func (a ApiClient) GetLineItem(lineItemId string) (LineItemResponseValue, error) {
	mopubUrl := a.BaseUrl + lineItemId
	req, err := http.NewRequest("GET", mopubUrl, nil)

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("x-api-key", a.ApiKey)

	if err != nil {
		return LineItemResponseValue{}, err
	}

	// initialize http client
	client := &http.Client{}
	if err != nil {
		return LineItemResponseValue{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return LineItemResponseValue{}, err
	}

	defer resp.Body.Close()
	fmt.Println("📩 getApi resp.StatusCode:", resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)

	var LineItemResponse LineItemResponse

	err = json.Unmarshal([]byte(str), &LineItemResponse)
	if err != nil {
		return LineItemResponseValue{}, err
	}

	result := LineItemResponse.LineItemResponseValue
	return result, nil
}

func (a ApiClient) PutLineItem(lineItemId string, lineItem LineItemPutBodyData) (LineItemResponseValue, error) {
	mopubUrl := a.BaseUrl + lineItemId

	c := &LineItemPutBody{
		Op:   "set",
		Data: lineItem,
	}

	buff, err := json.Marshal(c)

	if err != nil {
		return LineItemResponseValue{}, errors.Wrap(err, "Marshal...")
	}
	req, err := http.NewRequest(http.MethodPut, mopubUrl, bytes.NewBuffer(buff))

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("x-api-key", a.ApiKey)

	if err != nil {
		return LineItemResponseValue{}, errors.Wrap(err, "NewRequest...")
	}

	// initialize http client
	client := &http.Client{}
	if err != nil {
		return LineItemResponseValue{}, errors.Wrap(err, "client...")
	}

	resp, err := client.Do(req)
	if err != nil {
		return LineItemResponseValue{}, errors.Wrap(err, "client.Do...")
	}

	defer resp.Body.Close()
	fmt.Println("📮postApi resp.StatusCode:", resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println("str result of response...", str)
	var LineItemResponse LineItemResponse

	err = json.Unmarshal([]byte(str), &LineItemResponse)

	if err != nil {
		return LineItemResponseValue{}, err
	}

	result := LineItemResponse.LineItemResponseValue
	return result, err
}
