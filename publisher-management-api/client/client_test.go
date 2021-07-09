package client

import (
	"os"
	"testing"
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

func TestPutLineItem(t *testing.T) {
	apiKey := os.Getenv("MOPUB_PM_API_KEY")
	baseUrl := DefaultBaseUrl
	c := NewClient(apiKey, baseUrl)
	testKey := "fcc018399741425798e3503b554dd21d"
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
	}

	want := LineItemPutBodyData{
		Bid:          0.4,
		Name:         "Mopub_T7 - WW_OS_TEST",
		AllowVideo:   "video",
		VideoSetting: "skippable",
		Status:       "archived",
		BudgetType:   "unlimited",
		// Budget:       3,
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
