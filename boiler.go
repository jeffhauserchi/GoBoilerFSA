package main

   import "github.com/kataras/iris"
   import "github.com/kataras/iris/middleware/logger"

   //import _ "github.com/lib/pq"
   import "github.com/jinzhu/gorm"
   import _ "github.com/jinzhu/gorm/dialects/postgres"

   import "github.com/googollee/go-socket.io"

   import "fmt"
   import "os"


   ///////////////////////////////////////////////////////////
   // orm model
   ///////////////////////////////////////////////////////////
   type User struct {
     gorm.Model
     Email string `gorm:"type varchar(255);unique_index"`
     Password string `gorm:"type varchar(255)"`
     Salt string `gorm:"type varchar(255)"`
   }


   func main() {

   ///////////////////////////////////////////////////
   // GORM - Go ORM - connect and create a table
   //////////////////////////////////////////////////

   databaseName := "goboiler"

   db,err := gorm.Open("postgres","host=localhost port=5432 sslmode=disable dbname="+databaseName)

   if (err!=nil) {
     fmt.Printf("err=",err)
     os.Exit(-1)
   }
   defer db.Close()

   db.AutoMigrate(&User{})

   ///////////////////////////////////////////////////////////////
   // router initialization and logger init
   //////////////////////////////////////////////////////////////

    app := iris.Default()

    customLogger := logger.New(logger.Config{
		Status: true,
		IP: true,
		Method: true,
		Path: true,
		Query: true,
		//Columns: true,
	})
    app.Use(customLogger)


    ///////////////////////////////////////////////////
    // socket init and connection with router
    //////////////////////////////////////////////////

     server, err := socketio.NewServer(nil)
     if err != nil {
	    app.Logger().Fatal(err)
     }

     server.On("connection", func(so socketio.Socket) {
	    app.Logger().Infof("on connection")
     so.On("test event", func(msg string) {
            app.Logger().Infof("received test event %v",msg)
     })
     so.On("disconnection", func() {
            app.Logger().Infof("on disconnect")
	    })
     })

     server.On("error", func(so socketio.Socket, err error) {
            app.Logger().Errorf("error: %v", err)
     })

     app.Any("/socket.io/{p:path}", iris.FromStd(server))


    ///////////////////////////////////////////////////////////////
    // set up router to handle static
    /////////////////////////////////////////////////////////////

     pathToStatic := os.Getenv("BOILERMAKER_PUBLIC")

     app.StaticWeb("/", pathToStatic )

     app.Any("*", func(ctx iris.Context) {
        file:= pathToStatic+"index.html";
        ctx.ServeFile(file,false)
     })

     // Method:   GET
     // Resource: http://localhost:8080/hello
     //     app.Get("/hello", func(ctx iris.Context) {
     //       ctx.JSON(iris.Map{"message": "Hello iris web framework."})
     //     })

     //////////////////////////////////////////////////////////////
     // set up listener
     //////////////////////////////////////////////////////////////
     app.Run(iris.Addr(":8080"))
   }
