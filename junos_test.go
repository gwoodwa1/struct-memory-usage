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
	var replyWithCliPtr RpcReplyWithCliPointer
	err = xml.Unmarshal(xmlFile, &replyWithCliPtr)
	if err != nil {
		t.Fatalf("Error parsing XML: %v", err)
	}

	// 3. Verify the parsed data matches the expected results.
	// For the purpose of this example, let's check the Junos attribute as a basic test.
	expectedJunosAttr := "http://xml.juniper.net/junos/18.1R3/junos" // This should be replaced with your expected value.
	if replyWithCliPtr.Junos != expectedJunosAttr {
		t.Errorf("Expected Junos attribute to be %s, but got %s", expectedJunosAttr, replyWithCliPtr.Junos)
	}
}

func TestMemoryFootprintComparison(t *testing.T) {
	xmlFile, err := os.ReadFile("route_table.xml")
	if err != nil {
		t.Fatalf("Error reading XML file: %v", err)
	}
	var replyWithCliPtr RpcReplyWithCliPointer
	err = xml.Unmarshal(xmlFile, &replyWithCliPtr)
	if err != nil {
		t.Fatalf("Error parsing XML into RpcReply: %v", err)
	}
	var replyWithTablePtr RpcReplyWithTablePointer
	err = xml.Unmarshal(xmlFile, &replyWithTablePtr)
	if err != nil {
		t.Fatalf("Error parsing XML into RpcReplyWithTablePointer: %v", err)
	}

	var originalReply RpcReply
	err = xml.Unmarshal(xmlFile, &originalReply)
	if err != nil {
		t.Fatalf("Error parsing XML into RpcReply: %v", err)
	}

	// Compare the sizes of the three parsed structures
	sizeWithCliPtr := unsafe.Sizeof(replyWithCliPtr)
	sizeWithTablePtr := unsafe.Sizeof(replyWithTablePtr)
	sizeOriginalReply := unsafe.Sizeof(originalReply)
        t.Logf("Size of RpcReplyWithCliPointer: %v bytes", sizeWithCliPtr)
	t.Logf("Size of RpcReplyWithTablePointer: %v bytes", sizeWithTablePtr)
	t.Logf("Size of RpcReply: %v bytes", sizeOriginalReply)

	// Example verification for RpcReplyWithTablePointer
	expectedJunosAttr := "http://xml.juniper.net/junos/18.1R3/junos"
	if replyWithTablePtr.Junos != expectedJunosAttr {
		t.Errorf("Expected Junos attribute for RpcReplyWithTablePointer to be %s, but got %s", expectedJunosAttr, replyWithTablePtr.Junos)
	}

	// Example verification for RpcReply
	if originalReply.Junos != expectedJunosAttr {
		t.Errorf("Expected Junos attribute for RpcReply to be %s, but got %s", expectedJunosAttr, originalReply.Junos)
	}
}
