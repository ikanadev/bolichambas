package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func saveData(companies *[]Company) {
	now := time.Now()
	jsonData, err := json.MarshalIndent(companies, "", "  ")
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("jobs_%s.json", now.Format("2006-01-02"))
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	companies := []Company{}
	// companies = append(companies, ParseBisaJobs())
	// companies = append(companies, ParseFarmacorpJobs())
	// companies = append(companies, ParseSofiaJobs())
	// companies = append(companies, ParseEmbolJobs())
	// companies = append(companies, ParseMercantilSantaCruzJobs())
	// companies = append(companies, ParseBNBJobs())
	// companies = append(companies, parseTigoJobs())
	// companies = append(companies, parsePilAndinaJobs())
	// companies = append(companies, parseIntiJobs())
	// companies = append(companies, parseBcpJobs())
	// companies = append(companies, parsePedidosYaJobs())
	// companies = append(companies, parseGanaderoJobs())
	// companies = append(companies, parseEconomicoJobs())
	// companies = append(companies, parseBagoJobs())
	// companies = append(companies, parseGrupoVenadoJobs())
	// companies = append(companies, ParseSolJobs())
	// companies = append(companies, parseAlianzaJobs())
	// companies = append(companies, parseFieJobs())
	// companies = append(companies, parseVivaJobs())
	// companies = append(companies, parseGrupoRodaJobs())
	// companies = append(companies, parseLaPapeleraJobs())
	// companies = append(companies, parseDhlJobs())
	// companies = append(companies, parseTotalEnergiesJobs())
	// companies = append(companies, parseFairPlayJobs())
	// companies = append(companies, parseMonopolJobs())
	// companies = append(companies, parseUnividaJobs())
	// companies = append(companies, parseBdpJobs())
	companies = append(companies, ParseDiaconiaJobs()) // 245 so fas
	saveData(&companies)
}
