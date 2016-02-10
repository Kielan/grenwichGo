package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "grenwichGo"
	app.Usage = "Timezone converter"

	app.Action = func(c *cli.Context) {
		gmtFrom := os.Args[0]
		gmtTo := os.Args[2]
		timezoneConversion := timezoneMap[gmtFrom] + timezoneMap[gmtTo]

		var result string

		localTimeString := os.Args[2]

		if strings.ContainsAny(localTimeString, ":") {
			localTimeSlice := strings.Split(localTimeString, ":")

			localTimeFloat, err := strconv.ParseFloat(localTimeSlice[0], 64)
			if err != nil {
				log.Fatal(err)
			}

			conversion := localTimeFloat + timezoneConversion

			conversionString := strconv.FormatFloat(conversion, 'f', -1, 64)

			conversionString = strings.Join([]string{conversionString, localTimeSlice[1]}, ":")

			result = conversionString
		} else {
			conversion := timezoneConversion
			conversionString := strconv.FormatFloat(conversion, 'f', -1, 64)
			result = conversionString
		}

		fmt.Printf("%v to %v local time %v conversion diff \n", timezoneMap[gmtFrom], timezoneMap[gmtTo], result)
	}

	app.Run(os.Args)
}

var timezoneMap = map[string]float64{
	"gmt-12": -12,
	"gmt-11": -11,
	"gmt-10": -10,
	"gmt-9":  -9,
	"gmt-8":  -8,
	"gmt-7":  -7,
	"gmt-6":  -6,
	"gmt-5":  -5,
	"gmt-4":  -4,
	"gmt-3":  -3,
	"gmt-2":  -2,
	"gmt-1":  -1,
	"gmt0":   0,
	"gmt+1":  1,
	"gmt+2":  2,
	"gmt+3":  3,
	"gmt+4":  4,
	"gmt+5":  5,
	"gmt+6":  6,
	"gmt+7":  7,
	"gmt+8":  8,
	"gmt+9":  9,
	"gmt+10": 10,
	"gmt+11": 11,
	"gmt+12": 12,
}
