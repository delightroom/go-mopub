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
	}

	want := LineItemPutBodyData{
		Bid:          0.4,
		Name:         "Mopub_T7 - WW_OS_TEST",
		AllowVideo:   "video",
		VideoSetting: "skippable",
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
