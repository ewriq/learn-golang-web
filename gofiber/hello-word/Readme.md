### GoFiber ile hello word yazmak

**main.go diye bir dosya oluşturuyoruz ve terminale sırayla 
`go mod init GolangWeb`
`go get github.com/gofiber/fiber/v2` 
yazın.**

Sonra Main.go da
fiber paketmizi import ediyoruz ve app adında yeni fiber uygulaması oluşturuyoruz
```go
package main
import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()
}
```

Sonrasında Method ve Route tanımlamamız lazım
Gofiber da kullanabileceğimiz Methodlar
- **`Get`** <br>
`GET` methodu, belirtilen kaynağın bir temsilini ister. `GET` kullanan istekler yalnızca veri almalıdır.

- `Head` <br>
`HEAD` methodu, yanıt gövdesi olmadan `GET` isteği ile aynı olan bir yanıt talep eder.

- `Put` <br>
`PUT` methodu, hedef kaynağın tüm mevcut temsillerini istek yükü ile değiştirir

- `Post` <br>
`POST` methodu, belirtilen kaynağa bir varlık sunar ve genellikle sunucuda bir durum değişikliğine veya yan etkilere neden olur

- `Delete` <br>
`DELETE` methodu belirtilen kaynağı siler.

- `Connect` <br>
`CONNECT` methodu, hedef kaynak tarafından tanımlanan sunucuya bir tünel kurar

- `Options` <br>
`OPTIONS` methodu, hedef kaynak için iletişim seçeneklerini açıklar

- `Trace` <br>
`TRACE` methodu, hedef kaynağa giden yol boyunca bir mesaj geri dönüş testi gerçekleştirir."

- `Patch` <br>
`PATCH` methodu bir kaynağa kısmi değişiklikler uygular.

**Biz Hello world örneğimiz de GET methodunu kullanıcaz, route'umuz da "/" olacak
bu rotaya girenler bir json görecekler
"/"'ı index.html gibi düşünebilirsiniz**

```go
package main
import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mesaj":"merhaba dünya" //Sonucu Json olarak gönderir
		}) 
		/*
			SAYFA SONUCU
			{ "mesaj":"merhaba dünya" }
		*/
	})

	app.Listen("127.0.0.1:3000") //Servisin çalışacağı port
}
```

### Okumanız gerekenler
- [HTTP Methodları](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)
- [Gofiber Docs (Context)](https://docs.gofiber.io/api/ctx)
- [Gofiber Docs (App)](https://docs.gofiber.io/api/app)






 
