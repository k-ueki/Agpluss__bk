package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var endpoint string = "http://syllabus.aoyama.ac.jp/"

func main() {

	_, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	etarget := "ctl00%24CPH1%24rptPagerT%24ctl03%24lnkPage"

	factor := make([][2]string, 30)
	factor[0] = [2]string{"__EVENTTARGET", etarget}
	factor[1] = [2]string{"__EVENTARGUMENT", ""}
	factor[2] = [2]string{"__VIEWSTATEGENERATOR", "309A73F1"}
	factor[3] = [2]string{"YR", "2019"}
	factor[4] = [2]string{"BU", "BU1"}
	factor[5] = [2]string{"KW", ""}
	factor[6] = [2]string{"KM", ""}
	factor[7] = [2]string{"KI", ""}
	factor[8] = [2]string{"CP4", "on"}
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
	factor[27] = [2]string{"PG", "3"}
	factor[28] = [2]string{"PC", "123"}
	factor[29] = [2]string{"PI", "0"}

	for i, v := range factor {
		if i == 0 {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += fmt.Sprintf("%s=%s", v[0], v[1])
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
