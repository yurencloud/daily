package util

import (
	"regexp"
)

var result2 = `
commit 1c174701742051af0dd372912e58c7c7e9055a59
Author: mack-wang <641212003@qq.com>
Date:   Mon Sep 17 18:06:57 2018 +0800

    修改分时图数据和样式

commit fe9e68bd1c0cc4ca322537a8c86958623592acfd
Author: mack-wang <641212003@qq.com>
Date:   Mon Sep 17 16:45:44 2018 +0800

    提现地址添加修改和删除功能

commit c10e56faf207c0a7e688ddd752e9183b1ecd37d3
Author: mack-wang <641212003@qq.com>
Date:   Mon Sep 17 15:49:15 2018 +0800

    删除log

commit 749394698d6192b670b4ec4d641887c740b7a7e2
Author: mack-wang <641212003@qq.com>
Date:   Mon Sep 17 15:45:25 2018 +0800

    资产提现地址管理，修改和删除（未完成，仅测试）
`

func Deal(author string, today string, str string) (results []string) {
	records := Find(author+"[\\S\\s]+?"+today+"[\\S\\s]+?0800\\s+(?P<result>[\\S\\s]+?)\\s\\scommit", str)
	for _, record := range records {
		results = append(results, record[1])
	}
	return results
}

func Find(reg string, result string) [][]string {
	//"Server Hostname:\\s+(?P<result>[\\S\\s]+?)Server Port:")
	rege, _ := regexp.Compile(reg)
	return rege.FindAllStringSubmatch(result, -1)
}
