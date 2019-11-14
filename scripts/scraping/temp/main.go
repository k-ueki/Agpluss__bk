package main

import (
	"fmt"
	"strings"
)

func main() {
	// url1 := "http://syllabus.aoyama.ac.jp/?BU=BU1&CP4=on&DL=ja&GKB=&JG1=on&JG2=on&JG3=on&JG4=on&JG5=on&JG6=on&JG7=on&KI=&KM=&KW=&PC=123&PG=3&PI=0&ST=&YB1=on&YB2=on&YB3=on&YB4=on&YB5=on&YB6=on&YR=2019&__EVENTARGUMENT=&__EVENTTARGET=ctl00%2524CPH1%2524rptPagerT%2524ctl03%2524lnkPage&__VIEWSTATEGENERATOR=309A73F1"
	url1 := "http://syllabus.aoyama.ac.jp/?BU=BU1&CP4=on&DL=ja&GKB=&JG1=on&JG2=on&JG3=on&JG4=on&JG5=on&JG6=on&JG7=on&KI=&KM=&KW=&PC=123&PG=3&PI=0&ST=&YB1=on&YB2=on&YB3=on&YB4=on&YB5=on&YB6=on&YR=2019&__EVENTARGUMENT=&__VIEWSTATEGENERATOR=309A73F1?__EVENTTARGET=ctl00%24CPH1%24rptPagerT%24ctl03%24lnkPage"

	url2 := "http: //syllabus.aoyama.ac.jp/?__EVENTTARGET=ctl00%24CPH1%24rptPagerT%24ctl03%24lnkPage&__EVENTARGUMENT=&__VIEWSTATEGENERATOR=309A73F1&YR=2019&BU=BU1&KW=&KM=&KI=&CP4=on&YB1=on&YB2=on&YB3=on&YB4=on&YB5=on&YB6=on&JG1=on&JG2=on&JG3=on&JG4=on&JG5=on&JG6=on&JG7=on&GB1B_0=&GKB=&DL=ja&ST=&PG=3&PC=123&PI=0"

	endpoint1 := strings.Split(url1, "?")
	endpoint2 := strings.Split(url2, "?")

	factors1 := strings.Split(endpoint1[1], "&")
	factors2 := strings.Split(endpoint2[1], "&")

	for _, v := range factors1 {
		fmt.Println("1 : ", v)
		for i, p := range factors2 {
			fmt.Println("2 : ", p)
			if v == p {
				fmt.Println("break!")
				break
			}
			if i == len(factors2)-1 {
				fmt.Println("!!!!!!! ", v, p, " !!!!!!!!")
			}
		}
	}
}
