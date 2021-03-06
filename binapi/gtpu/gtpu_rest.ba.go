// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package gtpu

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RESTHandler(rpc RPCService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/gtpu_add_del_tunnel", func(w http.ResponseWriter, req *http.Request) {
		var request = new(GtpuAddDelTunnel)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.GtpuAddDelTunnel(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/gtpu_offload_rx", func(w http.ResponseWriter, req *http.Request) {
		var request = new(GtpuOffloadRx)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.GtpuOffloadRx(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/sw_interface_set_gtpu_bypass", func(w http.ResponseWriter, req *http.Request) {
		var request = new(SwInterfaceSetGtpuBypass)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.SwInterfaceSetGtpuBypass(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	return http.HandlerFunc(mux.ServeHTTP)
}
