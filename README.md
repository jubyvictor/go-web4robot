WEB4ROBOT
=========
WEB4ROBOT allows you to interface with the now out of market weatherboard for raspberry pi that was developed by web4robot.

The weatherboard is a I2C device with address 0x4E, and usually shows up as /dev/i2c-1. Please make sure that the linux modules i2c_dev and i2c_bcm2708 are loaded.

### How it works
The application uses the config file that is passed in as a command line argument and polls the device at the specified interval (poll_interval) in the config file. Each polled reading is posted as a data point using a websocket. Modify the upstream_url in the config file to change the target endpoint.


### Usage

web4robot /path/to/web4robot.json


### Credits

The application uses:

GO I2C module from 
https://github.com/d2r2/go-i2c

Gorillakit websocket from 
https://github.com/gorilla/websocket
   
