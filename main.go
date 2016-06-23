// main
package main

import (
	//"encoding/json"
	"fmt"
	//	"log"
	//	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsdutka/ps/controllers"
	"github.com/vsdutka/ps/models"
	"github.com/vsdutka/ps/shared/database"
	//	"github.com/vsdutka/wlog"
)

func handleRequests() {
	database.Open("./psdb.db")
	defer database.Close()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Brand{})
	database.DB.AutoMigrate(&models.Model{})
	database.DB.AutoMigrate(&models.Dealer{})
	database.DB.AutoMigrate(&models.Person{})
	database.DB.AutoMigrate(&models.Car{})
	database.DB.AutoMigrate(&models.WorkType{})

	router := gin.Default()
	router.Use(gin.Logger())
	// Listen and server on 0.0.0.0:8080
	router.POST("/api/v1/users/*action", controllers.UserController)
	router.POST("/api/v1/brands/*action", controllers.BrandController)
	router.POST("/api/v1/models/*action", controllers.ModelController)
	router.POST("/api/v1/dealers/*action", controllers.DealerController)
	router.POST("/api/v1/persons/*action", controllers.PersonController)
	router.POST("/api/v1/cars/*action", controllers.CarController)
	router.POST("/api/v1/wts/*action", controllers.WorkTypeController)
	router.Static("/admin", "./static/admin")
	router.Run(":3000")

	//	models.UserTableCreate()
	//	models.DealerTableCreate()
	//	models.PersonTableCreate()
	//	models.BrandTableCreate()
	//	models.WorkTypeTableCreate()
	//	models.TimeSlotTableCreate()

	//	r := httprouter.New()

	//	uc := controllers.NewUserController()
	//	dc := controllers.NewDealerController()
	//	pc := controllers.NewPersonController()
	//	bc := controllers.NewBrandController()
	//	wt := controllers.NewWorkTypeController()
	//	tsc := controllers.NewTimeSlotController()

	//	r.GET("/user/list", wlog.HTTPRouteWrapper(uc.Users))
	//	r.POST("/user/register", wlog.HTTPRouteWrapper(uc.RegisterUser))
	//	r.POST("/user/confirm", wlog.HTTPRouteWrapper(uc.ConfirmUser))
	//	/* -- */
	//	r.POST("/api/v1/dealers", wlog.HTTPRouteWrapper(dc.GetDealers))
	//	r.POST("/api/v1/dealers/get", wlog.HTTPRouteWrapper(dc.GetDealer))
	//	r.POST("/api/v1/dealers/add", wlog.HTTPRouteWrapper(dc.AddDealer))
	//	r.POST("/api/v1/dealers/upd", wlog.HTTPRouteWrapper(dc.UpdateDealer))
	//	r.POST("/api/v1/dealers/del", wlog.HTTPRouteWrapper(dc.DeleteDealer))
	//	/* -- */
	//	r.POST("/api/v1/persons", wlog.HTTPRouteWrapper(pc.GetPersons))
	//	r.POST("/api/v1/persons/get", wlog.HTTPRouteWrapper(pc.GetPerson))
	//	r.POST("/api/v1/persons/add", wlog.HTTPRouteWrapper(pc.AddPerson))
	//	r.POST("/api/v1/persons/upd", wlog.HTTPRouteWrapper(pc.UpdatePerson))
	//	r.POST("/api/v1/persons/del", wlog.HTTPRouteWrapper(pc.DeletePerson))
	//	/* -- */
	//	r.POST("/api/v1/brands", wlog.HTTPRouteWrapper(bc.GetBrands))
	//	r.POST("/api/v1/brands/get", wlog.HTTPRouteWrapper(bc.GetBrand))
	//	r.POST("/api/v1/brands/add", wlog.HTTPRouteWrapper(bc.AddBrand))
	//	r.POST("/api/v1/brands/upd", wlog.HTTPRouteWrapper(bc.UpdateBrand))
	//	r.POST("/api/v1/brands/del", wlog.HTTPRouteWrapper(bc.DeleteBrand))
	//	/* -- */
	//	r.POST("/api/v1/wts", wlog.HTTPRouteWrapper(wt.GetWorkTypes))
	//	r.POST("/api/v1/wts/get", wlog.HTTPRouteWrapper(wt.GetWorkType))
	//	r.POST("/api/v1/wts/add", wlog.HTTPRouteWrapper(wt.AddWorkType))
	//	r.POST("/api/v1/wts/upd", wlog.HTTPRouteWrapper(wt.UpdateWorkType))
	//	r.POST("/api/v1/wts/del", wlog.HTTPRouteWrapper(wt.DeleteWorkType))
	//	/* -- */
	//	r.POST("/api/v1/tss", wlog.HTTPRouteWrapper(tsc.GetTimeSlots))
	//	r.POST("/api/v1/tss/get", wlog.HTTPRouteWrapper(tsc.GetTimeSlot))
	//	r.POST("/api/v1/tss/add", wlog.HTTPRouteWrapper(tsc.AddTimeSlot))
	//	r.POST("/api/v1/tss/upd", wlog.HTTPRouteWrapper(tsc.UpdateTimeSlot))
	//	r.POST("/api/v1/tss/del", wlog.HTTPRouteWrapper(tsc.DeleteTimeSlot))

	//	r.ServeFiles("/admin/*filepath", http.Dir("./static/admin/"))

	//	// Fire up the server
	//	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	fmt.Println("PitStop V1")
	handleRequests()
}
