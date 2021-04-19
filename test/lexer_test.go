package test

import (
	"github.com/blastbao/gengine/builder"
	"github.com/blastbao/gengine/context"
	"github.com/blastbao/gengine/engine"
	"testing"
)

func Test_lexer(t *testing.T) {

	lexer_rule := `
rule "test" salience 1
begin
规则管理
end
`
	dataContext := context.NewDataContext()
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(lexer_rule)
	if e1 != nil {
		panic(e1)
	}

	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	if e != nil {
		panic(e)
	}

}
