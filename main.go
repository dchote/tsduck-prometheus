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

	// leaps
	if pid.Packets.DTSLeap >= 1 {
		sum := 0
		for sum < pid.Packets.DTSLeap {
			sum += 1
			models.TsPidDTSLeap.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
		}
	}

	if pid.Packets.PCRLeap >= 1 {
		sum := 0
		for sum < pid.Packets.PCRLeap {
			sum += 1
			models.TsPidPCRLeap.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
		}
	}

	if pid.Packets.PTSLeap >= 1 {
		sum := 0
		for sum < pid.Packets.PTSLeap {
			sum += 1
			models.TsPidPTSLeap.WithLabelValues(target, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
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

func updateContinuityValues(target, label string, tspContinuity models.TspContinuity) {
	if tspContinuity.Packets >= 1 {
		sum := 0
		for sum < tspContinuity.Packets {
			sum += 1
			models.TsDiscontinuityTotals.WithLabelValues(target, label, fmt.Sprint(tspContinuity.Pid), strconv.FormatInt(int64(tspContinuity.Pid), 16), tspContinuity.Type).Inc()
		}
	}
}

func updatePCRExtractValues(target, label string, logLine string) {
	var PIDHex = logLine[13:16]
	var Type = logLine[19:22]

	var lastComma = strings.LastIndex(logLine, ",")
	var lastMS = strings.LastIndex(logLine, "ms")

	var IntervalMS, _ = strconv.ParseFloat(logLine[lastComma+2:lastMS-1], 64)

	//fmt.Printf("tsp pcrextract: TYPE %v PID %v Interval: %v\n", Type, PIDHex, IntervalMS)

	models.TSPidPCRInterval.WithLabelValues(target, label, PIDHex, Type).Set(float64(IntervalMS))
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

	if recieve.Interval.UniquePackets > 0 && recieve.Interval.UniqueBytes > 0 {
		models.SRTIntervalReceivePacketSize.WithLabelValues(target, label).Set(float64(recieve.Interval.UniqueBytes / recieve.Interval.UniquePackets))
	}

	if recieve.Interval.DroppedPackets >= 1 {
		sum := 0
		for sum < recieve.Interval.DroppedPackets {
			sum += 1
			models.SRTIntervalReceiveDroppedPacketsTotal.WithLabelValues(target, label).Inc()
		}
	}

	if recieve.Interval.IgnoredLatePackets >= 1 {
		sum := 0
		for sum < recieve.Interval.IgnoredLatePackets {
			sum += 1
			models.SRTIntervalReceiveIgnoredLatePacketsTotal.WithLabelValues(target, label).Inc()
		}
	}

	if recieve.Interval.LostPackets >= 1 {
		sum := 0
		for sum < recieve.Interval.LostPackets {
			sum += 1
			models.SRTIntervalReceiveLostPacketsTotal.WithLabelValues(target, label).Inc()
		}
	}

	if recieve.Interval.ReorderDistancePackets >= 1 {
		sum := 0
		for sum < recieve.Interval.ReorderDistancePackets {
			sum += 1
			models.SRTIntervalReceiveReorderDistancePacketsTotal.WithLabelValues(target, label).Inc()
		}
	}

	if recieve.Interval.RetransmittedPackets >= 1 {
		sum := 0
		for sum < recieve.Interval.RetransmittedPackets {
			sum += 1
			models.SRTIntervalReceiveRetransmittedPacketsTotal.WithLabelValues(target, label).Inc()
		}
	}
}

func launchTsp(target, label string) {
	command := ""

	// input protocol specific items
	if protocol == "srt" {
		command = "tsp -I srt --caller " + target + " --statistics-interval 1000 --json-line "
	} else if protocol == "rtsp" {
		command = "gst-launch-1.0 -q rtspsrc location=\"rtsp://" + target + "\" ! rtpmp2tdepay ! filesink location=/dev/stdout | tsp -r "
	}

	// analysis items
	command += "-P analyze -i 1 --json-line -P continuity --json-line -P pcrextract -l -O drop"

	fmt.Printf("Starting %v\n", command)
	cmd := exec.Command("bash", "-c", command)

	// TSDuck outputs to stderr
	cmdReader, err := cmd.StderrPipe()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Started monitoring for %v\n", label)
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("error starting %v: %v\n", command, err)
		return
	}

	// Create buffer to read the stderr output
	scanner := bufio.NewScanner(cmdReader)

	var tspAnalyze models.TspAnalyze
	var tspSRT models.TspSRT
	var tspContinuity models.TspContinuity

	var analyzeMatch = "* analyze: "
	var srtMatch = "* srt: "
	var continuityMatch = "* continuity: "
	var pcrMatch = "* pcrextract: "

	for scanner.Scan() {
		s := scanner.Text()

		if len(s) > len(analyzeMatch) && s[:len(analyzeMatch)] == analyzeMatch {
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

		} else if len(s) > len(srtMatch) && s[:len(srtMatch)] == srtMatch {
			// capture srt statistics output
			t := strings.Replace(s, srtMatch, "", -1)
			// Unmarshal JSON into TSP SRT model
			json.Unmarshal([]byte(t), &tspSRT)

			go updateSRTGlobalValues(target, label, tspSRT.Global)
			go updateSRTRecieveValues(target, label, tspSRT.Receive)
			//fmt.Printf("tsp srt %v\n", tspSRT)
		} else if len(s) > len(continuityMatch) && s[:len(continuityMatch)] == continuityMatch {
			// capture output from continuity function

			t := strings.Replace(s, continuityMatch, "", -1)
			// Unmarshal JSON into TSP Continuity model
			json.Unmarshal([]byte(t), &tspContinuity)

			//fmt.Printf("tsp continuity %v\n", tspContinuity)

			go updateContinuityValues(target, label, tspContinuity)
		} else if len(s) > len(pcrMatch) && s[:len(pcrMatch)] == pcrMatch {
			// capture output from pcrextract function
			var logLine = s[len(pcrMatch):len(s)]

			//fmt.Printf("tsp pcrextract %v\n", logLine)

			go updatePCRExtractValues(target, label, logLine)
		}

	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("%v exited with error: %v\n", command, err)
	}

}

var (
	target     string
	label      string
	listenPort int
	protocol   string
)

func cliArguments() {
	usage := `
Usage: tsduck-prometheus [options]

Options:
  -t, --target=<host:port>      Specify target host:port for SRT listener
  -l, --label=<label>           Specify label to use in Prometheus
  -p, --port=<port>             Specify listen port [default: 8000]
  -P, --protocol=<protocol>     Specify protocol [default: srt]
  -h, --help                    Show this screen.
`
	args, _ := docopt.ParseArgs(usage, os.Args[1:], "")

	target, _ = args.String("--target")
	label, _ = args.String("--label")
	listenPort, _ = args.Int("--port")
	protocol, _ = args.String("--protocol")

	if target == "" || label == "" {
		docopt.PrintHelpAndExit(nil, usage)
	}
}

func main() {
	cliArguments()

	go func() {
		for {
			models.ConnectAttempts.WithLabelValues(target, label).Inc()

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
	r.MustRegister(models.TsPidDTSLeap)
	r.MustRegister(models.TsPidPCRLeap)
	r.MustRegister(models.TsPidPTSLeap)

	// connection
	r.MustRegister(models.ConnectAttempts)

	// continuity
	r.MustRegister(models.TsDiscontinuityTotals)

	// pcr extract
	r.MustRegister(models.TSPidPCRInterval)

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

	r.MustRegister(models.SRTIntervalReceiveDroppedPacketsTotal)
	r.MustRegister(models.SRTIntervalReceiveIgnoredLatePacketsTotal)
	r.MustRegister(models.SRTIntervalReceiveLostPacketsTotal)
	r.MustRegister(models.SRTIntervalReceiveReorderDistancePacketsTotal)
	r.MustRegister(models.SRTIntervalReceiveRetransmittedPacketsTotal)

	r.MustRegister(models.SRTIntervalReceivePacketSize)

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)
	http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
}
