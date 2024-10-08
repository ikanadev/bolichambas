package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

func saveData(companies *[]Company) {
	now := time.Now()
	// jsonData, err := json.MarshalIndent(companies, "", "  ")
	jsonData, err := json.Marshal(companies)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("jobs_%s.json", now.Format("2006-01-02"))
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func addCompanyData(companies *[]Company, mu *sync.Mutex, fn func() Company) {
	company := fn()
	mu.Lock()
	defer mu.Unlock()
	*companies = append(*companies, company)
}

var parsers = [](func() Company){
	parseBisaJobs,
	parseFarmacorpJobs,
	ParseSofiaJobs,
	ParseEmbolJobs,
	ParseMercantilSantaCruzJobs,
	ParseBNBJobs,
	parseTigoJobs,
	parsePilAndinaJobs,
	parseIntiJobs,
	parseBcpJobs,
	parsePedidosYaJobs,
	parseGanaderoJobs,
	parseEconomicoJobs,
	parseBagoJobs,
	parseGrupoVenadoJobs,
	ParseSolJobs,
	parseAlianzaJobs,
	parseFieJobs,
	parseVivaJobs,
	parseGrupoRodaJobs,
	parseLaPapeleraJobs,
	parseDhlJobs,
	parseTotalEnergiesJobs,
	parseFairPlayJobs,
	parseMonopolJobs,
	parseUnividaJobs,
	parseBdpJobs,
	ParseDiaconiaJobs,
	parseDatecJobs,
	parseCamsaJobs,
	parseBBOJobs,
	parseClinicaAmericasJobs,
	parseUPDSJobs,
	parseProesaJobs,
	parseUnionAgroJobs,
	parseSoboceJobs,
	parseOpalJobs,
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	start := time.Now()
	companies := []Company{}

	wg.Add(len(parsers))
	for _, fn := range parsers {
		go func() {
			addCompanyData(&companies, &mu, fn)
			wg.Done()
		}()
	}
	wg.Wait()

	total := 0
	for i := range companies {
		total += len(companies[i].Jobs)
	}
	fmt.Println("#####################")
	fmt.Printf("A total of %d jobs parsed in %v\n", total, time.Since(start))
	saveData(&companies)
}
