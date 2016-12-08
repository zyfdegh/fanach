package entity

// Stats contains usage infomation of a container
type Stats struct {
	RxBytes uint64 `json:"rxBytes"`
	TxBytes uint64 `json:"txBytesw"`
}
