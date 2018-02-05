package beater

import (
    "net/http"
    "io/ioutil"
    "github.com/regiocom/jolokiaperfbeat/config"
)

func GetStatsFromServer(config config.Config) ([]byte, error) {
    url := config.Baseurl
    if url[len(url)-1] != '/' {
        url += "/"
    }
    switch config.Provider {
    case "cxfCounter":
        url += "read/org.apache.cxf:bus.id=*,operation=\"*\",port=\"*\",service=\"*\",type=Performance.Counter.Server"
    case "cxfMetrics":
        url += "org.apache.cxf:bus.id=*,type=Metrics.Server,service=*,port=*,Operation=*,Attribute=*"
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return body, err
}
