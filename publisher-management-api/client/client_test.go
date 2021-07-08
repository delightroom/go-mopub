package client

import (
	"os"
	"testing"
)

// //Test GetLineItem will test GetLineItem method
// //required env :MOPUB_PM_API_KEY
// func TestGetLineItem(t *testing.T) {
// 	apiKey := os.Getenv("MOPUB_PM_API_KEY")
// 	baseUrl := DefaultBaseUrl
// 	c := NewClient(apiKey, baseUrl)
// 	testKey := "fcc018399741425798e3503b554dd21d"
// 	got, err := c.GetLineItem(testKey)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	gotBid := got.Bid
// 	want := 0.057

// 	if gotBid != want {
// 		t.Errorf("got %f want %f", gotBid, want)
// 	}
// }

func TestPutLineItem(t *testing.T) {
	apiKey := os.Getenv("MOPUB_PM_API_KEY")
	baseUrl := DefaultBaseUrl
	c := NewClient(apiKey, baseUrl)
	testKey := "fcc018399741425798e3503b554dd21d"
	lineItem := LineItemPutBodyData{
		Bid:  0.6,
		Name: "Mopub_T7 - WW_OS_TESTT",
	}
	got, err := c.PutLineItem(testKey, lineItem)
	if err != nil {
		t.Error(err)
	}
	gotBid := got.Bid
	want := 0.6

	if gotBid != want {
		t.Errorf("got %f want %f", gotBid, want)
	}
}
