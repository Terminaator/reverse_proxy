@echo off 
setlocal enableextensions 

set REVERSE_PROXY_SERVER=127.0.0.1:8000
set REVERSE_PROXY_SERVER_REDIRECT_URL=http://127.0.0.1:7000
set CODE=^
package temp ?^
import "net/http" ?^
func Run(res *http.Response) { res.Header.Set("access-control-allow-origin", "http://localhost:3000") }

call go run proxy.go

endlocal