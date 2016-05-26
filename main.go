package main


import (
    "log"
    "time"
    "github.com/jubyvictor/web4robot/device"
    "github.com/d2r2/go-i2c"
    "github.com/gorilla/websocket"
)

const(
    //I2cDeviceID is the hex addres of the weatherboard 0x4E
    I2cDeviceID = 0x4E
)

func main()  {
    log.Println("Starting web4robot client...")
    i2cDev, err := i2c.NewI2C(I2cDeviceID,1)
    if err != nil {
        log.Fatal(err)
    }
    weatherBoard := device.InitBoard(i2cDev)
    
    //Loop infinitely and make readings.
    for{
        reading := device.ReadChannels(weatherBoard)
        log.Println(reading)
        time.Sleep(1 * time.Second)
    }
    
    
}


