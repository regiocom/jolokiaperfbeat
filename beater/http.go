package beater

import (
    "net/http"
    "io/ioutil"
)

func GetStatsFromServer(basepath string) ([]byte, error) {
    url := basepath
    if url[len(url)-1] != '/' {
        url += "/"
    }
    url += "read/org.apache.cxf:bus.id=*,operation=\"*\",port=\"*\",service=\"*\",type=Performance.Counter.Server"
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return body, err
}
