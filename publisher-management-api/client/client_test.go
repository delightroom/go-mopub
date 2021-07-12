package client

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

//Test GetLineItem will test GetLineItem method
//required env :MOPUB_PM_API_KEY
func TestGetLineItem(t *testing.T) {
	apiKey := os.Getenv("MOPUB_PM_API_KEY")
	baseUrl := DefaultBaseUrl
	c := NewClient(apiKey, baseUrl)
	testKey := "69c0c8d4859341ffb61e06d15b3cccfc"
	got, err := c.GetLineItem(testKey)
	if err != nil {
		t.Error(err)
	}
	gotBid := got.Bid
	want := 0.01

	if gotBid != want {
		t.Errorf("got %f want %f", gotBid, want)
	}
}

//Test TestPutLineItem will test PutLineItem method
//required env :MOPUB_PM_API_KEY
func TestPutLineItem(t *testing.T) {
	apiKey := os.Getenv("MOPUB_PM_API_KEY")
	baseUrl := DefaultBaseUrl
	c := NewClient(apiKey, baseUrl)
	testKey := "fcc018399741425798e3503b554dd21d"
	adUnitKeys := []string{"a6c828b21aba4ae6bf2d8d7bfaf87e83", "1081d00562a24ca5869147a7e1a9a7df", "614f31b0392848a7851f4fc13115d5e6", "996c13fec41c4ef6a18a3d93104050ff", "5f5107f7ff7549b7a7b881b394057389", "a4184b7f4b4f4af19dc0c5db03b40fc8", "47648e92fd96424d964c1570f350f2ae"}
	lineItem := LineItemPutBodyData{
		Bid:          0.4,
		Name:         "Mopub_T7 - WW_OS_TEST",
		AllowVideo:   "video", //false
		VideoSetting: "skippable",
		Enabled:      false, //Enabled 값의 경우 response에 없어서 현재로써는 따로 browser에서 테스트 결과를 확인해야 한다.
		//Enabled 값만 바꿔서 호출하면 변경이 안되는데, 다른값도 바꿔서 전송하면 변경이 된다???
		Archived:   true, //true -> status(archived), false -> status(campaign-archived)
		BudgetType: "unlimited",
		// Budget:     3,
		AdUnitKeys: adUnitKeys,
	}

	resp, err := c.PutLineItem(testKey, lineItem)
	if err != nil {
		t.Error(err)
	}

	got := LineItemPutBodyData{
		Bid:          resp.Bid,
		Name:         resp.Name,
		AllowVideo:   resp.AllowVideo,
		VideoSetting: resp.VideoSetting,
		Status:       resp.Status,
		BudgetType:   resp.BudgetType,
		Budget:       resp.Budget,
		AdUnitKeys:   resp.AdUnitKeys,
	}

	want := LineItemPutBodyData{
		Bid:          0.4,
		Name:         "Mopub_T7 - WW_OS_TEST",
		AllowVideo:   "video",
		VideoSetting: "skippable",
		Status:       "archived",
		BudgetType:   "unlimited",
		// Budget:       3,
		AdUnitKeys: adUnitKeys,
	}

	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}
