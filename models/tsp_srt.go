package models

type TspSRT struct {
	Global  Global  `json:"global"`
	Receive Receive `json:"receive"`
	Send    Send    `json:"send"`
}

type Global struct {
	Instant struct {
		RTTMs float64 `json:"rtt-ms"`
	} `json:"instant"`
}

type Receive struct {
	Instant struct {
		AvgBelatedMs            int `json:"avg-belated-ms"`
		BufferAckBytes          int `json:"buffer-ack-bytes"`
		BufferAckMs             int `json:"buffer-ack-ms"`
		BufferAckPackets        int `json:"buffer-ack-packets"`
		BufferAvailBytes        int `json:"buffer-avail-bytes"`
		DeliveryDelayMs         int `json:"delivery-delay-ms"`
		MssBytes                int `json:"mss-bytes"`
		ReorderTolerancePackets int `json:"reorder-tolerance-packets"`
	} `json:"instant"`
	Interval struct {
		Bytes                     int     `json:"bytes"`
		DropBytes                 int     `json:"drop-bytes"`
		DroppedPackets            int     `json:"dropped-packets"`
		FilterExtraPackets        int     `json:"filter-extra-packets"`
		FilterNotRecoveredPackets int     `json:"filter-not-recovered-packets"`
		FilterRecoveredPackets    int     `json:"filter-recovered-packets"`
		IgnoredLatePackets        int     `json:"ignored-late-packets"`
		LossBytes                 int     `json:"loss-bytes"`
		LostPackets               int     `json:"lost-packets"`
		Packets                   int     `json:"packets"`
		RateMbps                  float64 `json:"rate-mbps"`
		ReorderDistancePackets    int     `json:"reorder-distance-packets"`
		RetransmittedPackets      int     `json:"retransmitted-packets"`
		SentAckPackets            int     `json:"sent-ack-packets"`
		SentNakPackets            int     `json:"sent-nak-packets"`
		UndecryptedBytes          int     `json:"undecrypted-bytes"`
		UndecryptedPackets        int     `json:"undecrypted-packets"`
		UniqueBytes               int     `json:"unique-bytes"`
		UniquePackets             int     `json:"unique-packets"`
	} `json:"interval"`
	Total struct {
		Bytes                     int `json:"bytes"`
		DropBytes                 int `json:"drop-bytes"`
		DroppedPackets            int `json:"dropped-packets"`
		ElapsedMs                 int `json:"elapsed-ms"`
		FilterExtraPackets        int `json:"filter-extra-packets"`
		FilterNotRecoveredPackets int `json:"filter-not-recovered-packets"`
		FilterRecoveredPackets    int `json:"filter-recovered-packets"`
		LossBytes                 int `json:"loss-bytes"`
		LostPackets               int `json:"lost-packets"`
		Packets                   int `json:"packets"`
		SentAckPackets            int `json:"sent-ack-packets"`
		SentNakPackets            int `json:"sent-nak-packets"`
		UndecryptedBytes          int `json:"undecrypted-bytes"`
		UndecryptedPackets        int `json:"undecrypted-packets"`
		UniqueBytes               int `json:"unique-bytes"`
		UniquePackets             int `json:"unique-packets"`
	} `json:"total"`
}

type Send struct {
	Instant struct {
		AvailBufferBytes           int     `json:"avail-buffer-bytes"`
		CongestionWindowPackets    int     `json:"congestion-window-packets"`
		DeliveryDelayMs            int     `json:"delivery-delay-ms"`
		EstimatedLinkBandwidthMbps float64 `json:"estimated-link-bandwidth-mbps"`
		FlowWindowPackets          int     `json:"flow-window-packets"`
		InFlightPackets            int     `json:"in-flight-packets"`
		IntervalPackets            int     `json:"interval-packets"`
		MaxBandwidthMbps           int     `json:"max-bandwidth-mbps"`
		MssBytes                   int     `json:"mss-bytes"`
		SndBufferBytes             int     `json:"snd-buffer-bytes"`
		SndBufferMs                int     `json:"snd-buffer-ms"`
		SndBufferPackets           int     `json:"snd-buffer-packets"`
	} `json:"instant"`
	Interval struct {
		Bytes              int `json:"bytes"`
		DropBytes          int `json:"drop-bytes"`
		DroppedPackets     int `json:"dropped-packets"`
		FilterExtraPackets int `json:"filter-extra-packets"`
		LostPackets        int `json:"lost-packets"`
		Packets            int `json:"packets"`
		ReceivedAckPackets int `json:"received-ack-packets"`
		ReceivedNakPackets int `json:"received-nak-packets"`
		RetransmitBytes    int `json:"retransmit-bytes"`
		RetransmitPackets  int `json:"retransmit-packets"`
		SendDurationUs     int `json:"send-duration-us"`
		SendRateMbps       int `json:"send-rate-mbps"`
		UniqueBytes        int `json:"unique-bytes"`
		UniquePackets      int `json:"unique-packets"`
	} `json:"interval"`
	Total struct {
		Bytes              int `json:"bytes"`
		DropBytes          int `json:"drop-bytes"`
		DroppedPackets     int `json:"dropped-packets"`
		ElapsedMs          int `json:"elapsed-ms"`
		FilterExtraPackets int `json:"filter-extra-packets"`
		LostPackets        int `json:"lost-packets"`
		Packets            int `json:"packets"`
		ReceivedAckPackets int `json:"received-ack-packets"`
		ReceivedNakPackets int `json:"received-nak-packets"`
		RestransBytes      int `json:"restrans-bytes"`
		RetransmitPackets  int `json:"retransmit-packets"`
		SendDurationUs     int `json:"send-duration-us"`
		UniqueBytes        int `json:"unique-bytes"`
		UniquePackets      int `json:"unique-packets"`
	} `json:"total"`
}
