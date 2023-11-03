### GoFiber ile hello word yazmak

**Main fonksiyonumuza şu kodu ekliyoruz**

```go
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello word")
		return c.JSON(fiber.Map{
          "hello":"word",
		})
	})
	
   app.Listen(":8080")
```
### Talimatlar :

- İlk önce app adında yeni fiber uygulaması oluşturuyoruz 
- Daha sonra app.Get ekini kullanarak / yolunu açıyoruz daha sonra c ve ctx i tanımlayıp içerisine json veri dönderecegi için c.JSON ekini kullanıp ekrana hello word yazıyoruz
- Son olarak web sitemizi :8080 portunda aktif hale getiriyoruz
