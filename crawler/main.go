package main

import (
	"fmt"
	"github.com/andelf/go-curl"
	"log"
	"time"
)

type CurlStats struct {
	SizeDownload      int64
	SizeUpload        int64
	SpeedDownload     float64
	SpeedUpload       float64
	TimeAppConnect    time.Duration
	TimeConnect       time.Duration
	TimeNameLookup    time.Duration
	TimePreTransfer   time.Duration
	TimeRedirect      time.Duration
	TimeStartTransfer time.Duration
	TimeTotal         time.Duration
}

func curlStat(url, iface string) (CurlStats, error) {
	easy := curl.EasyInit()
	if easy == nil {
		return CurlStats{}, fmt.Errorf("failed to initialize curl")
	}
	defer easy.Cleanup()

	if err := easy.Setopt(curl.OPT_URL, url); err != nil {
		return CurlStats{}, fmt.Errorf("failed to set URL: %v", err)
	}

	if iface != "" {
		if err := easy.Setopt(curl.OPT_INTERFACE, iface); err != nil {
			return CurlStats{}, fmt.Errorf("failed to bind to interface %s: %v", iface, err)
		}
	}

	if err := easy.Setopt(curl.OPT_VERBOSE, true); err != nil {
		return CurlStats{}, fmt.Errorf("failed to set verbose: %v", err)
	}

	if err := easy.Setopt(curl.OPT_TIMEOUT, 30); err != nil {
		return CurlStats{}, fmt.Errorf("failed to set timeout: %v", err)
	}

	var responseBody []byte
	if err := easy.Setopt(curl.OPT_WRITEFUNCTION, func(buf []byte, userdata interface{}) bool {
		responseBody = append(responseBody, buf...)
		return true
	}); err != nil {
		return CurlStats{}, fmt.Errorf("failed to set write function: %v", err)
	}
	if err := easy.Perform(); err != nil {
		return CurlStats{}, fmt.Errorf("failed to perform request: %v", err)
	}

	stats := CurlStats{}
	if val, err := easy.Getinfo(curl.INFO_SIZE_DOWNLOAD); err == nil {
		stats.SizeDownload = int64(val.(float64))
	}
	if val, err := easy.Getinfo(curl.INFO_SIZE_UPLOAD); err == nil {
		stats.SizeUpload = int64(val.(float64))
	}
	if val, err := easy.Getinfo(curl.INFO_SPEED_DOWNLOAD); err == nil {
		stats.SpeedDownload = val.(float64)
	}
	if val, err := easy.Getinfo(curl.INFO_SPEED_UPLOAD); err == nil {
		stats.SpeedUpload = val.(float64)
	}
	if val, err := easy.Getinfo(curl.INFO_APPCONNECT_TIME); err == nil {
		stats.TimeAppConnect = time.Duration(val.(float64) * float64(time.Second))
	}
	if val, err := easy.Getinfo(curl.INFO_CONNECT_TIME); err == nil {
		stats.TimeConnect = time.Duration(val.(float64) * float64(time.Second))
	}
	if val, err := easy.Getinfo(curl.INFO_NAMELOOKUP_TIME); err == nil {
		stats.TimeNameLookup = time.Duration(val.(float64) * float64(time.Second))
	}
	if val, err := easy.Getinfo(curl.INFO_PRETRANSFER_TIME); err == nil {
		stats.TimePreTransfer = time.Duration(val.(float64) * float64(time.Second))
	}
	if val, err := easy.Getinfo(curl.INFO_REDIRECT_TIME); err == nil {
		stats.TimeRedirect = time.Duration(val.(float64) * float64(time.Second))
	}
	if val, err := easy.Getinfo(curl.INFO_STARTTRANSFER_TIME); err == nil {
		stats.TimeStartTransfer = time.Duration(val.(float64) * float64(time.Second))
	}
	if val, err := easy.Getinfo(curl.INFO_TOTAL_TIME); err == nil {
		stats.TimeTotal = time.Duration(val.(float64) * float64(time.Second))
	}

	return stats, nil
}

func main() {
	url := "https://bing.com"
	iface := "wlp3s0"

	stats, err := curlStat(url, iface)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Curl Stats for %s (interface: %s):\n", url, iface)
	fmt.Printf("Size Download: %d bytes\n", stats.SizeDownload)
	fmt.Printf("Size Upload: %d bytes\n", stats.SizeUpload)
	fmt.Printf("Speed Download: %.2f bytes/sec\n", stats.SpeedDownload)
	fmt.Printf("Speed Upload: %.2f bytes/sec\n", stats.SpeedUpload)
	fmt.Printf("Time AppConnect: %v\n", stats.TimeAppConnect)
	fmt.Printf("Time Connect: %v\n", stats.TimeConnect)
	fmt.Printf("Time NameLookup: %v\n", stats.TimeNameLookup)
	fmt.Printf("Time PreTransfer: %v\n", stats.TimePreTransfer)
	fmt.Printf("Time Redirect: %v\n", stats.TimeRedirect)
	fmt.Printf("Time StartTransfer: %v\n", stats.TimeStartTransfer)
	fmt.Printf("Time Total: %v\n", stats.TimeTotal)
}
