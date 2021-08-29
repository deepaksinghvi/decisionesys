package main

import (
	"github.com/deepaksinghvi/decisionesys/pkg/facts"
	"github.com/deepaksinghvi/decisionesys/pkg/util"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	log "github.com/sirupsen/logrus"
)

func main() {
	// creating knowledgebase with rules
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	loaded := util.GetRuleJson("perfruncompleted.json")
	bs := pkg.NewBytesResource(loaded)

	err := ruleBuilder.BuildRuleFromResource("PerfRunRules", "0.0.1", bs)
	if err != nil {
		panic(err)
	}
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("PerfRunRules", "0.0.1")

	/**
	// by default MaxCycles are 5000 and if we do not Retract in rule,
	   rule engine will try it for the MaxCycle times.
	*/
	g := &engine.GruleEngine{MaxCycle: 40}
	dataCtx := ast.NewDataContext()
	// Rule execution with different facts
	ExecuteRules(g, knowledgeBase, util.GetValidPerfRun(&dataCtx), dataCtx)
	g = &engine.GruleEngine{MaxCycle: 40}
	dataCtx = ast.NewDataContext()
	ExecuteRules(g, knowledgeBase, util.GetInvalidPerfRun(&dataCtx), dataCtx)
	g = &engine.GruleEngine{MaxCycle: 40}
	dataCtx = ast.NewDataContext()
	ExecuteRules(g, knowledgeBase, util.GetAchievedTPS(&dataCtx), dataCtx)
	g = &engine.GruleEngine{MaxCycle: 40}
	dataCtx = ast.NewDataContext()
	ExecuteRules(g, knowledgeBase, util.GetUnAchivedTPS(&dataCtx), dataCtx)
}

func ExecuteRules(g *engine.GruleEngine, knowledgeBase *ast.KnowledgeBase, perfRunFact *facts.PerfRunFact, dataCtx ast.IDataContext) {
	// added datacontext
	err := g.Execute(dataCtx, knowledgeBase)
	if err != nil {
		log.Error(err)
	}
	util.PrintRun(*perfRunFact)
}



