package main

import(
        "log"
        "net/url"
        "net/http"
        "net/http/httputil"
        "fmt"
        "strings"
        "os"
        "flag"
        "strconv"
)

type  revProxyAndUrl struct{
  url *url.URL
  proxy *httputil.ReverseProxy
  }

var locations map[string]revProxyAndUrl
func main() {
    flagPort       := flag.Int("port", 80, "default port to listen, 80 should be used with sudo")
    flagConfigFile := flag.String("config", "config.txt", "config file location and name")
    flag.Parse()
    
    locations       = make(map[string]revProxyAndUrl)
    configsFC, err := os.ReadFile(*flagConfigFile)
    if err != nil {
        fmt.Print("config.txt should include lines like from:port to:port ")
        panic(err)
    }
    
    configsFcS := strings.Split(string(configsFC), "\n")
    for _,confLine := range configsFcS{
      config := strings.Split(confLine, " ")
      if len(config)<2{
        continue
      }
      remote,_             :=  url.Parse(config[1])
      locations[config[0]]  =  revProxyAndUrl{remote, httputil.NewSingleHostReverseProxy(remote)}
      }
    

    handler := func() func(http.ResponseWriter, *http.Request) {
            return func(w http.ResponseWriter, r *http.Request) {
              var foundLocation   *httputil.ReverseProxy
              bFound := false
              host := string(r.Host)     
              log.Print(r.Host)       
              log.Println(r.URL)
              for l0,l1 := range locations{
               if l0 == host {
                    bFound = true
                    log.Print("Match found : ")
                    log.Println(l0)
                    log.Println(r)
                    r.Host = l1.url.Host
                    foundLocation = l1.proxy
                }
              }
              if !bFound {
                  w.Write([]byte(host + " host cant found") )
                }else{
                  log.Println()
                  w.Header().Set("X-Ben", "Rad")
                  foundLocation.ServeHTTP(w, r)
              }
              
            }
    }
    
    http.HandleFunc("/", handler())
    err = http.ListenAndServe(":"+strconv.Itoa(*flagPort), nil)
    if err != nil {
            panic(err)
    }
}
