package main

import (
	"strings"
    "fmt"
    "os"
	"bufio"
	"io/ioutil"
    "net/http"
	"encoding/json"
	"text/tabwriter"
	"time"
)

type FundRankingResp struct {
	Status bool			`json:"status"`
	Error string		`json:"error"`
	Datas []Data		`json:"data"`
}

type Data struct {
	MstarId string			`json:"mstar_id"`
	ThaiFundCode string		`json:"thailand_fund_code"`
	NavReturn float64		`json:"nav_return"`
	Nav float64				`json:"nav"`
	NavDate string			`json:"nav_date"`
	AvgReturn float64		`json:"avg_return"`
}

func main() {

	fmt.Print("Please input time range to view a list of funds(1D, 1W, 1M, 1Y): ")
	reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
	inputParm := strings.TrimSuffix(input, "\n")
    
	period := ValidateInput(inputParm)
	if period == "" {
		fmt.Println("invalid input")
		os.Exit(1)
	}

	fundRankingSlice := GetFundRanking(period)

	PrintFundRanking(fundRankingSlice)

}

func ValidateInput(inputParm string) (string){
	if strings.EqualFold("1D", inputParm) {
		return "1D"
	} else if strings.EqualFold("1W", inputParm) {
		return "1W"
	} else if strings.EqualFold("1M", inputParm) {
		return "1M"
	} else if strings.EqualFold("1Y", inputParm) {
		return "1Y"
	} else {
		return "";
	}
}

func GetFundRanking(period string) ([]Data) {

	response, err := http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-" + period + ".json")
    if err != nil {
        fmt.Print(err.Error())
	    os.Exit(1)
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Print(err.Error())
	}

	fundRankingResp := FundRankingResp{}
	json.Unmarshal([]byte(responseData), &fundRankingResp)

	return fundRankingResp.Datas

}

func PrintFundRanking(fundRankingSlice []Data) {

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	fmt.Fprintln(w, "Name\tRank of fund\tUpdated date\tPerfomance\tPrice\t")
	layout := "2006-01-02T15:04:05.000Z"
	for i, fundRank := range fundRankingSlice {

		navDate, err := time.Parse(layout, fundRank.NavDate)
		if err != nil {
			fmt.Println(err)
		}

		concatenated := fmt.Sprintf("%s\t%d\t%s\t%.5f\t%.4f", fundRank.ThaiFundCode, i+1, navDate.Format("2006/01/02 15:04:05"), fundRank.NavReturn, fundRank.Nav)
		fmt.Fprintln(w, concatenated)
	}
	w.Flush()

}