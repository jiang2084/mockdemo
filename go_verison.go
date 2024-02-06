package main

import "github.com/jiang2084/mockdemo/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
