## Rule Based Decision System

The two major components of rule-based artificial intelligence models are 
* Set of rules 
* Set of facts

A basic artificial intelligence model can be with the help of these two components.

How rule based AI is different from ML?
Machine learning systems are probabilistic and rule-based AI models are deterministic.


### [Grule Engine](https://github.com/hyperjumptech/grule-rule-engine)
Inspired by the acclaimed JBOSS Drools build for golang. It has its own dsl.
For this example I have used json based rules

###

#### Facts
PerfRunFact is used as the fact
```
type PerfRunFact struct {
    TPS int
    CPURunTime_AVG float32
    RunDuration	float32
    RunStatus string
    
    // Following are used as assignment after rules processing
    FactErrorMessage string
    FactResult string // VALID or INVALID
    TPSResult string // TPS_VALID or TPS_INVALID
}
```

ResultFact is used as the literals
```
type ResultFact struct {
    //PerfRunFact
    PerfRunFactCompleted string `default:"COMPLETED"`
    PerfRunFactFailed string `default:"FAILED"`
    PerfRunFactAborted string `default:"ABORTED"`
    PerfRunFactInProgress string `default:"INPROGRESS"`
    
    //FactResult
    InvalidFactResult string `default:"INVALID"`
    ValidFactResult string `default:"VALID"`
}
```


Rules are added in rules/perfruncompleted.json