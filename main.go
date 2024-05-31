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

	"github.com/docopt/docopt-go"
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

func updateSRTGlobalValues(target, label string, global models.Global) {
	models.SRTRTTMs.WithLabelValues(target, label).Set(float64(global.Instant.RTTMs))
}

func updateSRTRecieveValues(target, label string, recieve models.Receive) {
	models.SRTIntervalReceiveBytes.WithLabelValues(target, label).Set(float64(recieve.Interval.Bytes))
	models.SRTIntervalReceiveDropBytes.WithLabelValues(target, label).Set(float64(recieve.Interval.DropBytes))
	models.SRTIntervalReceiveDroppedPackets.WithLabelValues(target, label).Set(float64(recieve.Interval.DroppedPackets))
	models.SRTIntervalReceiveIgnoredLatePackets.WithLabelValues(target, label).Set(float64(recieve.Interval.IgnoredLatePackets))
	models.SRTIntervalReceiveLossBytes.WithLabelValues(target, label).Set(float64(recieve.Interval.LossBytes))
	models.SRTIntervalReceiveLostPackets.WithLabelValues(target, label).Set(float64(recieve.Interval.LostPackets))
	models.SRTIntervalReceivePackets.WithLabelValues(target, label).Set(float64(recieve.Interval.Packets))
	models.SRTIntervalRateMbps.WithLabelValues(target, label).Set(float64(recieve.Interval.RateMbps))
	models.SRTIntervalReceiveReorderDistancePackets.WithLabelValues(target, label).Set(float64(recieve.Interval.ReorderDistancePackets))
	models.SRTIntervalReceiveRetransmittedPackets.WithLabelValues(target, label).Set(float64(recieve.Interval.RetransmittedPackets))
	models.SRTIntervalReceiveSentAckPackets.WithLabelValues(target, label).Set(float64(recieve.Interval.SentAckPackets))
	models.SRTIntervalReceiveSentNakPackets.WithLabelValues(target, label).Set(float64(recieve.Interval.SentNakPackets))
}

func launchTsp(target, label string) {
	// Launch TSP pointing at our SRT target
	cmd := exec.Command("tsp", "-I", "srt", "--caller", target, "--statistics-interval", "1000", "--json-line", "-P", "analyze", "-i", "1", "--json-line", "-O", "drop")
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

	var tspAnalyze models.TspAnalyze
	var tspSRT models.TspSRT

	var analyzeMatch = "* analyze: "
	var srtMatch = "* srt: "

	for scanner.Scan() {
		s := scanner.Text()

		if s[:len(analyzeMatch)] == analyzeMatch {
			// capture output from analyze function

			t := strings.Replace(s, analyzeMatch, "", -1)
			// Unmarshal JSON into TSP Analyze model
			json.Unmarshal([]byte(t), &tspAnalyze)

			//fmt.Printf("tsp analyze %v\n", tspAnalyze)

			// Update PID metrics
			for _, pid := range tspAnalyze.Pids {
				go updatePidValues(target, label, pid)
			}

			// Update Service metrics
			for _, service := range tspAnalyze.Services {
				go updateServiceValues(target, label, service)
			}

			// Update TS metrics
			go updateTsValues(target, label, tspAnalyze.Ts)

			// Update TS table metrics
			for _, table := range tspAnalyze.Tables {
				go updateTableValues(target, label, table)
			}

		} else if s[:len(srtMatch)] == srtMatch {
			// capture srt statistics output
			t := strings.Replace(s, srtMatch, "", -1)
			// Unmarshal JSON into TSP SRT model
			json.Unmarshal([]byte(t), &tspSRT)

			go updateSRTGlobalValues(target, label, tspSRT.Global)
			go updateSRTRecieveValues(target, label, tspSRT.Receive)
			//fmt.Printf("tsp srt %v\n", tspSRT)
		}

	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("process error %v\n", err)
	}

}

var (
	target     string
	label      string
	listenPort int
)

func cliArguments() {
	usage := `
Usage: tsduck-prometheus [options]

Options:
  -t, --target=<host:port>      Specify target host:port for SRT listener
  -l, --label=<label>           Specify label to use in Prometheus
  -p, --port=<port>             Specify listen port [default: 8000]
  -h, --help                    Show this screen.
`
	args, _ := docopt.ParseArgs(usage, os.Args[1:], "")

	target, _ = args.String("--target")
	label, _ = args.String("--label")
	listenPort, _ = args.Int("--port")

	if target == "" || label == "" {
		docopt.PrintHelpAndExit(nil, usage)
	}
}

func main() {
	cliArguments()

	go func() {
		for {
			models.ConReconnectAttempts.WithLabelValues(target, label).Inc()

			launchTsp(target, label)

			fmt.Printf("tsp exited... restarting in 15s\n")
			time.Sleep(15 * time.Second)
		}
	}()

	r := prometheus.NewRegistry()
	// Register prometheus metrics

	// transport stream
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

	// connection
	r.MustRegister(models.ConReconnectAttempts)

	// srt
	r.MustRegister(models.SRTRTTMs)
	r.MustRegister(models.SRTIntervalReceiveBytes)
	r.MustRegister(models.SRTIntervalReceiveDropBytes)
	r.MustRegister(models.SRTIntervalReceiveDroppedPackets)
	r.MustRegister(models.SRTIntervalReceiveIgnoredLatePackets)
	r.MustRegister(models.SRTIntervalReceiveLossBytes)
	r.MustRegister(models.SRTIntervalReceiveLostPackets)
	r.MustRegister(models.SRTIntervalReceivePackets)
	r.MustRegister(models.SRTIntervalRateMbps)
	r.MustRegister(models.SRTIntervalReceiveReorderDistancePackets)
	r.MustRegister(models.SRTIntervalReceiveRetransmittedPackets)
	r.MustRegister(models.SRTIntervalReceiveSentAckPackets)
	r.MustRegister(models.SRTIntervalReceiveSentNakPackets)

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)
	http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
}
