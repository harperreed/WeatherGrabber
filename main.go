package main

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type RainwiseWeather struct {
	DeviceTime    string  `json:"time"`
	BatteryLevel  float64 `json:"batt"`
	Signal        int     `json:"signal"`
	SignalQuality int     `json:"quality"`
	US            struct {
		AirTemperature struct {
			Current           float64 `json:"tic"`
			Average           float64 `json:"tia"`
			TodayHigh         float64 `json:"tdh"`
			TodayIntervalHigh float64 `json:"tih"`
			TodayLow          float64 `json:"tdl"`
			TodayIntervalLow  float64 `json:"til"`
		} `json:"atmp"`
		RelativeHumidity struct {
			Current           int `json:"ric"`
			Average           int `json:"ria"`
			TodayHigh         int `json:"rdh"`
			TodayIntervalHigh int `json:"rih"`
			TodayLow          int `json:"rdl"`
			TodayIntervalLow  int `json:"ril"`
		} `json:"rh"`
		BarometricPressure struct {
			Current           float64 `json:"bic"`
			Average           float64 `json:"bia"`
			TodayHigh         float64 `json:"bdh"`
			TodayIntervalHigh float64 `json:"bih"`
			TodayLow          float64 `json:"bdl"`
			TodayIntervalLow  float64 `json:"bil"`
		} `json:"bp"`
		Wind struct {
			Current   float64 `json:"wic"`
			Wict      int     `json:"wict"`
			Average   float64 `json:"wia"`
			TodayHigh float64 `json:"wdh"`
			Wdht      int     `json:"wdht"`
			Wih       float64 `json:"wih"`
			Wiht      int     `json:"wiht"`
		} `json:"wnd"`
		Rainfall struct {
			Today  float64 `json:"rfd"`
			Minute float64 `json:"rfm"`
		} `json:"rf"`
		/*
			SolarRad struct {
				Src int `json:"src"`
				Srd int `json:"srd"`
				Srm int `json:"srm"`
			} `json:"sr"`
			SolarRad2 struct {
				Sr2C int `json:"sr2c"`
				Sr2D int `json:"sr2d"`
				Sr2M int `json:"sr2m"`
			} `json:"sr2"`
			UVIndex struct {
				Uvc int `json:"uvc"`
				Uvd int `json:"uvd"`
				Uvm int `json:"uvm"`
			} `json:"uv"`
			LeafWetness struct {
				Lwc int `json:"lwc"`
				Lwd int `json:"lwd"`
				Lwm int `json:"lwm"`
			} `json:"lw"`
			Temperature1 struct {
				T1Ic float64 `json:"t1ic"`
				T1Ia float64 `json:"t1ia"`
				T1Dh float64 `json:"t1dh"`
				T1Ih float64 `json:"t1ih"`
				T1Dl float64 `json:"t1dl"`
				T1Il float64 `json:"t1il"`
			} `json:"tmp1"`
			Temperature2 struct {
				T2Ic float64 `json:"t2ic"`
				T2Ia float64 `json:"t2ia"`
				T2Dh float64 `json:"t2dh"`
				T2Ih float64 `json:"t2ih"`
				T2Dl float64 `json:"t2dl"`
				T2Il float64 `json:"t2il"`
			} `json:"tmp2"`
			SoilMoisture struct {
				Sic int `json:"sic"`
				Sia int `json:"sia"`
				Sdh int `json:"sdh"`
				Sih int `json:"sih"`
				Sdl int `json:"sdl"`
				Sil int `json:"sil"`
			} `json:"sm"`
			InsideTemperature struct {
				Itic float64 `json:"itic"`
				Itia float64 `json:"itia"`
				Itdh float64 `json:"itdh"`
				Itih float64 `json:"itih"`
				Itdl float64 `json:"itdl"`
				Itil float64 `json:"itil"`
			} `json:"itmp"`
		*/
	} `json:"us"`
	/*
		Metric struct {
			Atmp struct {
				Tic float64 `json:"tic"`
				Tia float64 `json:"tia"`
				Tdh float64 `json:"tdh"`
				Tih float64 `json:"tih"`
				Tdl float64 `json:"tdl"`
				Til float64 `json:"til"`
			} `json:"atmp"`
			Rh struct {
				Ric int `json:"ric"`
				Ria int `json:"ria"`
				Rdh int `json:"rdh"`
				Rih int `json:"rih"`
				Rdl int `json:"rdl"`
				Ril int `json:"ril"`
			} `json:"rh"`
			Bp struct {
				Bic float64 `json:"bic"`
				Bia float64 `json:"bia"`
				Bdh float64 `json:"bdh"`
				Bih float64 `json:"bih"`
				Bdl float64 `json:"bdl"`
				Bil float64 `json:"bil"`
			} `json:"bp"`
			Wnd struct {
				Wic  float64 `json:"wic"`
				Wict int     `json:"wict"`
				Wia  float64 `json:"wia"`
				Wdh  float64 `json:"wdh"`
				Wdht int     `json:"wdht"`
				Wih  float64 `json:"wih"`
				Wiht int     `json:"wiht"`
			} `json:"wnd"`
			Rf struct {
				Rfd float64 `json:"rfd"`
				Rfm float64 `json:"rfm"`
			} `json:"rf"`
			Sr struct {
				Src int `json:"src"`
				Srd int `json:"srd"`
				Srm int `json:"srm"`
			} `json:"sr"`
			Sr2 struct {
				Sr2C int `json:"sr2c"`
				Sr2D int `json:"sr2d"`
				Sr2M int `json:"sr2m"`
			} `json:"sr2"`
			Uv struct {
				Uvc int `json:"uvc"`
				Uvd int `json:"uvd"`
				Uvm int `json:"uvm"`
			} `json:"uv"`
			Lw struct {
				Lwc int `json:"lwc"`
				Lwd int `json:"lwd"`
				Lwm int `json:"lwm"`
			} `json:"lw"`
			Tmp1 struct {
				T1Ic float64 `json:"t1ic"`
				T1Ia float64 `json:"t1ia"`
				T1Dh float64 `json:"t1dh"`
				T1Ih float64 `json:"t1ih"`
				T1Dl float64 `json:"t1dl"`
				T1Il float64 `json:"t1il"`
			} `json:"tmp1"`
			Tmp2 struct {
				T2Ic float64 `json:"t2ic"`
				T2Ia float64 `json:"t2ia"`
				T2Dh float64 `json:"t2dh"`
				T2Ih float64 `json:"t2ih"`
				T2Dl float64 `json:"t2dl"`
				T2Il float64 `json:"t2il"`
			} `json:"tmp2"`
			Sm struct {
				Sic int `json:"sic"`
				Sia int `json:"sia"`
				Sdh int `json:"sdh"`
				Sih int `json:"sih"`
				Sdl int `json:"sdl"`
				Sil int `json:"sil"`
			} `json:"sm"`
			Itmp struct {
				Itic float64 `json:"itic"`
				Itia float64 `json:"itia"`
				Itdh float64 `json:"itdh"`
				Itih float64 `json:"itih"`
				Itdl float64 `json:"itdl"`
				Itil float64 `json:"itil"`
			} `json:"itmp"`
		} `json:"metric"`
	*/
}

