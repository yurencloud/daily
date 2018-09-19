package main

import (
	"daily/util"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func main() {
	isWeek := flag.Bool("week", false, "获取周报（过去7天日报汇总）")
	start := flag.Int("start", 1, "获取之前start天-end天之间的所有日报")
	end := flag.Int("end", 0, "获取之前start天-end天之间的所有日报")
	flag.Parse()

	result, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	config := util.Config{}
	json.Unmarshal(result, &config)

	var CN_NUMBER = [10]string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}

	layout := "Mon Jan 2"
	now := time.Now()
	var days []string
	if *isWeek {
		for i := 0; i < 7; i++ {
			ago, _ := time.ParseDuration("-" + strconv.Itoa(24*i) + "h")
			days = append(days, now.Add(ago).Format(layout))
		}
	}

	if *start != 0 {
		if *end > *start {
			log.Println("因为是过去的start天到end天, 所以start要大于end!")
			return
		}
		for i := *end; i < *start; i++ {
			ago, _ := time.ParseDuration("-" + strconv.Itoa(24*i) + "h")
			days = append(days, now.Add(ago).Format(layout))
		}
	}

	if !*isWeek && *start == 0 {
		ago, _ := time.ParseDuration("-24h")
		days = append(days, now.Add(ago).Format(layout))
	}

	var results string
	var daysString string
	for _, v := range days {
		daysString += v + "、"
	}

	results += "涉及日期：" + daysString + "\r\n"

	for i, v := range config.Repositories {
		results += "\r\n"+CN_NUMBER[i] + "、" + v.Title + "\r\n"
		f := 0
		for _, today := range days {
			logs, _ := util.Exec(v.Path + " && git log")
			res := util.Deal(config.Author, today, logs)
			for _, each := range res {
				f++
				results += strconv.Itoa(f) + "." + each + "\r\n"
			}
		}
	}

	ioutil.WriteFile("日报.txt", []byte(results), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
