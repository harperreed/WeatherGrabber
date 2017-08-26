# WeatherGrabber
A cron job that will regularly grab your RainWise IP100 weather.json file and post it to a webhook 

Just build the executable:

`go get`  
`go build`

and then fill out the details in the config file

Then run it!

    harper@machine {~/src/go/src/weather}$ ./WeatherGrabber                                                                                                                                           
    [WeatherGrabber] 2017/08/26 14:32:13 Starting up...
    [WeatherGrabber] 2017/08/26 14:32:13 Grabbing weather
    [WeatherGrabber] 2017/08/26 14:32:14 Pushing weather to webhook
    harper@machine {~/src/go/src/weather}$

Should work

## Why?

I am using this to offload my weather data to an app to show historical weather information about my PWS. This is pretty straight forward. 

I run this every 5 min via a local cron. WFM ;)
