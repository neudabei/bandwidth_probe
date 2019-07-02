package main

import (
  "io"
  "net/http"
  "os"
  "time"
  "fmt"
)

func main() {
  fileUrl := "http://ipv4.download.thinkbroadband.com/20MB.zip"

  for {
    err := DownloadFile("20MB.zip", fileUrl)
    if err != nil {
      panic(err)
    }
  }
}

func DownloadFile(filepath string, url string) error {

  start := time.Now()

  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  out, err := os.Create(filepath)
  if err != nil {
    return err
  }
  defer out.Close()

  _, err = io.Copy(out, resp.Body)

  elapsed := time.Since(start)
  fmt.Printf("Download speed %.2f Mb/s\n", DownloadSpeedCalculator(elapsed, filepath))

  return err
}

func DownloadSpeedCalculator(timeElapsed time.Duration, filepath string) float64 {
  result := fileSizeInMegaBits(filepath) / (float64(timeElapsed) / 1000000000)
  return result
}

func fileSizeInMegaBits(filepath string) float64 {
  fileInfo, _ := os.Stat(filepath)
  sizeInBytes := fileInfo.Size()
  sizeInMegaBits := float64(sizeInBytes) * 0.000008

  return sizeInMegaBits
}

