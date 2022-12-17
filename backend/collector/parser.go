package collector

import (
	"backend/entity"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strings"
	"time"
)

func ParseHtmlToClassObjects(htmls string) []entity.Class {
	var (
		trTags  []*html.Node
		classes []entity.Class
	)

	root, err := html.Parse(strings.NewReader(htmls))
	if err != nil {
		log.Fatalln(err)
	}

	parseTableTags(root, &trTags)

	for _, tr := range trTags {
		var tdTags []*html.Node
		var td *html.Node

		parseTrTags(tr, &tdTags)
		td = tdTags[2]
		tdTags = append(tdTags[:1], tdTags[3:]...)

		var otherFields []string
		parseTdTags(tdTags, &otherFields)

		var rawDataOfScheduleField []string
		exactDataFromTdTag(td, &rawDataOfScheduleField)
		scheduleField := getLessonByDay(rawDataOfScheduleField)

		subjectField := strings.Split(otherFields[0], "-")[0]
		c := entity.NewClass(otherFields[0], subjectField, otherFields[1], otherFields[2], scheduleField)
		classes = append(classes, *c)
	}
	return classes
}

func parseTableTags(root *html.Node, trTags *[]*html.Node) {
	if root.Type == html.ElementNode && root.Data == "tr" {
		for _, attr := range root.Attr {
			if attr.Val == "cssListAlternativeItem" || attr.Val == "cssListItem" {
				*trTags = append(*trTags, root)
				return
			}
		}
	}

	for n := root.FirstChild; n != nil; n = n.NextSibling {
		parseTableTags(n, trTags)
	}
}

func parseTrTags(root *html.Node, tdTags *[]*html.Node) {
	if root.Type == html.ElementNode && root.Data == "td" {
		for _, attr := range root.Attr {
			if attr.Key == "align" {
				return
			}
		}
		*tdTags = append(*tdTags, root)
		return
	}
	for n := root.FirstChild; n != nil; n = n.NextSibling {
		parseTrTags(n, tdTags)
	}
}

func parseTdTags(tdTags []*html.Node, rawData *[]string) {
	for _, td := range tdTags {
		if td.FirstChild.Type == html.TextNode {
			*rawData = append(*rawData, strings.TrimSpace(td.FirstChild.Data))
		}
	}
}

func exactDataFromTdTag(td *html.Node, rawData *[]string) {
	if td.Type == html.TextNode && strings.TrimSpace(td.Data) != "" {
		space := regexp.MustCompile(`\s+`)
		number := regexp.MustCompile(`\(\d\)`)
		s := space.ReplaceAllString(td.Data, " ")
		if !number.MatchString(s) {
			*rawData = append(*rawData, s)
		}
	}

	for n := td.FirstChild; n != nil; n = n.NextSibling {
		exactDataFromTdTag(n, rawData)
	}
}

func getLessonByDay(texts []string) map[string]string {
	var (
		from, to string
		m        = make(map[string]string)
		re1      = regexp.MustCompile(`\s*Từ.+đến.+`)
		re2      = regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
		re3      = regexp.MustCompile(`Thứ \d tiết.+`)
		re4      = regexp.MustCompile(`Thứ \d`)
		re5      = regexp.MustCompile(`tiết .+[^/:]`)
	)

	for _, elem := range texts {
		if re1.MatchString(elem) {
			twoDates := re2.FindAllString(elem, -1)
			from, to = twoDates[0], twoDates[1]
		}

		if re3.MatchString(elem) {
			weekDay := re4.FindString(elem)
			lesson := re5.FindString(elem)
			dates := getAllDatesInDuration(from, to, weekDay)
			for _, d := range dates {
				m[d] = lesson
			}
		}
	}

	return m
}

func getAllDatesInDuration(from, to, weekDay string) []string {
	const layout = "02/01/2006"
	var (
		wd       time.Weekday
		dates    []string
		start, _ = time.Parse(layout, from)
		end, _   = time.Parse(layout, to)
	)

	switch weekDay {
	case "Thứ 2":
		wd = time.Monday
	case "Thứ 3":
		wd = time.Tuesday
	case "Thứ 4":
		wd = time.Wednesday
	case "Thứ 5":
		wd = time.Thursday
	case "Thứ 6":
		wd = time.Friday
	case "Thứ 7":
		wd = time.Saturday
	default:
		wd = time.Sunday
	}

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == wd {
			dates = append(dates, d.Format(layout))
		}
	}

	return dates
}
