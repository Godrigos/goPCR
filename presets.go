package main

import (
	"encoding/json"
	"io/ioutil"
)

func stock() Solutions {
	var defaultValues Solutions

	//file, err := ioutil.ReadFile(execDir + "/presets/stock.json")
	file, err := ioutil.ReadFile("./presets/stock.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &defaultValues)
	if err != nil {
		panic(err)
	}

	return defaultValues
}

func pcr() Solutions {
	var pcrDefaultValues Solutions

	//file, err := ioutil.ReadFile(execDir + "/presets/pcr.json")
	file, err := ioutil.ReadFile("./presets/pcr.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &pcrDefaultValues)
	if err != nil {
		panic(err)
	}

	return pcrDefaultValues
}

func mix() Mix {
	var mixDefaultValues Mix

	//file, err := ioutil.ReadFile(execDir + "/presets/mix.json")
	file, err := ioutil.ReadFile("./presets/mix.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &mixDefaultValues)
	if err != nil {
		panic(err)
	}

	return mixDefaultValues
}
