package warehouse

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var tc *TestState
var testT *testing.T

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty",
	Paths:  []string{"../features"},
}

func TestMain(m *testing.M) {
	// Store the testing.T instance
	t := &testing.T{}
	testT = t

	status := godog.TestSuite{
		Name:                 "warehouse",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Pass testT to NewTestContext
		tc = NewTestContext(testT)
		return ctx, nil
	})

	initCalculateCapacitySteps(ctx)
	initFindAvailableWarehouseSteps(ctx)
	initGetFullyUtilizedDatesSteps(ctx)
	initGetLeastUsedWarehouseSteps(ctx)
}
