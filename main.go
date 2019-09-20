package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func getData(baseUrl string, item string) {
	// 创建各个省份的csv文件
	fName := item + ".csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	if writer.Write([]string{"bank_no", "bank_name", "mobile", "zip_code", "address", "swift_code", "bank_province"}) != nil {
		log.Fatal(err)
	}

	page := 1
	// TODO 先写死最多只请求400页
	for page < 400 {
		st := strconv.Itoa(page)
		reqUrl := baseUrl + "/" + string(st)

		// Instantiate default collector
		c := colly.NewCollector()

		// 异步监听
		c.OnHTML("table tr", func(e *colly.HTMLElement) {
			var tdArr []string

			e.ForEach("td", func(_ int, elem *colly.HTMLElement) {
				tdArr = append(tdArr, elem.Text)
			})

			if tdArr != nil {
				tdArr = append(tdArr, item)

				err := writer.Write(tdArr)
				if err != nil {
					log.Fatal(err)
				}
			}
		})

		fmt.Println("请求", reqUrl)
		page++
		err := c.Visit(reqUrl)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("请求", item, "结束")
}

func main() {
	connStr := "user=default dbname=root password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	id := 1
	_, err = db.Query("SELECT * FROM nice_bank_info WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	provinces := map[string]string{
		"jiangsu":      "江苏",
		"guangdong":    "广东",
		"shandong":     "山东",
		"hebei":        "河北",
		"zhejiang":     "浙江",
		"fujian":       "福建",
		"liaoning":     "辽宁",
		"anhui":        "安徽",
		"hubei":        "湖北",
		"sichuan":      "四川",
		"shanxisheng":  "陕西",
		"hunan":        "湖南",
		"shanxi":       "山西",
		"guizhou":      "贵州",
		"henan":        "河南",
		"heilongjiang": "黑龙江",
		"jilin":        "吉林",
		"xinjiang":     "新疆",
		"shanghai":     "上海",
		"gansu":        "甘肃",
		"yunnan":       "云南",
		"beijing":      "北京",
		"neimenggu":    "内蒙古",
		"tianjin":      "天津",
		"jiangxi":      "江西",
		"chongqing":    "重庆",
		"guangxi":      "广西",
		"ningxia":      "宁夏",
		"hainan":       "海南",
		"qinghai":      "青海",
		"xianggang":    "香港",
		"xicang":       "西藏",
		"aomen":        "澳门",
	}

	argsCount := len(os.Args)

	if argsCount <= 1 {
		fmt.Println("至少输入一个参数：pachou or ruku")
		return
	}

	switch os.Args[1] {
	case "pachon": // 爬虫
		var baseUrl string
		for key, item := range provinces {
			baseUrl = "http://5cm.cn/bank/" + key
			getData(baseUrl, item)
		}
	case "ruku": // 入库
		fmt.Println("读取csvStart---------")
		for _, item := range provinces {
			fName := item + ".csv"
			file, err := os.Open(fName)
			if err != nil {
				log.Fatal(err)
			}

			reader := csv.NewReader(file)
			rows, _ := reader.ReadAll()

			fmt.Println("入库", fName, "start-----")
			sqlstr := ""
			for k, v := range rows {
				if k == 0 {
					continue
				}

				sqlstr += "INSERT INTO nice_bank_info(bank_no,bank_name,mobile,zip_code,address,bank_province) " +
					"VALUES('" + v[0] + "','" + v[1] + "','" + v[2] + "','" + v[3] + "','" + v[4] + "','" + v[6] + "');"
			}

			_, err = db.Exec(sqlstr)
			if err != nil {
				log.Fatal(sqlstr, err)
			}

			fmt.Print()
			fmt.Println("入库", "end-----")
		}
	}

}
