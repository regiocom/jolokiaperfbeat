package tests

import (
    "io/ioutil"
    "testing"
    "github.com/regiocom/jolokiaperfbeat/beater"
)

/*
 * Test CreateEvent for CodahaleMetricsProvider data
 */
func TestCreateEventTotals(t *testing.T) {
    dat, err := ioutil.ReadFile("data2.json")
    if err != nil {
        t.Fatal("tests/data2.json not found")
    }

    resp, err2 := beater.Convert(dat)
    if err2 != nil {
        t.Fatal(err2)
    }

    sd := beater.ServiceData{Source:"Metrics.Server"}
    key := "org.apache.cxf:Attribute=Totals,Operation=validate,bus.id=IdService-3.0.23-SNAPSHOT,port=\"IDServiceSOAP\",service=\"{http://bpo.regiocom.com/}IDServiceImplService\",type=Metrics.Server"
    event := beater.CreateEvent(1, "test", sd, resp.Counters[key])

    eventKey := "responsetime.p75"
    _, found := event[eventKey]
    if !found {
        t.Fatalf("Key %s not found", eventKey)
    }
}
