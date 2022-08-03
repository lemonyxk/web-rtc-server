/**
* @program: web-rtc-server
*
* @description:
*
* @author: lemo
*
* @create: 2022-08-03 23:02
**/

package main

type Offer struct {
	Data string `json:"data"`
	To   string `json:"to"`
	From string `json:"from"`
}

type Answer struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}
