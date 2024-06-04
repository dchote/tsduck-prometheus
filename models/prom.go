package models

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Transport Stream

var TsBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_bitrate_bytes",
		Help: "The overall TS bitrate based upon 188 byte TS packet.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPcrBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pcr_bitrate_bytes",
		Help: "The PCR bitrate of the TS based upon 188 byte TS packet.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPidBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_bitrate_bytes",
		Help: "The bitrate for an individual PID based upon a 188 byte TS packet.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidServiceCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_service_count_total",
		Help: "The total number of services within the a given PID in the TS.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidDiscontinuity = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_discontinuities_total",
		Help: "The discontinuities per PID.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidDuplicated = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_duplicated_total",
		Help: "The number of duplicated PIDs seen for a given PID since the start of monitoring.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidDTSLeap = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_dts_leap_total",
		Help: "The dts leaps per PID.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidPCRLeap = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_pcr_leap_total",
		Help: "The pcr leaps per PID.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidPTSLeap = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_pts_leap_total",
		Help: "The pts leaps per PID.",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "description"},
)

var TsServiceBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_service_bitrate_bytes",
		Help: "The bitrate of each service carried in the TS based upon a 188 byte TS packet.",
	},
	[]string{"target", "label", "service_id", "ts_id", "name", "provider", "type_name", "pcr_pid", "pmt_pid"},
)

var TsPidMinRepititionMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_min_repitition_ms",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal"},
)

var TsPidMaxRepititionMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_max_repitition_ms",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal"},
)

var TsPidMinRepititionPkt = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_min_repitition_pkt",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal"},
)

var TsPidMaxRepititionPkt = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_max_repitition_pkt",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal"},
)

var TsPidRepititionMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_repitition_ms",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal"},
)

var TsPidRepititionPkt = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_repitition_pkt",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal"},
)

var TsPacketInvalidSync = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_packet_invalid_sync_total",
		Help: "The number of invalid sync packets detected.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPacketSuspectIgnored = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_packet_suspect_ignored_total",
		Help: "The number of suspect packets ignored.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPacketTeiErrors = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_packet_tei_count_total",
		Help: "The number of transport errors detected.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPidCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_count_total",
		Help: "The total number of pids detected in the TS.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPcrPidCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_pcr_count_total",
		Help: "The total number of PCR pids detected in the TS.",
	},
	[]string{"target", "label", "ts_id"},
)

var TsPidUnferencedCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_unreferenced_count_total",
		Help: "The total number of unreferenced pids detected in the TS.",
	},
	[]string{"target", "label", "ts_id"},
)

// Connection

var ConnectAttempts = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "con_reconnect_attempts_total",
		Help: "The total number of connection attempts.",
	},
	[]string{"target", "label"},
)

// Continuity

var TsDiscontinuityTotals = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_discontinuity_totals",
		Help: "",
	},
	[]string{"target", "label", "pid", "pid_hexadecimal", "type"},
)

// SRT

var SRTRTTMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_rtt_ms",
		Help: "SRT Round-Trip Time in ms.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveBytes = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_bytes",
		Help: "The number of bytes passed in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveDropBytes = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_drop_bytes",
		Help: "The number of bytes dropped in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveDroppedPackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_dropped_packets",
		Help: "The number of packets dropped in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveDroppedPacketsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "srt_interval_receive_dropped_packets_total",
		Help: "The total number of packets dropped.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveIgnoredLatePackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_ignored_late_packets",
		Help: "The number of packets ignored in reporting interval due to being late.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveIgnoredLatePacketsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "srt_interval_receive_ignored_late_packets_total",
		Help: "The total number of packets ignored.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveLossBytes = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_loss_bytes",
		Help: "The number of bytes lost in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveLostPackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_lost_packets",
		Help: "The number of packets lost in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveLostPacketsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "srt_interval_receive_lost_packets_total",
		Help: "The total number of packets lost.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceivePackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_packets",
		Help: "The number of packets in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalRateMbps = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_rate",
		Help: "The transmit rate during the reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveReorderDistancePackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_reorder_distance_packets",
		Help: "The number of packets reordered in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveReorderDistancePacketsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "srt_interval_receive_reorder_distance_packets_total",
		Help: "The total number of packets reordered.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveRetransmittedPackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_retransmitted_packets",
		Help: "The number of packets retransmitted in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveRetransmittedPacketsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "srt_interval_receive_retransmitted_packets_total",
		Help: "The total number of packets retransmitted.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveSentAckPackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_sent_ack_packets",
		Help: "The number of ack packets in reporting interval.",
	},
	[]string{"target", "label"},
)

var SRTIntervalReceiveSentNakPackets = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "srt_interval_receive_sent_nak_packets",
		Help: "The number of nak packets in reporting interval.",
	},
	[]string{"target", "label"},
)
