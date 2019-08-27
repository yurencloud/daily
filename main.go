package main

import (
	"daily/util"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func exclude(item string, rules []string) string {
	var result string
	for _, v := range rules {
		if strings.Contains(item, v) {
			result = strings.Replace(item, v, "", 1)
		}
	}
	return result
}

func arrayIsContain(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	isWeek := flag.Bool("week", false, "获取周报（过去7天日报汇总）")
	isMonth := flag.Bool("month", false, "获取月报（过去30天日报汇总）")
	isQuarter := flag.Bool("quarter", false, "获取季报（过去30天日报汇总）")
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
	var title = "日报.txt"

	if *isWeek {
		title = "周报.txt"
		*start = 7
	}
	if *isMonth {
		title = "月报.txt"
		*start = 30
	}

	if *isQuarter {
		title = "季报.txt"
		*start = 90
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

	var results string
	var daysString string
	for _, v := range days {
		daysString += v + "、"
	}

	results += "涉及日期：" + daysString + "\r\n"

	for i, v := range config.Repositories {
		results += "\r\n" + CN_NUMBER[i] + "、" + v.Title + "\r\n"
		f := 0
		var list []string
		for _, today := range days {
			logs, _ := util.Exec(v.Path + " && git log")
			res := util.Deal(config.Author, today, logs)
			for _, each := range res {
				if !strings.Contains(each, "Merge") {
					var logs = strings.Trim(exclude(each, config.Exclude), "")
					// 过滤空日志
					if len(logs) != 0 && !arrayIsContain(list, logs) {
						// 日志去重
						list = append(list, logs)
						f++
					}
				}
			}

		}
		for e, v := range list {
			results += strconv.Itoa(e+1) + "." + v + "\r\n"
		}
	}

	fmt.Println(results)

	ioutil.WriteFile(title, []byte(results), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
