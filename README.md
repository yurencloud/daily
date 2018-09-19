## 快速生成日报、周报

支持 windows、mac、linux

## 使用方法

#### windows下可以直接点击daily.exe会在当前目录下生成"日报.txt"

#### 输出日报
./daily

#### 输出周报
./daily -week

#### 输出指定日期的所有日报
> 输出第10天前到第5天前期间的所有日报，end默认为0

./daily -start 10 -end 5


#### 配置文件
输入你的git username,会筛选出你所有的提交记录
再输入项目切换路径即可
> mac或linux用户请删除config.json文件，并将config.linux.json重命名成config.json