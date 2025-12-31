package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/proxy"
)

const (
	torProxyAddr = "127.0.0.1:9150" // Tor Browser için 9150
	timeout      = 60 * time.Second
	outputDir    = "output"
	reportFile   = "scan_report.log"
)

func main() {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal("Output klasörü oluşturulamadı:", err)
	}

	report, err := os.OpenFile(reportFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Report dosyası açılamadı:", err)
	}
	defer report.Close()
	logger := log.New(io.MultiWriter(os.Stdout, report), "", log.LstdFlags)

	filename := "targets.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	urls, err := readTargets(filename)
	if err != nil {
		log.Fatal("Hedef dosyası okunamadı:", err)
	}

	if len(urls) == 0 {
		logger.Println("[INFO] Hedef listesi boş!")
		return
	}

	logger.Printf("[INFO] %d adet adres yüklendi. Tarama başlıyor...\n", len(urls))

	client, err := createTorClient()
	if err != nil {
		log.Fatal("Tor client oluşturulamadı:", err)
	}

	// Tor kontrolü
	logger.Println("[INFO] Tor bağlantısı kontrol ediliyor...")
	testResp, err := client.Get("https://check.torproject.org")
	if err != nil {
		logger.Printf("[WARN] Tor kontrolü başarısız: %v\n", err)
	} else {
		testBody, _ := io.ReadAll(testResp.Body)
		testResp.Body.Close()
		if strings.Contains(string(testBody), "Congratulations") {
			logger.Println("[SUCCESS] Tor bağlantısı DOĞRULANDI! IP gizli.")
		} else {
			logger.Println("[WARN] Tor bağlantısı şüpheli.")
		}
	}
	logger.Println("")

	for _, rawURL := range urls {
		scanURL(rawURL, client, logger)
	}

	logger.Println("[INFO] Tarama tamamlandı!")
}

func readTargets(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.HasPrefix(line, "http") {
			line = "http://" + line
		}
		urls = append(urls, line)
	}
	return urls, scanner.Err()
}

func createTorClient() (*http.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", torProxyAddr, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	dialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}

	transport := &http.Transport{
		DialContext:           dialContext,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: timeout,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}, nil
}

func scanURL(rawURL string, client *http.Client, logger *log.Logger) {
	// 1. HTML ÇEKME KISMI
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		logger.Printf("[ERR] %s -> %v\n", rawURL, err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "close")

	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		logger.Printf("[ERR] %s -> Bağlantı hatası: %v (%.2fs)\n", rawURL, err, duration.Seconds())
		// Hata olsa bile screenshot denemek istersen buradan return kaldırırsın
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logger.Printf("[ERR] %s -> HTTP %d (%.2fs)\n", rawURL, resp.StatusCode, duration.Seconds())
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Printf("[ERR] %s -> Body okuma hatası: %v\n", rawURL, err)
			} else {
				host := sanitizeFilename(rawURL)
				htmlPath := filepath.Join(outputDir, host+".html")
				if err := os.WriteFile(htmlPath, body, 0644); err != nil {
					logger.Printf("[ERR] %s -> HTML yazma hatası: %v\n", rawURL, err)
				} else {
					logger.Printf("[SUCCESS] %s -> %d KB HTML kaydedildi (%.2fs)\n", rawURL, len(body)/1024, duration.Seconds())
				}
			}
		}
	}

	// 2. EKRAN GÖRÜNTÜSÜ KISMI (chromedp ile, Tor üzerinden)
	logger.Printf("[INFO] %s -> Ekran görüntüsü alınıyor...\n", rawURL)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer("socks5://"+torProxyAddr),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.WindowSize(1920, 1080), // daha gerçekçi görünüm
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	taskCtx, cancel = context.WithTimeout(taskCtx, 90*time.Second)
	defer cancel()

	var screenshotBuf []byte
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(rawURL),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.FullScreenshot(&screenshotBuf, 95), // kaliteli screenshot
	)

	host := sanitizeFilename(rawURL)

	if err != nil {
		logger.Printf("[ERR] %s -> Screenshot hatası: %v\n", rawURL, err)
		return
	}

	ssPath := filepath.Join(outputDir, host+".png")
	if err := os.WriteFile(ssPath, screenshotBuf, 0644); err != nil {
		logger.Printf("[ERR] %s -> Screenshot yazma hatası: %v\n", rawURL, err)
		return
	}

	logger.Printf("[SUCCESS] %s -> Screenshot kaydedildi (%d KB)\n", rawURL, len(screenshotBuf)/1024)
}

func sanitizeFilename(u string) string {
	parsed, err := url.Parse(u)
	if err != nil || parsed.Host == "" {
		// Parse hatası varsa manuel temizle
		u = strings.TrimPrefix(u, "http://")
		u = strings.TrimPrefix(u, "https://")
		return strings.NewReplacer(":", "_", "/", "_", "\\", "_").Replace(u)
	}
	host := strings.ReplaceAll(parsed.Host, ":", "_")
	return strings.ReplaceAll(host, "/", "_")
}