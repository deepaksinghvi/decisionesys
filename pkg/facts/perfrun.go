package facts

import "fmt"

type PerfRunFact struct {
	TPS int
	TPSResult string // TPS_NOT_ACHIEVED or TPS_ACHIEVED
	CPURunTime_AVG float32
	RunDuration	float32
	RunStatus string
	FactMessage string
	FactResult string // VALID or INVALID
}

type ResultFact struct {

	//PerfRunFact
	PerfRunFactCompleted string `default:"COMPLETED"`  // COMPLETED
	PerfRunFactFailed string `default:"FAILED"`// FAILED
	PerfRunFactAborted string `default:"ABORTED"`// ABORTED
	PerfRunFactInProgress string `default:"INPROGRESS"`// INPROGRESS

	//FactResult
	InvalidFactResult string `default:"INVALID"`// INVALID
	ValidFactResult string `default:"VALID"`// VALID
}

func (f *PerfRunFact) GetFactResult() string {
	return fmt.Sprintf("GetFactResult")
}
