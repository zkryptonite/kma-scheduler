package collector

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

const uri = "http://qldt.actvn.edu.vn/CMCSoft.IU.Web.Info/Login.aspx" 

func GetCookies(username, password string) ([]*http.Cookie, error) {
	method := "POST"
	encyptionPwd := GetMD5Hash(password)
	body := url.Values{}

	body.Add("__EVENTTARGET", "")
	body.Add("__EVENTARGUMENT", "")
	body.Add("__LASTFOCUS", "")
	body.Add("__VIEWSTATE", "kR8VUoxVMbcvgjKMvXJH06KQeJLsU/oUow7ZUa43izQjh4bP2szISIYk6dr2Ysu+Er0TZaBm6w6ioMet86LvpNOE4R8zWvxKN6k5DxXL3Z3AIFiNDKL+dpr291hbVbIyT0t9rWcqU/I2vwg+4rjEHndYpZj5c2nzBSKMzwJZqCYktX1m/kUAwUkzoiqtJNmTNLMpv1OcB9fLLEeVyW+wKpFS9EWqNZgVpBG+/jOpbz9yOzJCTMH21tGNa8JRSc07GrVniFCCELmFgVE4A70C8PzybzFZLAkv2V91+vKBYi3zScGlry5d9d5mMcFjpakiCrmfLay6mzC2jLwJUoT8hRx0zTbVaBHkMJUViva0NRfza6zjazEADV8wi8RlGlB7n9njkpLd3NPe7lq5w3cWk5Sn2cum1xm/bu9enhdoqHw36t0Hpj9VWDQOsmgj3MqhThlAmTSrWErDey8OQ5lzFbKRG8uP6HGsTqexQdPiGMOPpQzv/2AQRlmxEMEDNpef2eQEv/6kNbCSyqmEUkvW5ErMxdF5TOdYuVbwfm6eLUmpS69ivLIqouf1beItEKsuboLV9OAFhKNnnV1plpPXICto4Q6YakXEaLmwJ/MdxfvR9FzmiZfScYQd8BlRAbG/rgtAKzOvxeS3M5iehlJb81/toT2dxZKx3eElpjl2UFjDkrODPQoWxyidUA18XP80z6ayTFnkCVaLVTFBLZwY5DwpS4AgdChJHUKbLfbdj4UequUD9fLT/T9+GoWydqXgvkDJjQFLFJyR41H5ZUcc/7B9iYhXVDbFd4czMwu9OjKWYqaqVw1sUw6+AcFhtAa6awTMxdD9ZAmaX+H6M6hrM78r7YDzxh+PPl9O29lqa5rps/TrQg19XWgNYejSiJKkAtMRYcMrpAy/KTRCzyjgI4FhB0/N2FSnxSpwI5yQFc2WfMXsrksLYoWw5gWqSWqdAkbYJCJylwkk7vutHAflm7n+IDXSXSh3f41euA+EtCqfQO3Y+ZsXnkCbrfahMF0GilCxIporKuOO+X/OW8EqO0yzuWJtDBvR7My+9uGjISFYM+mO1HLB8LtpDLG1ewUxcyawcEdhgackcFW/2AZfG9/HPBp+3oyBAXnlrS9UD3TIpjIHbeWsdWGd7WTQarDZlMAyKJ9h+9ALwDgBnlDv6lN/y6/xnvLo0oYbC1FqSNzDdWq6JRtcZJfNf6euKU0eJdbq0QA064OLsQst3igmMueFahgniaoOXljZklrKxIs2MeWTSwNAxHfCesDE3S2+xQxvlGov65DXGZhftxCy2kyXF0N5+c9aoqhIPi1PP8C3GXPybFkVysD98+LPgXaFRPTDuPYtrVdgSouG7CB3iPZG/0eJqRujkbCEjut5XRUAB5Ms3tjzVeSfEeSv55cJp4i3g4cuIOg6rTgcdZ40JjYWM5lGHMPVevnDTygzYLCmV6POs5MGEogNhffhA9LiTX8Tj+tcOMWb7gMVnMuHbSpIii1rSU3vb06pqBExWNhoVuyVzDWmtMHYguo2AwoLnQwDnACXgqHccjMSd1KqgsqeIAyzlES3JFHpdJf5ZkKDMsIFKmc7iIDVrxJHNiP36nGnXmAVP9wv3nQPdCrHAgdttuou+hb1Wz2BtdIX8/BG5wCOukVHon1QEjEC2dzm4hhb49iftDfyNR48AzGouTYxWPgVRPKU40xvQYeTck9Sddd/KZyfQBX8wgQw+7cbYqQ3+We9LN8rd4pqv9uTJbzuGL4XJ+MQ7jlEuVZH1YtGmU7/si8s5nhPwEN8IDYQpSaJRe7uzk1mP+ys68t50g==")
	body.Add("__VIEWSTATEGENERATOR", "D620498B")
	body.Add("PageHeader1$drpNgonNgu", "E43296C6F24C4410A894F46D57D2D3AB")
	body.Add("PageHeader1$hidisNotify", "0")
	body.Add("PageHeader1$hidValueNotify", "")
	body.Add("txtUserName", username)
	body.Add("txtPassword", encyptionPwd)
	body.Add("btnSubmit", "Đăng+nhập")
	body.Add("hidUserId", "")
	body.Add("hidUserFullName", "")
	body.Add("hidTrainingSystemId", "")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(method, uri, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Origin", "http://qldt.actvn.edu.vn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Referer", "http://qldt.actvn.edu.vn/CMCSoft.IU.Web.Info/Login.aspx")
	req.Header.Add("Accept-Language", "vi-VN,vi;q=0.9,fr-FR;q=0.8,fr;q=0.7,en-US;q=0.6,en;q=0.5")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res.Cookies(), nil
}
