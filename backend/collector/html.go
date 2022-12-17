package collector

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetRequiredHtmlFromAccount(cookies []*http.Cookie) (string, error) {
	var (
		url       = "http://qldt.actvn.edu.vn/CMCSoft.IU.Web.Info/Reports/Form/StudentTimeTable.aspx"
		method    = "GET"
		newCookie = "ASP.NET_SessionId" + cookies[0].Value + "; " + "SignIn=" + cookies[1].Value
		client    = &http.Client{}
		req, err  = http.NewRequest(method, url, nil)
	)

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Referer", "http://qldt.actvn.edu.vn/CMCSoft.IU.Web.Info/Home.aspx")
	req.Header.Add("Accept-Language", "vi-VN,vi;q=0.9,fr-FR;q=0.8,fr;q=0.7,en-US;q=0.6,en;q=0.5")
	req.Header.Add("Cookie", newCookie)

	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	html := filterHtml(string(body))

	return html ,nil
}

func filterHtml(rawHtml string) string {
	re := regexp.MustCompile(`<table[\s\S]*?id="gridRegistered"[\s\S]*?>[\s\S]*?</table>`)
	return re.FindString(rawHtml)
}
