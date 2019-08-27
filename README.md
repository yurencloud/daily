#### 快速生成日报、周报、月报、季报
支持 windows、mac、linux

#### 下载使用


#### 更新日志
1. 去除Merge日志
2. 支持过滤删除日志前缀例如：`fix: `,`~`,`+`
3. 去除重复的日志
4. 去除空字符串日志

#### 使用方法
1. 输出日报
```
./daily
```

2. 输出周报
```
./daily -week
```

3. 输出月报
```
./daily -month
```

4. 输出季报
```
./daily -quarter
```

5. 输出指定日期的所有日报

输出第10天前到第5天前期间的所有日报，end默认为0
```
./daily -start 10 -end 5
```

#### 配置文件
Author：输入你的git username,会筛选出你所有的提交记录
Exclude: 去除每条日志的前缀，支持自定义
Repositories: 再输入项目切换路径即可
> mac或linux用户请删除config.json文件，并将config.linux.json重命名成config.json
```
{
  "Author": "mackwang",
  "Exclude": ["fix:", "feat:", "update:", "add:"],
  "Repositories": [
    {
      "Title": "商家中心",
      "Path": "cd /Users/wanglecheng/Documents/code/A"
    },
    {
      "Title": "运营后台",
      "Path": "cd /Users/wanglecheng/Documents/code/B"
    },
    {
      "Title": "自定义项目名称",
      "Path": "cd /Users/wanglecheng/Documents/code/C"
    }
  ]
}
```
> windows 下的`Path`要 类似这样 `E: && \\document\\code\\A`