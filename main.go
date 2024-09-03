package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"samplehttpserver/types"
	"samplehttpserver/utils"
	"syscall"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		startServer()
	}()

	<-done
}

var log = logrus.New()
var tenantList = make([]types.Param, 0)

func startServer() {
	tenantList = append(tenantList, types.Param{SubCode: "subcode1"})
	tenantList = append(tenantList, types.Param{SubCode: "subcode2"})
	tenantList = append(tenantList, types.Param{SubCode: "subcode3"})
	tenantList = append(tenantList, types.Param{SubCode: "subcode4"})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/getparams.execute", pluginHandler)
	router.HandleFunc("/api/v1/tenants", tenantListHandler)
	log.Info("Starting Argo plugin server on port 8000")
	http.ListenAndServe(":8000", router)
}

func pluginHandler(w http.ResponseWriter, r *http.Request) {
	argoResponseBody := types.Response{types.Output{Parameters: tenantList}}
	resp, err := utils.ConvertToJson(argoResponseBody)
	if err != nil {
		log.Errorf("Error converting argo plugin handler response to json: %s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Error creating response. Check logs of operator for more details."))
	} else {
		log.Debugf("Successfully built plugin gen reponse and sent.")
		w.WriteHeader(200)
		w.Write(resp)
	}
}

func tenantListHandler(w http.ResponseWriter, req *http.Request) {
	resp, err := utils.ConvertToJson(tenantList)
	if err != nil {
		log.Errorf("Error converting tenant list to json: %s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Error Getting the tenant list. Check logs of operator for more details."))
	} else {
		log.Debugf("Successfully retrieved tenant list and reponse sent.")
		w.WriteHeader(200)
		w.Write(resp)
	}
}
