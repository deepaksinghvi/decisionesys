package util

import (
	"encoding/json"
	"fmt"
	"github.com/deepaksinghvi/decisionesys/pkg/facts"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/mcuadros/go-defaults"
	log "github.com/sirupsen/logrus"
	"os"
)

func PrintRun(perfrunFact facts.PerfRunFact) {
	perfrunJSON, err := json.MarshalIndent(perfrunFact, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("-----")
	log.Println("Processed Perf Run Fact \n", string(perfrunJSON))
	fmt.Println("-----")
}


func GetValidPerfRun(dataCtx *ast.IDataContext) *facts.PerfRunFact {

	perfRunFact := &facts.PerfRunFact{
		TPS:              199,
		CPURunTime_AVG:   1.0,
		RunDuration:      200.4,
		RunStatus:        "COMPLETED",
	}
	return addFacts(*dataCtx, perfRunFact)
}

func GetInvalidPerfRun(dataCtx *ast.IDataContext) *facts.PerfRunFact {
	perfRunFact := &facts.PerfRunFact{
		TPS:              199,
		CPURunTime_AVG:   2.0,
		RunDuration:      200.4,
		RunStatus:        "ABORTED",
	}
	return addFacts(*dataCtx, perfRunFact)
}

func GetAchievedTPS(dataCtx *ast.IDataContext) *facts.PerfRunFact {
	perfRunFact := &facts.PerfRunFact{
		TPS:              199,
		CPURunTime_AVG:   3.0,
		RunDuration:      200.4,
		RunStatus:        "COMPLETED",
	}
	return addFacts(*dataCtx, perfRunFact)
}

func GetUnAchivedTPS(dataCtx *ast.IDataContext) *facts.PerfRunFact {
	perfRunFact := &facts.PerfRunFact{
		TPS:              99,
		CPURunTime_AVG:   4.0,
		RunDuration:      200.4,
		RunStatus:        "COMPLETED",
	}
	return addFacts(*dataCtx, perfRunFact)
}




func addFacts(dataCtx ast.IDataContext, perfRunFact *facts.PerfRunFact) *facts.PerfRunFact {
	err := dataCtx.Add("PerfRunFact", perfRunFact)
	if err != nil {
		panic(err)
	}

	// ResultFact to be used as the status literls
	resultFact := &facts.ResultFact{}
	defaults.SetDefaults(resultFact)
	err = dataCtx.Add("ResultFact", resultFact)
	if err != nil {
		panic(err)
	}

	return perfRunFact
}



func GetRuleJson(rulefilename string) []byte {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "/pkg/rules/"+rulefilename)
	if err != nil {
		panic(err)
	}
	underlying := pkg.NewReaderResource(f)
	resource := pkg.NewJSONResourceFromResource(underlying)
	loaded, err := resource.Load()
	if err != nil {
		log.Error(err)
	}
	// print rule
	log.Println(string(loaded))
	return loaded
}
