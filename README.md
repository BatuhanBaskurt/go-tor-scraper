# Go Tor Scraper

Go ile yazılmış, Tor ağı üzerinden anonim HTTP istekleri gönderen yüksek performanslı bir araç.

Yerel bir metin dosyasında (.txt) listelenen domainleri (normal web veya .onion) sırayla tarar ve yanıtları işler.

## Temel Özellikler

- Tam anonimlik (SOCKS5 proxy ile Tor trafiği)
- .onion adreslerine doğrudan erişim
- Toplu domain tarama (yüzlerce satır destekler)
- Hafif ve hızlı (Go avantajı)
- Basit kullanım, karmaşık konfigürasyon yok

## Gereksinimler

- Go 1.18 veya üstü
- Çalışan Tor servisi (varsayılan port: 9050)

## Kurulum

```bash
git clone https://github.com/BatuhanBaskurt/go-tor-scraper.git
cd go-tor-scraper
go mod tidy

Kullanım Rehberi

Taramak istediğiniz domainleri targets.txt gibi bir dosyaya her satıra bir tane yazın.
Örnek targets.txt içeriği:

texthttp://example.com
http://example.onion
https://check.torproject.org

Çalıştırın:
Bashgo run main.go targets.txt

Not: Sadece dosya adını yazmanız yeterli, program içeriği satır satır otomatik okur.

Teknik Altyapı ve Çalışma Mantığı
Araç, Go'nun net/http kütüphanesini golang.org/x/net/proxy ile genişleterek Tor proxy'si üzerinden bağlantı kurar.

Go// Tor SOCKS5 Proxy Yapılandırması
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

Notlar

Tor ağı doğal olarak yavaş olabilir, özellikle .onion sitelerde.
Yoğun kullanımda Tor exit node'lar engellenebilir.
Yalnızca yasal ve etik amaçlarla kullanılmalıdır.

Batuhan Başkurt
https://github.com/BatuhanBaskurt