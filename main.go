package main

import (
	    "fmt"
		"encoding/xml"
        "unsafe"
)

type RpcReplyOptimized struct {
	XMLName          xml.Name `xml:"rpc-reply"`
	Text             string   `xml:",chardata"`
	Junos            string   `xml:"junos,attr"`
	Cli *Cli `xml:"cli"`
	RouteInformation struct {
		Text       string `xml:",chardata"`
		Xmlns      string `xml:"xmlns,attr"`
		RouteTable []struct {
			Text               string `xml:",chardata"`
			TableName          string `xml:"table-name"`
			DestinationCount   int `xml:"destination-count"`
			TotalRouteCount    int `xml:"total-route-count"`
			ActiveRouteCount   int    `xml:"active-route-count"`
			HolddownRouteCount int `xml:"holddown-route-count"`
			HiddenRouteCount   int`xml:"hidden-route-count"`
			Rt                 []struct {
				Text          string `xml:",chardata"`
				Style         string `xml:"style,attr"`
				RtDestination struct {
					Text string `xml:",chardata"`
					Emit string `xml:"emit,attr"`
				} `xml:"rt-destination"`
				RtEntry struct {
					Text          string `xml:",chardata"`
					ActiveTag     string `xml:"active-tag"`
					CurrentActive string `xml:"current-active"`
					LastActive    string `xml:"last-active"`
					ProtocolName  string `xml:"protocol-name"`
					Preference    string `xml:"preference"`
					Age           struct {
						Text    string `xml:",chardata"`
						Seconds int`xml:"seconds,attr"`
					} `xml:"age"`
					Nh []struct {
						Text             string `xml:",chardata"`
						SelectedNextHop  string `xml:"selected-next-hop"`
						To               string `xml:"to"`
						Via              string `xml:"via"`
						NhLocalInterface string `xml:"nh-local-interface"`
					} `xml:"nh"`
					Metric int `xml:"metric"`
					NhType string `xml:"nh-type"`
				} `xml:"rt-entry"`
			} `xml:"rt"`
		} `xml:"route-table"`
	} `xml:"route-information"`
} 

type Cli struct {
	Text string `xml:",chardata"`
	Banner string `xml:"banner"`
  }

type RpcReply struct {
	XMLName          xml.Name `xml:"rpc-reply"`
	Text             string   `xml:",chardata"`
	Junos            string   `xml:"junos,attr"`
	RouteInformation struct {
		Text       string `xml:",chardata"`
		Xmlns      string `xml:"xmlns,attr"`
		RouteTable []struct {
			Text               string `xml:",chardata"`
			TableName          string `xml:"table-name"`
			DestinationCount   int `xml:"destination-count"`
			TotalRouteCount    int `xml:"total-route-count"`
			ActiveRouteCount   int    `xml:"active-route-count"`
			HolddownRouteCount int `xml:"holddown-route-count"`
			HiddenRouteCount   int `xml:"hidden-route-count"`
			Rt                 []struct {
				Text          string `xml:",chardata"`
				Style         string `xml:"style,attr"`
				RtDestination struct {
					Text string `xml:",chardata"`
					Emit string `xml:"emit,attr"`
				} `xml:"rt-destination"`
				RtEntry struct {
					Text          string `xml:",chardata"`
					ActiveTag     string `xml:"active-tag"`
					CurrentActive string `xml:"current-active"`
					LastActive    string `xml:"last-active"`
					ProtocolName  string `xml:"protocol-name"`
					Preference    string `xml:"preference"`
					Age           struct {
						Text    string `xml:",chardata"`
						Seconds string `xml:"seconds,attr"`
					} `xml:"age"`
					Nh []struct {
						Text             string `xml:",chardata"`
						SelectedNextHop  string `xml:"selected-next-hop"`
						To               string `xml:"to"`
						Via              string `xml:"via"`
						NhLocalInterface string `xml:"nh-local-interface"`
					} `xml:"nh"`
					Metric string `xml:"metric"`
					NhType string `xml:"nh-type"`
				} `xml:"rt-entry"`
			} `xml:"rt"`
		} `xml:"route-table"`
	} `xml:"route-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
} 

func main() {
	var r1 RpcReply
	var r2 RpcReplyOptimized
	fmt.Println("==============================================================")
	fmt.Printf("Total Memory Usage StructType: RpcReply %T => [%d]\n", r1, unsafe.Sizeof(r1))
	fmt.Println("==============================================================")
	fmt.Printf("Total Memory Usage StructType: RpcReplyOptimized %T => [%d]\n", r2, unsafe.Sizeof(r2))
	fmt.Println("==============================================================")
}