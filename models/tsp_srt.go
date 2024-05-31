package models

type TspSRT struct {
	Global struct {
		Instant struct {
			RttMs float64 `json:"rtt-ms,omitempty"`
		} `json:"instant,omitempty"`
	} `json:"global,omitempty"`
	Receive struct {
		Instant struct {
			AvgBelatedMs            int `json:"avg-belated-ms,omitempty"`
			BufferAckBytes          int `json:"buffer-ack-bytes,omitempty"`
			BufferAckMs             int `json:"buffer-ack-ms,omitempty"`
			BufferAckPackets        int `json:"buffer-ack-packets,omitempty"`
			BufferAvailBytes        int `json:"buffer-avail-bytes,omitempty"`
			DeliveryDelayMs         int `json:"delivery-delay-ms,omitempty"`
			MssBytes                int `json:"mss-bytes,omitempty"`
			ReorderTolerancePackets int `json:"reorder-tolerance-packets,omitempty"`
		} `json:"instant,omitempty"`
		Interval struct {
			Bytes                     int     `json:"bytes,omitempty"`
			DropBytes                 int     `json:"drop-bytes,omitempty"`
			DroppedPackets            int     `json:"dropped-packets,omitempty"`
			FilterExtraPackets        int     `json:"filter-extra-packets,omitempty"`
			FilterNotRecoveredPackets int     `json:"filter-not-recovered-packets,omitempty"`
			FilterRecoveredPackets    int     `json:"filter-recovered-packets,omitempty"`
			IgnoredLatePackets        int     `json:"ignored-late-packets,omitempty"`
			LossBytes                 int     `json:"loss-bytes,omitempty"`
			LostPackets               int     `json:"lost-packets,omitempty"`
			Packets                   int     `json:"packets,omitempty"`
			RateMbps                  float64 `json:"rate-mbps,omitempty"`
			ReorderDistancePackets    int     `json:"reorder-distance-packets,omitempty"`
			RetransmittedPackets      int     `json:"retransmitted-packets,omitempty"`
			SentAckPackets            int     `json:"sent-ack-packets,omitempty"`
			SentNakPackets            int     `json:"sent-nak-packets,omitempty"`
			UndecryptedBytes          int     `json:"undecrypted-bytes,omitempty"`
			UndecryptedPackets        int     `json:"undecrypted-packets,omitempty"`
			UniqueBytes               int     `json:"unique-bytes,omitempty"`
			UniquePackets             int     `json:"unique-packets,omitempty"`
		} `json:"interval,omitempty"`
		Total struct {
			Bytes                     int `json:"bytes,omitempty"`
			DropBytes                 int `json:"drop-bytes,omitempty"`
			DroppedPackets            int `json:"dropped-packets,omitempty"`
			ElapsedMs                 int `json:"elapsed-ms,omitempty"`
			FilterExtraPackets        int `json:"filter-extra-packets,omitempty"`
			FilterNotRecoveredPackets int `json:"filter-not-recovered-packets,omitempty"`
			FilterRecoveredPackets    int `json:"filter-recovered-packets,omitempty"`
			LossBytes                 int `json:"loss-bytes,omitempty"`
			LostPackets               int `json:"lost-packets,omitempty"`
			Packets                   int `json:"packets,omitempty"`
			SentAckPackets            int `json:"sent-ack-packets,omitempty"`
			SentNakPackets            int `json:"sent-nak-packets,omitempty"`
			UndecryptedBytes          int `json:"undecrypted-bytes,omitempty"`
			UndecryptedPackets        int `json:"undecrypted-packets,omitempty"`
			UniqueBytes               int `json:"unique-bytes,omitempty"`
			UniquePackets             int `json:"unique-packets,omitempty"`
		} `json:"total,omitempty"`
	} `json:"receive,omitempty"`
	Send struct {
		Instant struct {
			AvailBufferBytes           int     `json:"avail-buffer-bytes,omitempty"`
			CongestionWindowPackets    int     `json:"congestion-window-packets,omitempty"`
			DeliveryDelayMs            int     `json:"delivery-delay-ms,omitempty"`
			EstimatedLinkBandwidthMbps float64 `json:"estimated-link-bandwidth-mbps,omitempty"`
			FlowWindowPackets          int     `json:"flow-window-packets,omitempty"`
			InFlightPackets            int     `json:"in-flight-packets,omitempty"`
			IntervalPackets            int     `json:"interval-packets,omitempty"`
			MaxBandwidthMbps           int     `json:"max-bandwidth-mbps,omitempty"`
			MssBytes                   int     `json:"mss-bytes,omitempty"`
			SndBufferBytes             int     `json:"snd-buffer-bytes,omitempty"`
			SndBufferMs                int     `json:"snd-buffer-ms,omitempty"`
			SndBufferPackets           int     `json:"snd-buffer-packets,omitempty"`
		} `json:"instant,omitempty"`
		Interval struct {
			Bytes              int `json:"bytes,omitempty"`
			DropBytes          int `json:"drop-bytes,omitempty"`
			DroppedPackets     int `json:"dropped-packets,omitempty"`
			FilterExtraPackets int `json:"filter-extra-packets,omitempty"`
			LostPackets        int `json:"lost-packets,omitempty"`
			Packets            int `json:"packets,omitempty"`
			ReceivedAckPackets int `json:"received-ack-packets,omitempty"`
			ReceivedNakPackets int `json:"received-nak-packets,omitempty"`
			RetransmitBytes    int `json:"retransmit-bytes,omitempty"`
			RetransmitPackets  int `json:"retransmit-packets,omitempty"`
			SendDurationUs     int `json:"send-duration-us,omitempty"`
			SendRateMbps       int `json:"send-rate-mbps,omitempty"`
			UniqueBytes        int `json:"unique-bytes,omitempty"`
			UniquePackets      int `json:"unique-packets,omitempty"`
		} `json:"interval,omitempty"`
		Total struct {
			Bytes              int `json:"bytes,omitempty"`
			DropBytes          int `json:"drop-bytes,omitempty"`
			DroppedPackets     int `json:"dropped-packets,omitempty"`
			ElapsedMs          int `json:"elapsed-ms,omitempty"`
			FilterExtraPackets int `json:"filter-extra-packets,omitempty"`
			LostPackets        int `json:"lost-packets,omitempty"`
			Packets            int `json:"packets,omitempty"`
			ReceivedAckPackets int `json:"received-ack-packets,omitempty"`
			ReceivedNakPackets int `json:"received-nak-packets,omitempty"`
			RestransBytes      int `json:"restrans-bytes,omitempty"`
			RetransmitPackets  int `json:"retransmit-packets,omitempty"`
			SendDurationUs     int `json:"send-duration-us,omitempty"`
			UniqueBytes        int `json:"unique-bytes,omitempty"`
			UniquePackets      int `json:"unique-packets,omitempty"`
		} `json:"total,omitempty"`
	} `json:"send,omitempty"`
}
