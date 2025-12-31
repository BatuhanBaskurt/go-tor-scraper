Markdown

# Go Tor Scraper 🕸️ 🛡️

**Go Tor Scraper**, yerel bir metin dosyasındaki domain listesini (URL) okuyan ve her bir adrese Tor ağının sunduğu tam anonimlik katmanı üzerinden HTTP istekleri gönderen yüksek performanslı bir Go aracıdır. Özellikle `.onion` servislerinin durumunu kontrol etmek ve kimlik gizleyerek toplu veri taraması yapmak için optimize edilmiştir.

---

## ✨ Özellikler

* **Toplu Dosya İşleme:** Yüzlerce satırlık `.txt` dosyalarını hızlıca parse eder.
* **Tor SOCKS5 Entegrasyonu:** Tüm ağ trafiğini otomatik olarak `127.0.0.1:9050` proxy hattına yönlendirir.
* **Onion Ağ Desteği:** Standart HTTP istemcilerinin erişemediği Tor Gizli Servislerine (.onion) sorunsuz erişim sağlar.
* **Hata Yönetimi:** Bağlantı zaman aşımı (timeout) veya erişilemeyen servisleri terminalde raporlar.
* **IP Maskeleme:** Hedef sunucular gerçek IP adresinizi asla göremez, sadece Tor çıkış düğümlerini (exit nodes) görür.

---

## 🚀 Kurulum

Projeyi çalıştırmadan önce sisteminizde **Go** (1.18+) ve aktif bir **Tor** servisinin kurulu olduğundan emin olun.

```bash
# Projeyi klonlayın
git clone [https://github.com/BatuhanBaskurt/go-tor-scraper.git](https://github.com/BatuhanBaskurt/go-tor-scraper.git)

# Proje dizinine girin
cd go-tor-scraper

# Bağımlılıkları yükleyin
go mod tidy
🖥️ Kullanım
Taramak istediğiniz domainleri her satıra bir tane gelecek şekilde bir metin dosyasına kaydedin. Ardından programı, dosya adını argüman olarak vererek çalıştırın:

Bash

go run main.go targets.txt
Not: Tor servisinin arka planda çalıştığından ve varsayılan 9050 portunun açık olduğundan emin olun.

🛠️ Teknik Altyapı
Bu araç, Go'nun standart net/http paketini golang.org/x/net/proxy kütüphanesi ile genişleterek Tor ağına bağlar.

Go

// Tor Proxy Yapılandırması
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
if err != nil {
    log.Fatal("Proxy bağlantısı kurulamadı!")
}

transport := &http.Transport{Dial: dialer.Dial}
client := &http.Client{Transport: transport}