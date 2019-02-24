package main

import (
	"github.com/espn-go-api/pkg/league"
	"github.com/espn-go-api/pkg/analytics"
)

func main() {
	hc := league.League{
		leagueID := 361793,
		seasonID := 2018,
		seasonLength := 16,
		espnS2 := 'AEBUBSDyP6%2FXzKzc88QaBaBq7iJUzklumOk5JLbx0DKGcLvTHkUJGTXmmqOOEoxPNjvvEszYuAHflwr7lPURWhtNhLgvWID%2BHhGc7u%2FdWAQM0yRsTR0resXumgJS1tMSEn4UdFy%2FE32XGAWcg4A1PZHYGkK7FFKxEP1HnBxUgYHNFWbPo7cJGPIBE7oDadz0CLlLHrRf9E9oNKzmL5FLny57unKhynNnr9gVR%2FGmZEGCYFbX%2BRTmJPIHIf%2FhnwQbZzY4mNG3%2F7jk9h3cq6ZW7kuL',
		swid := '{7A703CA3-FB06-4A12-A868-46266E40A33D}'
	}

	analytics.Start(hc)
}