type Weather struct {
	DeviceTime    string  `json:"time"`
	BatteryLevel  float64 `json:"batt"`
	Signal        int     `json:"signal"`
	SignalQuality int     `json:"quality"`
	US            struct {
		AirTemperature struct {
			Current           float64 `json:"tic"`
			Average           float64 `json:"tia"`
			TodayHigh         float64 `json:"tdh"`
			TodayIntervalHigh float64 `json:"tih"`
			TodayLow          float64 `json:"tdl"`
			TodayIntervalLow  float64 `json:"til"`
		} `json:"atmp"`
		RelativeHumidity struct {
			Current           int `json:"ric"`
			Average           int `json:"ria"`
			TodayHigh         int `json:"rdh"`
			TodayIntervalHigh int `json:"rih"`
			TodayLow          int `json:"rdl"`
			TodayIntervalLow  int `json:"ril"`
		} `json:"rh"`
		BarometricPressure struct {
			Current           float64 `json:"bic"`
			Average           float64 `json:"bia"`
			TodayHigh         float64 `json:"bdh"`
			TodayIntervalHigh float64 `json:"bih"`
			TodayLow          float64 `json:"bdl"`
			TodayIntervalLow  float64 `json:"bil"`
		} `json:"bp"`
		Wind struct {
			Current   float64 `json:"wic"`
			Wict      int     `json:"wict"`
			Average   float64 `json:"wia"`
			TodayHigh float64 `json:"wdh"`
			Wdht      int     `json:"wdht"`
			Wih       float64 `json:"wih"`
			Wiht      int     `json:"wiht"`
		} `json:"wnd"`
		Rainfall struct {
			Today  float64 `json:"rfd"`
			Minute float64 `json:"rfm"`
		} `json:"rf"`
	} `json:"us"`
}

