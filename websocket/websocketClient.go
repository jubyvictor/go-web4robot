package websocket

import(
    "log"
    "golang.org/x/net/websocket"
)
//ServerConfig defines the config necessary to connect to the upstream server.
type ServerConfig struct{
    url string
    protocol string    
    origin string    
}

//Init initializes the websocket connection
func Init(config *ServerConfig) *websocket.Conn{    
    ws, err := websocket.Dial(config.url,config.protocol,config.origin)
    if err != nil {
        log.Printf("WARNING: Could not connect to url %s",config.url)
    }   
    return ws    
}

//Destroy closes the websocket connection
func Destroy(websocket *websocket.Conn){
    
}



