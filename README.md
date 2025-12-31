# Go Tor Scraper 🕸️ 🛡️

**Go Tor Scraper**, Tor ağının sunduğu anonimlikten yararlanarak web sitelerinden veri çekmenize olanak tanıyan, Go (Golang) ile geliştirilmiş hızlı ve güvenli bir araçtır. `.onion` uzantılı sitelere erişim sağlamak veya standart web kazıma işlemlerinde IP engellemelerinden kaçınmak için idealdir.

---

## ✨ Özellikler

* **Tor Entegrasyonu:** Tüm trafik Tor ağı üzerinden proxy edilerek anonimlik sağlanır.
* **Onion Desteği:** Standart tarayıcıların erişemediği `.onion` adreslerini kazıyabilir.
* **Eşzamanlılık (Concurrency):** Go'nun `goroutine` yapısı sayesinde yüksek hızda tarama.
* **IP Rotasyonu:** Tor üzerinden her istekte veya belirli aralıklarla yeni bir kimlik (IP) alma imkanı.
* **Kolay Kurulum:** Minimum bağımlılık ve hızlı konfigürasyon.

---

## 🚀 Başlangıç

Bu projeyi yerel makinenizde çalıştırmak için aşağıdaki adımları takip edin.

### 📋 Gereksinimler

* **Go:** (Sürüm 1.18 veya üzeri)
* **Tor Service:** Bilgisayarınızda bir Tor servisinin çalışıyor olması gerekir.
    * **macOS:** `brew install tor`
    * **Linux:** `sudo apt install tor`
    * **Windows:** [Tor Project](https://www.torproject.org/download/tor-browser-alpha/) üzerinden uzman paketini indirin.

### ⚙️ Kurulum

1.  **Projeyi klonlayın:**
    ```bash
    git clone [https://github.com/BatuhanBaskurt/go-tor-scraper.git](https://github.com/BatuhanBaskurt/go-tor-scraper.git)
    cd go-tor-scraper
    ```

2.  **Bağımlılıkları indirin:**
    ```bash
    go mod tidy
    ```

3.  **Tor Servisi:** Tor servisinin çalıştığından emin olun (Varsayılan port genellikle **9050**'dir).

### 🖥️ Kullanım

Uygulamayı çalıştırmak için terminale şu komutu girin:

```bash
go run main.go -url "[http://check.torproject.org](http://check.torproject.org)"
🛠️ Teknik Detaylar
Bu araç, HTTP isteklerini Tor SOCKS5 proxy hattına yönlendirmek için özel bir http.Client yapılandırması kullanır.

Go

// Örnek Proxy Bağlantı Yapısı
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
httpTransport := &http.Transport{Dial: dialer.Dial}
client := &http.Client{Transport: httpTransport}