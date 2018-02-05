package main

// This file is mandatory as otherwise the jolokiaperfbeat.test binary is not generated correctly.

import (
    "flag"
    "testing"
    "github.com/regiocom/jolokiaperfbeat/beater"
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

func TestServiceDataExctract2(t *testing.T) {
    s := "org.apache.cxf:bus.id=BankcheckService-3.1.7-hotfix-1,operation=\"validate\",port=\"BankcheckServiceSOAP\",service=\"{http://regiocom.com/}BankcheckService\",type=Performance.Counter.Server"
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

/*
 * Test CodahaleMetricsProvider data conversion
 */
func TestConvert2(t *testing.T) {
    dat, err := ioutil.ReadFile("tests/data2.json")
    if err != nil {
        t.Fatal("tests/data2.json not found")
    }

    resp, err2 := beater.Convert(dat)
    if err2 != nil {
        t.Fatal(err2)
    }

    actualLength := len(resp.Counters)
    expected := 40
    if actualLength != expected {
        t.Fatalf("Expected %d, actual %d", expected, actualLength)
    }

    key := "org.apache.cxf:Attribute=Totals,Operation=validate,bus.id=IdService-3.0.23-SNAPSHOT,port=\"IDServiceSOAP\",service=\"{http://bpo.regiocom.com/}IDServiceImplService\",type=Metrics.Server"
    perfCounter, found := resp.Counters[key]
    if !found {
        t.Fatalf("Key %s not found", key)
    }

    if perfCounter.NumInvocations != 0 {
        t.Fatalf("NumInvocations must be zero")
    }

    p99 := 77.85287699999999
    if perfCounter.Percentile99th != p99 {
        t.Fatalf("Percentile99th: expected f, actual %f", p99, perfCounter.Percentile99th)
    }
}
