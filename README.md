# printrouter

Simple library that allows to print information about http handlers for mux.Router:

ex: 

``[mode] [method] /[url]   --> [package][router].[handler] ([amount] handlers)``

``[debug] [POST  ] /test                     --> github.com/myProject/mypackage/myRouter.myHandler (1 handlers)``

### OVERVIEW

Default mode is 'debug' it means it will print all the routers if SetMode("notDebug") wasn't call.

## Usage

    export print "github.com/eu-ga/printrouter"
    
    r := mux.NewRouter().PathPrefix("/").Subrouter()
    
    router := print.NewRouter(r, "/")
    
Now **router** can be used as a regular mux.Router