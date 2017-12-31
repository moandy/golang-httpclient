package main

import (
	"flag"
	"Agency"
)

var native bool

func main()  {
	flag.BoolVar(&native, "n", false, "Whether to use native method")
	flag.Parse()
	if native {
		customer := Agency.GetCustomerDetails()
		destinations := Agency.GetRecommendedDestinations(customer)
		var infos [10]Agency.Info
		for index, dest := range destinations {
			q := Agency.GetQuote(dest)
			w := Agency.GetWeatherForecast(dest)
			infos[index] = Agency.Info{Destinations:dest, Quote:q, Weather:w}
		}
		return infos
	} 
	else {
		customer := Agency.GetCustomerDetails()
		destinations := Agency.GetRecommendedDestinations(customer)
		var infos [10]Agency.Info
		quotes := [10]chan Agency.Quoting{}
		weathers := [10]chan Agency.Weather{}
		for i := range quotes {
			quotes[i] = make(chan Agency.Quoting)
		}
		for i := range weathers {
			weathers[i] = make(chan Agency.Weather)
		}
		for index, dest := range destinations {
			i := index
			d := dest
			go func() {
				quotes[i] <- Agency.GetQuote(d)
			}()
			go func() {
				weathers[i] <- Agency.GetWeatherForecast(d)
			}()
		}
		for index, dest := range destinations {
			infos[index] = Agency.Info{Destinations:dest, Quote:<-quotes[index], Weather:<-weathers[index]}
		}
		return infos
	}
}
