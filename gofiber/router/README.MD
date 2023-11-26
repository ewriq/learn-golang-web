# GoFiber ile Router 
### Router nedir
-  **Router projelerimizi daha okunabilir hale getirir karmaşıklıgı düzeltir**

### Nasıl yapılır Nedir 

- **Dosyaları  ve klasörleri belirtildigi gibi oluşturunuz.**

- **Aşşagıdaki işlem ile routeri belirtiriz yani yönlendiririz yani web sitemizin maini orası olur yani yönlendiricisi orası olur (Routers.Initalize kısmı)**

```go
func main() {
	app := fiber.New() // Fiber uygulaması oluşturur

	Routers.Initalize(app) //Yönendiriciye oluşturulan uygulama verilir

	app.Listen(":8080") //Sunucu 8080 portunda başlatılır 
}
```

- **Handlerden home kısmını alır pathnamesi / olan get istegine verir**
```go
func Initalize(app *fiber.App){

    app.Get("/", Handler.Home) //Home import edilir pathname / olarak belirlenir 
	
}
```
- **Home fonksüyonu ise düz yazdıgımız gibi koyulur**
```go
func Home(c *fiber.Ctx) error {
	c.JSON(fiber.Map{ //Json olarak  { "hello":"word" } döndürülür
		"hello":"world",
	})
	return nil
}
```

