package filter_csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type WHSite struct {
	UniqueNumber       string
	IDNo               string
	Rev                string
	NameEN             string
	NameFR             string
	NameES             string
	NameRU             string
	NameAR             string
	NameZH             string
	ShortDescriptionEN string
	ShortDescriptionFR string
	ShortDescriptionES string
	ShortDescriptionRU string
	ShortDescriptionAR string
	ShortDescriptionZH string
	JustificationEN    string
	JustificationFR    string
	DateInscribed      string
	SecondaryDates     string
	Danger             string
	DateEnd            string
	DangerList         string
	Longitude          string
	Latitude           string
	AreaHectares       string
	C1                 string
	C2                 string
	C3                 string
	C4                 string
	C5                 string
	C6                 string
	N7                 string
	N8                 string
	N9                 string
	N10                string
	CriteriaTxt        string
	Category           string
	CategoryShort      string
	StatesNameEN       string
	StatesNameFR       string
	StatesNameES       string
	StatesNameRU       string
	StatesNameAR       string
	StatesNameZH       string
	RegionEN           string
	RegionFR           string
	ISOCode            string
	UDNPCode           string
	Transboundary      bool
}

// Filter sites by category or state
func FilterByCategoryORState(sites []WHSite, category, stateNameEN string) []WHSite {
	if len(category) == 0 && len(stateNameEN) == 0 {
		return sites
	}

	var filteredSites []WHSite
	for _, site := range sites {

		if strings.EqualFold(site.Category, category) || strings.EqualFold(site.StatesNameEN, stateNameEN) {
			whSite := WHSite{
				NameEN:             site.NameEN,
				ShortDescriptionEN: site.ShortDescriptionEN,
				JustificationEN:    site.JustificationEN,
				Category:           site.Category,
			}
			filteredSites = append(filteredSites, whSite)
		}
	}
	return filteredSites
}

func GetAllSites(linhas [][]string) []WHSite {
	var sites []WHSite
	for _, linha := range linhas {
		site := WHSite{
			UniqueNumber:       linha[0],
			IDNo:               linha[1],
			Rev:                linha[2],
			NameEN:             linha[3],
			NameFR:             linha[4],
			NameES:             linha[5],
			NameRU:             linha[6],
			NameAR:             linha[7],
			NameZH:             linha[8],
			ShortDescriptionEN: linha[9],
			ShortDescriptionFR: linha[10],
			ShortDescriptionES: linha[11],
			ShortDescriptionRU: linha[12],
			ShortDescriptionAR: linha[13],
			ShortDescriptionZH: linha[14],
			JustificationEN:    linha[15],
			JustificationFR:    linha[16],
			DateInscribed:      linha[17],
			SecondaryDates:     linha[18],
			Danger:             linha[19],
			DateEnd:            linha[20],
			DangerList:         linha[21],
			Longitude:          linha[22],
			Latitude:           linha[23],
			AreaHectares:       linha[24],
			C1:                 linha[25],
			C2:                 linha[26],
			C3:                 linha[27],
			C4:                 linha[28],
			C5:                 linha[29],
			C6:                 linha[30],
			N7:                 linha[31],
			N8:                 linha[32],
			N9:                 linha[33],
			N10:                linha[34],
			CriteriaTxt:        linha[35],
			Category:           linha[36],
			CategoryShort:      linha[37],
			StatesNameEN:       linha[38],
			StatesNameFR:       linha[39],
			StatesNameES:       linha[40],
			StatesNameRU:       linha[41],
			StatesNameAR:       linha[42],
			StatesNameZH:       linha[43],
			RegionEN:           linha[44],
			RegionFR:           linha[45],
			ISOCode:            linha[46],
			UDNPCode:           linha[47],
		}
		sites = append(sites, site)
	}
	return sites
}

// Read the csv file and return all lines
func ReadCSV() ([]WHSite, error) {
	file, err := os.Open("whc-sites.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	linhas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all lines:", err)
		return nil, err
	}

	sites := GetAllSites(linhas)
	return sites, err
}
