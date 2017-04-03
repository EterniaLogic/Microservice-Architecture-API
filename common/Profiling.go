package common
import "github.com/gorilla/mux"
import "net/http/pprof" // /debug/pprof/heap

func ProfilingRouterHandler(router *mux.Router){
	dbgprefix := "/debug/pprof";
	router.HandleFunc(dbgprefix+"/", pprof.Index);
	router.HandleFunc(dbgprefix+"/block", pprof.Index);
	router.HandleFunc(dbgprefix+"/heap", pprof.Index);
	router.HandleFunc(dbgprefix+"/threadcreate", pprof.Index);
	router.HandleFunc(dbgprefix+"/goroutine", pprof.Index);
	router.HandleFunc(dbgprefix+"/cmdline", pprof.Cmdline);
	router.HandleFunc(dbgprefix+"/profile", pprof.Profile);
	router.HandleFunc(dbgprefix+"/symbol", pprof.Symbol);
}