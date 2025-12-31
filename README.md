# Go Tor Scraper 🕸️ 🛡️

**Go Tor Scraper**, bir metin dosyasındaki domain listesini alarak her birine Tor ağının anonimliği üzerinden HTTP istekleri atan hızlı bir araçtır. `.onion` adreslerini toplu kontrol etmek için idealdir.

---

## ✨ Özellikler

* **Toplu Tarama:** Dosyadan domain listesini otomatik okur.
* **Tor Proxy:** Tüm trafik SOCKS5 üzerinden (`127.0.0.1:9050`) anonim geçer.
* **Onion Erişimi:** Standart tarayıcıların giremediği gizli servisleri tarar.
* **Performans:** Go ile optimize edilmiş bağlantı yönetimi.

---

## 🚀 Başlangıç

### 📋 Gereksinimler

1.  **Tor Servisi:** Bilgisayarında Tor çalışıyor olmalı (Varsayılan Port: `9050`).
2.  **Go:** 1.18+ sürümü.

### ⚙️ Kurulum

```bash
git clone https://github.com/BatuhanBaskurt/go-tor-scraper.git
cd go-tor-scraper
go mod tidy
🖥️ Kullanım
Taramak istediğin domainleri bir dosyaya (örneğin targets.txt) alt alta ekle, sonra direkt şu şekilde çalıştır:

Bash

go run main.go targets.txt
Not: Dosya adını komutun sonuna boşluk bırakarak yazman yeterlidir. Program dosyadaki tüm adresleri sırayla tarayacaktır.

🛠️ Teknik Detay
Program, Go'nun net/http kütüphanesini Tor'un SOCKS5 proxy'sine bağlayarak çalışır. Bu sayede her istekte gerçek IP adresin gizli tutulur.

Go

// SOCKS5 üzerinden anonim bağlantı yapısı
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)