package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	basic "github.com/elcamino/go101/01_basic"
	types "github.com/elcamino/go101/02_types"
	arrays "github.com/elcamino/go101/03_arrays"
	control "github.com/elcamino/go101/04_control"
	errors "github.com/elcamino/go101/05_errors"
	funcs "github.com/elcamino/go101/06_funcs"
	maps "github.com/elcamino/go101/07_maps"
	jayson "github.com/elcamino/go101/08_json"
	httpea "github.com/elcamino/go101/09_http"
	timedemo "github.com/elcamino/go101/10_time"
	concurrency "github.com/elcamino/go101/11_concurrency"
	contextdemo "github.com/elcamino/go101/12_context"
	logging "github.com/elcamino/go101/13_logging"
	configuration "github.com/elcamino/go101/14_configuration"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	loglevel := flag.String("loglevel", "info", "set the logrus loglevel")
	callVars := flag.Bool("vars", false, "call the vars demo")
	callNames := flag.Bool("names", false, "call the names demo")
	callControl := flag.Bool("control", false, "call the control demo")
	callFuncs := flag.Bool("funcs", false, "call the funcs demo")
	callArrays := flag.Bool("arrays", false, "call the arrays demo")
	callTypes := flag.Bool("types", false, "call the types demo")
	callMaps := flag.Bool("maps", false, "call the maps demo")
	callErrors := flag.Bool("errors", false, "call the errors demo")
	callInterfaces := flag.Bool("interfaces", false, "call the interfaces demo")
	callPointers := flag.Bool("pointers", false, "call the pointers demo")
	callConsts := flag.Bool("consts", false, "call the consts demo")
	callJSON := flag.Bool("json", false, "call the json demo")
	callHTTP := flag.Bool("http", false, "call the HTTP demo")
	callTime := flag.Bool("time", false, "call the Time demo")
	callConcurrency := flag.Bool("concurrency", false, "call the concurrency demo")
	callBuffered := flag.Bool("buffered", false, "call the buffered demo")
	callContext := flag.Bool("context", false, "call the context demo")
	callStdlog := flag.Bool("stdlog", false, "call the standard log demo")
	callLogrus := flag.Bool("logrus", false, "call the logrus demo")
	callViper := flag.Bool("viper", false, "call the viper demo")

	// flag.Parse()

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	switch {
	case *callVars:
		basic.Vars()
	case *callNames:
		basic.Names()
	case *callControl:
		control.Run()
	case *callFuncs:
		funcs.Run()
	case *callArrays:
		arrays.Run()
	case *callTypes:
		types.Run()
	case *callMaps:
		maps.Maps()
	case *callErrors:
		errors.Errors()
	case *callInterfaces:
		types.Interfaces()
	case *callPointers:
		basic.Pointers()
	case *callConsts:
		basic.Consts()
	case *callTypes:
		basic.Types()
	case *callJSON:
		jayson.Jayson()
	case *callHTTP:
		httpea.Http()
	case *callTime:
		timedemo.Time()
	case *callConcurrency:
		concurrency.Concurrency()
	case *callBuffered:
		concurrency.Buffered()
	case *callContext:
		contextdemo.Context()
	case *callStdlog:
		logging.Logging()
	case *callLogrus:
		logging.Logrus(*loglevel)
	case *callViper:
		configuration.Viper()
	default:
		flag.Usage()
	}

}

func init() {
	seed := time.Now().UnixNano()
	fmt.Printf("[init] random seed: %d\n", seed)
	rand.Seed(seed)
}