func (w *Weather) MarshalJSON() ([]byte, error) {
	type WeatherAlias struct {
		DeviceTime    string  `json:"deviceTime"`
		BatteryLevel  float64 `json:"batteryLevel"`
		Signal        int     `json:"signal"`
		SignalQuality int     `json:"signalQuality"`
		US            struct {
			AirTemperature struct {
				Current           float64 `json:"current"`
				Average           float64 `json:"average"`
				TodayHigh         float64 `json:"todayHigh"`
				TodayIntervalHigh float64 `json:"todayIntervalHigh"`
				TodayLow          float64 `json:"todayLow"`
				TodayIntervalLow  float64 `json:"todayIntervalLow"`
			} `json:"airTemperature"`
			RelativeHumidity struct {
				Current           int `json:"current"`
				Average           int `json:"average"`
				TodayHigh         int `json:"todayHigh"`
				TodayIntervalHigh int `json:"todayIntervalHigh"`
				TodayLow          int `json:"todayLow"`
				TodayIntervalLow  int `json:"todayIntervalLow"`
			} `json:"relativeHumidity"`
			BarometricPressure struct {
				Current           float64 `json:"current"`
				Average           float64 `json:"average"`
				TodayHigh         float64 `json:"todayHigh"`
				TodayIntervalHigh float64 `json:"todayIntervalHigh"`
				TodayLow          float64 `json:"todayLow"`
				TodayIntervalLow  float64 `json:"todayIntervalLow"`
			} `json:"barometricPressure"`
			Wind struct {
				Current   float64 `json:"current"`
				Wict      int     `json:"currentDirection"`
				Average   float64 `json:"average"`
				TodayHigh float64 `json:"todayHigh"`
				Wdht      int     `json:"todayHighDirection"`
				Wih       float64 `json:"intervalHigh"`
				Wiht      int     `json:"intervalHighDirection"`
			} `json:"wind"`
			Rainfall struct {
				Today  float64 `json:"today"`
				Minute float64 `json:"interval"`
			} `json:"rainfall"`
		} `json:"US"`
	}
	var a WeatherAlias = WeatherAlias(*w)
	return json.Marshal(&a)
}

func main() {

	//set up logging
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[WeatherGrabber] ")
	log.Println("Starting up...")

	viper.SetConfigName("weathergrabber-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	ip100_ip := viper.GetString("ip100ip")
	ip100_url := "http://" + ip100_ip + "/weather.json"

	log.Println("Grabbing weather ")

	rsp, err := http.Get(ip100_url)

	defer rsp.Body.Close()

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// Fill the record with the data from the JSON
	var weather Weather

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(rsp.Body).Decode(&weather); err != nil {
		log.Println(err)
	}

	log.Println("Pushing weather to webhook ")
	jsonStr, err := json.Marshal(&weather)

	url := viper.GetString("webhook")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("api_key", viper.GetString("api_key"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
