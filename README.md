# proxy
A minimal http proxy in golang using gorilla router
#### quick start
The package is pretty self-explanatory, and hasn't complicated process.
#####install
```
go get github.com/ghoroubi/proxy
```   

```
package main
import (
    "github.com/ghoroubi/proxy"
    "gorilla/mux"
    "net/http"
    "context"
)

func main(){
// Declare router    
r:= mux.NewRouter().StrictSlash(true)
    
// Declare route
// In this example we supposed that all methods would be called
// So, all of http methods are invoked.
// Pay attention to the Path declartion
// You can use your favorite router package
// Instead of gorilla/mux, but the point is how to set the path
// The proxy server will redirect requests to the path that is after _/forward/_
// The Host of proxy is set in the context
// You can set any other values in context
// And handle it inside the proxy package  
ctx := context.WithValue(context.Background(), "url", &url.URL{
		Scheme: "http",
		Host:   "your_proxy_host:your_proxy_port", // remove [:port] if there isn't any
	})
r.Methods(http.MethodPost,
           http.MethodOptions,
           http.MethodPut,
           http.MethodDelete,
           http.MethodGet).
  Path("/forward/{rest:.*}").Handler(proxyHandler(ctx))
}
```

That's all ! you can add your other paths and just keep the /forward/ ( or your favorite prefix) for proxy. 
 
 #### Contribution
 Please add your comment and refactor any parts to make the package more flexible to others.