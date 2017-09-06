package main

// This file is mandatory as otherwise the jolokiaperfbeat.test binary is not generated correctly.

import (
	"flag"
	"testing"
        "github.com/ErikWegner/jolokiaperfbeat/beater"
        "io/ioutil"
//        "fmt"
)

var systemTest *bool

func init() {
	systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")
}

// Test started when the test binary is started. Only calls main.
func TestSystem(t *testing.T) {

	if *systemTest {
		main()
	}
}

func TestConvert(t *testing.T) {
    dat, err := ioutil.ReadFile("tests/data1.txt")
    if err != nil {
        t.Fatal("tests/data1.txt not found")
    }

    resp, err2 := beater.Convert(dat)
    if err2 != nil {
        t.Fatal(err2)
    }

    actualLength := len(resp.Counters)
    expected := 26
    if actualLength != expected {
        t.Fatalf("Expected %d, actual %d", expected, actualLength)
    }
}

func TestServiceDataExctract(t *testing.T) {
    s := "org.apache.cxf:bus.id=BankcheckService-3.1.7,operation=\"validate\",port=\"BankcheckServiceSOAP\",service=\"{http://regiocom.com/}BankcheckService\",type=Performance.Counter.Server"
    r := beater.ServiceDataExtract(s)
    if r.ServiceName != "BankcheckService" {
        t.Fatalf("Service mismatch: %s", r.ServiceName)
    }
    if r.Version != "3.1.7" {
        t.Fatalf("Version mismatch: %s", r.Version)
    }
    if r.Operation != "validate" {
        t.Fatalf("Operation mismatch: %s", r.Operation)
    }
}