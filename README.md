Önceki görüntülerde yaşadığın kayma sorunu, metnin içindeki Bash ve Go gibi etiketlerin kod bloklarının dışına çıkmasından kaynaklanıyordu. Aşağıdaki metni olduğu gibi, hiçbir satırı değiştirmeden kopyalayıp README.md dosyasına yapıştırırsan, GitHub üzerinde profesyonel, uzun ve hatasız bir dökümanın olacak.

Markdown

# Go Tor Scraper 🕸️ 🛡️

**Go Tor Scraper**, yerel bir metin dosyasında listelenen domain adreslerini (standart web veya .onion) okuyan ve her birine Tor ağının anonimlik katmanı üzerinden güvenli HTTP istekleri gönderen, yüksek performanslı bir Go (Golang) aracıdır.



---

## ✨ Temel Özellikler

* **Tam Anonimlik:** Tüm ağ trafiği SOCKS5 proxy (`127.0.0.1:9050`) üzerinden Tor ağına yönlendirilir, gerçek IP adresiniz gizlenir.
* **Toplu İşleme Kapasitesi:** Tek bir `.txt` dosyası içerisindeki yüzlerce domaini sırayla ve verimli bir şekilde tarar.
* **Onion Ağ Erişimi:** Standart HTTP istemcilerinin ulaşamadığı Tor Gizli Servislerine (`.onion`) doğrudan erişim sağlar.
* **Hafif ve Hızlı:** Go dilinin düşük sistem kaynağı tüketimi ve hız avantajını kullanarak optimize edilmiştir.
* **Kolay Kullanım:** Karmaşık konfigürasyonlarla uğraşmadan, sadece dosya yolunu belirterek çalıştırılabilir.

---

## 🚀 Kurulum ve Hazırlık

Projeyi çalıştırmadan önce sisteminizde **Go (1.18+)** ve aktif bir **Tor** servisinin kurulu olduğundan emin olmalısınız.

### 📋 Gereksinimler

1.  **Tor Servisi:** Tor daemon arka planda çalışıyor olmalıdır (Varsayılan port: `9050`).
2.  **Go:** Geliştirme ortamınızda Go kurulu olmalıdır.

### ⚙️ Adımlar

```bash
# Projeyi GitHub'dan klonlayın
git clone [https://github.com/BatuhanBaskurt/go-tor-scraper.git](https://github.com/BatuhanBaskurt/go-tor-scraper.git)

# Proje klasörüne geçiş yapın
cd go-tor-scraper

# Gerekli bağımlılıkları indirin
go mod tidy
🖥️ Kullanım Rehberi
Taramak istediğiniz domainleri (örneğin targets.txt) her satıra bir tane gelecek şekilde hazırlayın. Ardından programı aşağıdaki şekilde çalıştırın:

Bash

go run main.go targets.txt
Önemli: Komutun sonuna sadece hedef dosyanın adını yazmanız yeterlidir. Program, dosya içeriğini otomatik olarak satır satır işleyecektir.

🛠️ Teknik Altyapı ve Çalışma Mantığı
Bu araç, Go'nun standart net/http kütüphanesini golang.org/x/net/proxy paketiyle genişleterek bir proxy dialer oluşturur. Bu yapı, tüm HTTP trafiğini yerel Tor portuna tüneller.

Go

// Tor SOCKS5 Proxy Yapılandırması
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
if err != nil {
    log.Fatal("Tor proxy bağlantısı kurulamadı: ", err)
}

// Özel Transport ve Client oluşturma
transport := &http.Transport{Dial: dialer.Dial}
client := &http.Client{
    Transport: transport,
    Timeout:   time.Second * 30, // 30 saniye timeout süresi
}