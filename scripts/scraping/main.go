package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Classes struct {
	Name        string `db:"name"`
	Semester    string `db:"semester"`
	Credits     int    `db:"credits"`
	Teacher     string `db:"teacher"`
	Description string `db:"description"`
	Year        int    `db:"year"`
	DayAt       string `db:"day_at"`
	Campus      string `db:"campus"`
	// Prerequisite string
	// Evaluation string
}

var endpoint string = "http://syllabus.aoyama.ac.jp/"

func gormConnect() *gorm.DB {
	CONNECT := "root:@/agpluss"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	_, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	etarget := "ctl00%24CPH1%24rptPagerT%24ctl03%24lnkPage"

	count := 0

	for i := 1; i <= 395; i++ {
		url := ""
		page := strconv.Itoa(i)

		factor := make([][2]string, 30)
		factor[0] = [2]string{"__EVENTTARGET", etarget}
		factor[1] = [2]string{"__EVENTARGUMENT", ""}
		factor[2] = [2]string{"__VIEWSTATEGENERATOR", "309A73F1"}
		factor[3] = [2]string{"YR", "2019"}
		factor[4] = [2]string{"BU", "BU1"}
		factor[5] = [2]string{"KW", ""}
		factor[6] = [2]string{"KM", ""}
		factor[7] = [2]string{"KI", ""}
		factor[8] = [2]string{"CP1", "on"}
		// factor[8] = [2]string{"CP4", "on"}
		factor[9] = [2]string{"YB1", "on"}
		factor[10] = [2]string{"YB1", "on"}
		factor[11] = [2]string{"YB2", "on"}
		factor[12] = [2]string{"YB3", "on"}
		factor[13] = [2]string{"YB4", "on"}
		factor[14] = [2]string{"YB5", "on"}
		factor[15] = [2]string{"YB6", "on"}
		factor[16] = [2]string{"JG1", "on"}
		factor[17] = [2]string{"JG2", "on"}
		factor[18] = [2]string{"JG3", "on"}
		factor[19] = [2]string{"JG4", "on"}
		factor[20] = [2]string{"JG5", "on"}
		factor[21] = [2]string{"JG6", "on"}
		factor[22] = [2]string{"JG7", "on"}
		factor[23] = [2]string{"GB1B_0", ""}
		factor[24] = [2]string{"GKB", ""}
		factor[25] = [2]string{"DL", "ja"}
		factor[26] = [2]string{"ST", ""}
		factor[27] = [2]string{"PG", page}
		factor[28] = [2]string{"PC", "395"}
		// factor[29] = [2]string{"PI", page_i}

		for i, v := range factor {
			if i == 0 {
				url += "?"
			} else {
				url += "&"
			}
			url += fmt.Sprintf("%s=%s", v[0], v[1])
		}

		resp, err := http.Get(endpoint + url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		selection := doc.Find("table.result > tbody > tr")

		// href, _ := selection.Attr("href")
		hrefs := [20]string{}
		dayats := [20]string{}
		campus := [20]string{}
		temp2 := [10]string{}

		selection.Each(func(index int, s *goquery.Selection) {
			sel := s.Find("td.col2 > span#CPH1_gvw_kensaku_lblJigen_0 > span")
			sel.Each(func(i int, se *goquery.Selection) {
				temp2[i] = se.Text()
			})

			tmp, _ := s.Find("td.col8 > a").Attr("href")
			hrefs[index] = tmp
			dayats[index] = temp2[1]
			campus[index] = temp2[0]
		})

		// fmt.Println(temp2)
		// fmt.Println(hrefs)
		// fmt.Println(dayats)
		// fmt.Println(campus)
		//
		// return

		/////詳細のスクレイピング
		for i, href := range hrefs {
			if href == "" {
				break
			}
			res, err := http.Get(endpoint + href)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			doc2, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			selection2 := doc2.Find("table.editTable > tbody > tr > td")

			temp := [150]string{}
			selection2.Each(func(index int, s *goquery.Selection) {
				temp[index] = s.Text()
			})

			// bytes, _ := ioutil.ReadAll(res.Body)
			// fmt.Println("nya-nn", string(bytes))

			credits, _ := strconv.Atoi(temp[4])
			year, _ := strconv.Atoi(temp[0])
			class := Classes{
				Name:        temp[1],
				Semester:    temp[3],
				Credits:     credits,
				Teacher:     temp[5],
				Description: temp[7],
				Year:        year,
				DayAt:       dayats[i],
				Campus:      campus[i],
			}
			count++

			fmt.Printf("%#v\n", class)
			fmt.Println(count)

			db.Create(&class)
		}
	}
}
