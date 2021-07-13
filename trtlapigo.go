package trtlapigo

import "fmt"

// Layer 0

type Response struct {
	Config    Config    `json:"config"`
	Lastblock Lastblock `json:"lastblock"`
	Network   Network   `json:"network"`
	Health    Health    `json:"health"`
	Pool      Pool      `json:"pool"`
}

// Layer 1

type Config struct {
	Ports []Port `json:"ports"`
}

type Lastblock struct {
	Difficulty int    `json:"difficulty"`
	Height     int    `json:"height"`
	Timestamp  int64  `json:"timestamp"`
	Reward     int    `json:"reward"`
	Hash       string `json:"hash"`
}

type Network struct {
	Difficulty int `json:"difficulty"`
	Height     int `json:"height"`
}

type Health struct {
	TurtleCoin TurtleCoin `json:"TurtleCoin"`
}

type Pool struct {
	Stats              Stats    `json:"stats"`
	Blocks             []string `json:"blocks"`
	TotalBlock         int      `json:"totalBlocks"`
	TotalBlockSolo     int      `json:"totalBlocksSolo"`
	TotalDiff          int64    `json:"totalDiff"`
	TotalDiffSolo      int64    `json:"totalDiffSolo"`
	TotalShares        int64    `json:"totalShares"`
	TotalSharesSolo    int64    `json:"totalSharesSolo"`
	Payments           []string `json:"payments"`
	TotalPayments      int      `json:"totalPayments"`
	TotalMinersPaid    int      `json:"totalMinersPaid"`
	Miners             int      `json:"miners"`
	MinersSolo         int      `json:"minersSolo"`
	Workers            int      `json:"workers"`
	Hashrate           int      `json:"hashrate"`
	HashrateSolo       int      `json:"hashrateSolo"`
	RoundScore         int64    `json:"roundScore"`
	RoundScoreSolo     int64    `json:"roundScoreSolo"`
	LastBlockFound     int64    `json:"lastBlockFound"`
	LastBlockFoundSolo int64    `json:"lastBlockFoundSolo"`
}

// Layer 2

type Port struct {
	Port       int    `json:"port"`
	Difficulty int    `json:"difficulty"`
	Desc       string `json:"desc"`
}

type TurtleCoin struct {
	Daemon string `json:"daemon"`
	Wallet string `json:"wallet"`
}

type Stats struct {
	LastBlockProp int64 `json:"lastBlockFoundprop"`
	LastBlockSolo int64 `json:"lastBlockFoundsolo"`
}

func init() {
	fmt.Println("trtlapigo initialized...")
}
