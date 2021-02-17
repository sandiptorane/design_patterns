package facade

import (
	"fmt"
	"testing"
)

func TestOpenWeather_ResponseParser(t *testing.T){
	r := GetMockData()
	openWeatherMap := CurrentWeatherData{
		APIKey: "",
	}
    weather, err := openWeatherMap.responseParser(r)

    if err != nil{
    	t.Fatal(err)
	}

	if weather.ID==54875{
		t.Errorf("Madrid id is 54875, not %d\n", weather.ID)
	}

	fmt.Println(weather.Main.Temp)
}
