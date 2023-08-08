package main

import (
	"encoding/xml"
	"os"
	"testing"
	"unsafe"
)


func TestParseRouteTable(t *testing.T) {
	// 1. Read the XML file from disk.
	xmlFile, err := os.ReadFile("route_table.xml")
	if err != nil {
		t.Fatalf("Error reading XML file: %v", err)
	}

	// 2. Parse the XML file into the specified struct.
	var reply RpcReplyOptimized
	err = xml.Unmarshal(xmlFile, &reply)
	if err != nil {
		t.Fatalf("Error parsing XML: %v", err)
	}

	// 3. Verify the parsed data matches the expected results.
	// For the purpose of this example, let's check the Junos attribute as a basic test.
	expectedJunosAttr := "http://xml.juniper.net/junos/18.1R3/junos" // This should be replaced with your expected value.
	if reply.Junos != expectedJunosAttr {
		t.Errorf("Expected Junos attribute to be %s, but got %s", expectedJunosAttr, reply.Junos)
	}
}

func TestMemoryFootprintComparison(t *testing.T) {
	xmlFile, err := os.ReadFile("route_table.xml")
	if err != nil {
		t.Fatalf("Error reading XML file: %v", err)
	}

	var replyOptimized RpcReplyOptimized
	err = xml.Unmarshal(xmlFile, &replyOptimized)
	if err != nil {
		t.Fatalf("Error parsing XML into RpcReplyOptimized: %v", err)
	}

	var reply RpcReply
	err = xml.Unmarshal(xmlFile, &reply)
	if err != nil {
		t.Fatalf("Error parsing XML into RpcReply: %v", err)
	}

	// Compare the sizes of the two parsed structures
	sizeOptimized := unsafe.Sizeof(replyOptimized)
	sizeReply := unsafe.Sizeof(reply)

	t.Logf("Size of RpcReplyOptimized: %v bytes", sizeOptimized)
	t.Logf("Size of RpcReply: %v bytes", sizeReply)

	// Example verification for RpcReplyOptimized
	expectedJunosAttr := "http://xml.juniper.net/junos/18.1R3/junos"
	if replyOptimized.Junos != expectedJunosAttr {
		t.Errorf("Expected Junos attribute for RpcReplyOptimized to be %s, but got %s", expectedJunosAttr, replyOptimized.Junos)
	}

	// Example verification for RpcReply
	if reply.Junos != expectedJunosAttr {
		t.Errorf("Expected Junos attribute for RpcReply to be %s, but got %s", expectedJunosAttr, reply.Junos)
	}
}
