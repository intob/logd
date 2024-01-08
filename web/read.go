package web

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fxamacker/cbor/v2"
	"github.com/swissinfo-ch/logd/msg"
)

func (s *Webserver) handleRead(w http.ResponseWriter, r *http.Request) {
	if !s.isAuthedForReading(r) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	offset := 0
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, "limit must be an integer", http.StatusBadRequest)
			return
		}
	}
	limit := 1000
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "limit must be an integer", http.StatusBadRequest)
			return
		}
		if limit > 1000 {
			limit = 1000
		}
	}
	envQ := r.URL.Query().Get("env")
	svcQ := r.URL.Query().Get("svc")
	fnQ := r.URL.Query().Get("fn")
	results := make([]*msg.Msg, 0)
	pages := 0
	bufSize := int(s.Buf.Size())
	offset32 := uint32(offset)
	limit32 := uint32(limit)
	for len(results) < int(limit) && pages*limit < bufSize/10 {
		items := s.Buf.Read(offset32, limit32)
		pages++
		offset += limit
		m := &msg.Msg{}
		for _, i := range items {
			buf := bytes.NewBuffer(*i)
			gzr, err := gzip.NewReader(buf)
			if err != nil {
				fmt.Println("handleRead: gzip reader:", err)
				continue
			}
			dec := cbor.NewDecoder(gzr)
			err = dec.Decode(m)
			if err != nil {
				fmt.Println("hadleRead: decode cbor:", err)
				continue
			}
			if envQ != "" && envQ != m.Env {
				continue
			}
			if svcQ != "" && svcQ != m.Svc {
				continue
			}
			if fnQ != "" && fnQ != m.Fn {
				continue
			}
			results = append(results, m)
		}
	}
	data, err := json.Marshal(results)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(data)
}
