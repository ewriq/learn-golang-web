# GoFiber ile Router 

- **Dosyaları  ve klasörleri belirtildigi gibi oluşturunuz.**

- **Aşşagıdaki işlem ile routeri belirtiriz yani yönlendiririz yani web sitemizin maini orası olur yani yönlendiricisi orası olur (Routers.Initalize kısmı)**

```go
func main() {
	app := fiber.New() // Fiber uygulaması oluşturur

	Routers.Initalize(app) //Yönendiriciye oluşturulan uygulama verilir

	app.Listen(":8080") //Sunucu 8080 portunda başlatılır 
}
```

