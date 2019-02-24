package league

import (
	"net/http"
	"net/url"
)

func (l League) GetLeague() {
	req := http.Request{
		URL: url.URL{Host: "http://games.espn.com/ffl/api/v2/"},
	}
	espnS2 := http.Cookie{
		Name: "espn_s2",
		Value: "AEBUBSDyP6%2FXzKzc88QaBaBq7iJUzklumOk5JLbx0DKGcLvTHkUJGTXmmqOOEoxPNjvvEszYuAHflwr7lPURWhtNhLgvWID%2BHhGc7u%2FdWAQM0yRsTR0resXumgJS1tMSEn4UdFy%2FE32XGAWcg4A1PZHYGkK7FFKxEP1HnBxUgYHNFWbPo7cJGPIBE7oDadz0CLlLHrRf9E9oNKzmL5FLny57unKhynNnr9gVR%2FGmZEGCYFbX%2BRTmJPIHIf%2FhnwQbZzY4mNG3%2F7jk9h3cq6ZW7kuL",
		Method: "Get",
	}
	swid = http.Cookie{
		Name: "SWID",
		Value: "{7A703CA3-FB06-4A12-A868-46266E40A33D}",
	}

	req.AddCookie(espnS2)
	req.AddCookie(swid)

	client := http.Client{}

	resp, err := client.Do(req)
}
