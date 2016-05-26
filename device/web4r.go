package device

import(
    "github.com/d2r2/go-i2c"
    "time"
    "log"
)

const(
    boardAddress = 0x4E
    externalTempReg = 0x42
    humidityReg = 0x43
    pressureReg = 0x47
)

//WeatherBoard method, exposes a Wrapper around the weather board
type WeatherBoard struct {
    i2cdev *i2c.I2C
}


//Reading represents a reading from the device.
type Reading struct{
    deviceID string
    ts int64
    internalTemp float32
    humidity float32
    pressure float32
    
}

//InitBoard Initializes the weatherboard with the connection to the i2c device. 
func InitBoard(i2c *i2c.I2C) (*WeatherBoard){
    weatherBoard := WeatherBoard{i2c}
    return &weatherBoard
}

//ReadChannels returns a reading from the weatherboard
func ReadChannels(board *WeatherBoard, ) Reading{
    
    readStart := time.Now().UTC().Unix()
    
    tempInside, err := board.i2cdev.ReadRegU8(externalTempReg) //oC    
    humidity, err := board.i2cdev.ReadRegU8(humidityReg) //%
    pressure, err := board.i2cdev.ReadRegU8(pressureReg)
    adjustedPressure  :=  0.1 * float32(pressure) //kPa
    
    readEnd := time.Now().UTC().Unix()
    
    //Time stamp is average of start and end to be simple.    
    reading := Reading{ deviceID:"x", 
        ts: (readEnd + readStart)/2, 
        internalTemp: float32(tempInside), 
        humidity: float32(humidity), 
        pressure: adjustedPressure}
         
    return reading
}