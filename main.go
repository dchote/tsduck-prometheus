package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"tsduck-prometheus/models"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func updatePidValues(target, label string, pid models.Pids) {
	// PID Bitrate
	models.TsPidBitrate.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Set(float64(pid.Bitrate))
	// PID Service Count
	models.TsPidServiceCount.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Set(float64(pid.SeviceCount))
	// PID Discontinuities Count
	if pid.Packets.Discontinuities >= 1 {
		sum := 0
		for sum < pid.Packets.Discontinuities {
			sum += 1
			models.TsPidDiscontinuity.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
		}
	}
	// Duplicate PID Count
	if pid.Packets.Duplicated >= 1 {
		sum := 0
		for sum < pid.Packets.Duplicated {
			sum += 1
			models.TsPidDuplicated.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
		}
	}
}

func updateServiceValues(target, label string, service models.Services) {
	// TS bitrate per service
	models.TsServiceBitrate.WithLabelValues(target, label, strconv.Itoa(service.Id), strconv.Itoa(service.TsId), service.Name, service.Provider, service.TypeName, strconv.Itoa(service.PcrPid), strconv.Itoa(service.PmtPid)).Set(float64(service.Bitrate))
}

func updateTsValues(target, label string, ts models.Ts) {
	// Total TS Bitrate (188byte/pkt)
	models.TsBitrate.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Set(float64(ts.Bitrate))
	// PCR TS Bitrate (188byte/pkt)
	models.TsPcrBitrate.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Set(float64(ts.PcrBitrate))
	// TS Packet Invalid Sync Count
	if ts.Packets.InvalidSyncs >= 1 {
		sum := 0
		for sum < ts.Packets.InvalidSyncs {
			sum += 1
			models.TsPacketInvalidSync.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Inc()
		}
	}
	// TS Packet Suspect Ignored Count
	if ts.Packets.SuspectIgnored >= 1 {
		sum := 0
		for sum < ts.Packets.SuspectIgnored {
			sum += 1
			models.TsPacketSuspectIgnored.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Inc()
		}
	}
	// TS Packet Transport Error Count
	if ts.Packets.TransportErrors >= 1 {
		sum := 0
		for sum < ts.Packets.TransportErrors {
			sum += 1
			models.TsPacketTeiErrors.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Inc()
		}
	}
	// TS Total PID Count
	models.TsPidCount.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Set(float64(ts.Pids.Total))
	// TS PCR PID Count
	models.TsPcrPidCount.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Set(float64(ts.Pids.Pcr))
	// TS Unreferenced PID Count
	models.TsPidUnferencedCount.WithLabelValues(target, label, strconv.Itoa(ts.Id)).Set(float64(ts.Pids.Unreferenced))
	// TsServiceClearCount (future)
	// TsServiceScrambledCount (future)
	// TsServiceCount (future)
}

func updateTableValues(target, label string, tables models.Tables) {
	// PID Max Repitition MS
	models.TsPidMaxRepititionMs.WithLabelValues(target, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MaxRepititionMs))
	// PID Max Repitition Pkt
	models.TsPidMaxRepititionPkt.WithLabelValues(target, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MaxRepititionPkt))
	// PID Min Repitition MS
	models.TsPidMinRepititionMs.WithLabelValues(target, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MinRepititionMs))
	// PID In Repitition Pkt
	models.TsPidMinRepititionPkt.WithLabelValues(target, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MinRepititionPkt))
	// PID Repitition MS
	models.TsPidRepititionMs.WithLabelValues(target, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.RepetitionMs))
	// PID Repitition Pkt
	models.TsPidRepititionPkt.WithLabelValues(target, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.RepetitionPkt))
}

func launchTsp(target, label string) {
	var tsp models.Tsp

	// Launch TSP pointing at our SRT target
	cmd := exec.Command("tsp", "-I", "srt", "--caller", target, "-P", "analyze", "-i", "1", "--json-line", "-O", "drop")
	// TSDuck outputs to stderr
	cmdReader, err := cmd.StderrPipe()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Started monitoring for %v\n", label)
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("process error %v\n", err)
		return
	}

	// Create buffer to read the stderr output
	scanner := bufio.NewScanner(cmdReader)

	for scanner.Scan() {
		s := scanner.Text()
		// Hacky way of removing the beginning of the JSON line which prepends TSDucks JSON output
		t := strings.Replace(s, "* analyze: ", "", -1)
		// Unmarshal JSON into TSP model
		json.Unmarshal([]byte(t), &tsp)

		// Update PID metrics
		for _, pid := range tsp.Pids {
			go updatePidValues(target, label, pid)
		}

		// Update Service metrics
		for _, service := range tsp.Services {
			go updateServiceValues(target, label, service)
		}

		// Update TS metrics
		go updateTsValues(target, label, tsp.Ts)

		// Update TS table metrics
		for _, table := range tsp.Tables {
			go updateTableValues(target, label, table)
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("process error %v\n", err)
	}

}

func main() {
	// Parse Args
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Not enough arguments! Plase parse at least one input e.g. 10.205.203.64:3333,My_Service")
	}

	for _, item := range args {
		s := strings.Split(item, ",")
		if len(s) < 2 {
			log.Fatal("Not enough arguments! Required format is target:port,source,label, e.g. 10.205.203.64:3333,My_Service")
		}
		// Launch TSDuck (tsp subprocess)
		go func() {
			for {
				launchTsp(s[0], s[1])

				fmt.Printf("tsp exited... restarting in 15s\n")
				time.Sleep(15 * time.Second)
			}
		}()
	}

	r := prometheus.NewRegistry()
	// Register prometheus metrics
	r.MustRegister(models.TsBitrate)
	r.MustRegister(models.TsPcrBitrate)
	r.MustRegister(models.TsPidBitrate)
	r.MustRegister(models.TsPidServiceCount)
	r.MustRegister(models.TsPidDiscontinuity)
	r.MustRegister(models.TsPidDuplicated)
	r.MustRegister(models.TsServiceBitrate)
	r.MustRegister(models.TsPidMaxRepititionMs)
	r.MustRegister(models.TsPidMaxRepititionPkt)
	r.MustRegister(models.TsPidMinRepititionMs)
	r.MustRegister(models.TsPidMinRepititionPkt)
	r.MustRegister(models.TsPidRepititionMs)
	r.MustRegister(models.TsPidRepititionPkt)
	r.MustRegister(models.TsPacketInvalidSync)
	r.MustRegister(models.TsPacketSuspectIgnored)
	r.MustRegister(models.TsPacketTeiErrors)
	r.MustRegister(models.TsPidCount)
	r.MustRegister(models.TsPcrPidCount)
	r.MustRegister(models.TsPidUnferencedCount)

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)
	http.ListenAndServe(":8000", nil)
}
