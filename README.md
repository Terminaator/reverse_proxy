# GO reverse proxy server for windows
run from CMD _**start.bat**_

## changing reverse proxy configuration
    set REVERSE_PROXY_SERVER=ADD_SERVER_IP
    set REVERSE_PROXY_SERVER_REDIRECT_URL=ADD_SERVER_REDIRECT_URL
    set CODE=ADD_RUNNABLE_CODE
    
## reverse proxy configuration example
    set REVERSE_PROXY_SERVER=127.0.0.1:8000
    set REVERSE_PROXY_SERVER_REDIRECT_URL=http://127.0.0.1:7000
    set CODE=^
    package temp ?^
    import "net/http" ?^
    func Run(res *http.Response) { res.Header.Set("access-control-allow-origin", "http://localhost:3000") }
    
## reverse proxy configuration example explanation
Inside CODE variable symbol _**?**_ will be replaced with linebreaks. \
Symbol **_^_** is used by bat file for defining linebreak. \
CODE variable must only contain three things to run reverse proxy successfully.

    //CODE is parsed and runned at runtime
    package temp //package name is hardcoded into proxy.go
    import "net/http" //this is needed for a Run function
    func Run(res *http.Response) {} //proxy.go needs this function at runtime. Also hardcoded.