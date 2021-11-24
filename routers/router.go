package routers

import (
	"fmt"
	"io"
	"os"
	"ottosfa-api-apk/docs"
	v22Controller "ottosfa-api-apk/v2.2/controllers"
	v22merchantsController "ottosfa-api-apk/v2.2/controllers/merchants"
	v22salesController "ottosfa-api-apk/v2.2/controllers/sales"
	v22sfaController "ottosfa-api-apk/v2.2/controllers/sfa"
	v22tokoOttopayController "ottosfa-api-apk/v2.2/controllers/tokoOttopay"
	v22updateQrisController "ottosfa-api-apk/v2.2/controllers/updateQRIS"
	v23Controller "ottosfa-api-apk/v2.3/controllers"
	v23merchantsController "ottosfa-api-apk/v2.3/controllers/merchants"
	v23sfaController "ottosfa-api-apk/v2.3/controllers/sfa"
	v23sfaIdmController "ottosfa-api-apk/v2.3/controllers/sfa-idm"
	v23tokoOttopayController "ottosfa-api-apk/v2.3/controllers/tokoOttopay"
	v24sfaController "ottosfa-api-apk/v2.4/controllers/sfa"
	v25indomarcoController "ottosfa-api-apk/v2.5/controllers/indomarco"
	v25sfaController "ottosfa-api-apk/v2.5/controllers/sfa"
	v26indomarcoController "ottosfa-api-apk/v2.6/controllers/indomarco"
	v26sfaController "ottosfa-api-apk/v2.6/controllers/sfa"
	ver "ottosfa-api-apk/version"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/opentracing/opentracing-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"ottodigital.id/library/httpserver/ginserver"
	ottologer "ottodigital.id/library/logger"
	"ottodigital.id/library/ottotracing"
	"ottodigital.id/library/utils"
)

// ServerEnv ..
type ServerEnv struct {
	ServiceName     string `envconfig:"SERVICE_NAME" default:"OTTO-SFA-API"`
	OpenTracingHost string `envconfig:"OPEN_TRACING_HOST" default:"13.250.21.165:5775"`
	DebugMode       string `envconfig:"DEBUG_MODE" default:"debug"`
	ReadTimeout     int    `envconfig:"READ_TIMEOUT" default:"120"`
	WriteTimeout    int    `envconfig:"WRITE_TIMEOUT" default:"120"`
}

var (
	server ServerEnv
)

//type OttoSvrGin interface {
//	ServerUp (listenAddr string, router *gin.Engine) error
//}

func init() {
	if err := envconfig.Process("SERVER", &server); err != nil {
		fmt.Println("Failed to get SERVER env:", err)
	}

	//header = "/v2"

	//accesspointcheckmerchantottopay = header + "/merchants/check_merchant_ottopay"
	//accesspointcheckmerchantbyid = header + "/merchants/check_merchant_id"
	//accesspointreporthistorydetail = header + "/merchants/history_detail"
	//accesspointreporthistorysummary = header + "/merchants/history_summary"
	//accesspointchangemerchantphone = header + "/merchants/change_merchant_phone"
	//accesspointreportbysales = header + "/merchants/report_by_sales"
	//accesspointupdatephotoprofilesales = header + "/sales/profile/update_photo"
	//accesspointchangepasswordsales = header + "/sfa/auth/change_password"
	//accesspointcheckprofilsales = header + "/sfa/auth/profile"
	//accesspointcheckversionapps = header + "/sfa/check_version"
	//accesspointlogout = header + "/sfa/auth/logout"
	//accesspointchecktoken = header + "/sfa/auth/check_token"
	//accesspointupdatemerchantindomarco = header + "/merchants/update_merchant_indomarco"
	//accesspointcheckaccountindomarco = header + "/sfa/check_indomarco_account"
	//accesspointchangephotopinktp = header + "/sfa/auth/change_pin_and_ktp"

	//debugmode = beego.AppConfig.DefaultString("apps.debug", "release")
	//
	//err := envconfig.Process("OTTOSFA_API_APK", &server)
	//if err != nil {
	//	fmt.Println("Failed to get OTTOSFA_API_APK env:", err)
	//}
	//
	//nameService = utils.GetEnv("OTTOSFA_API_APK", "OTTOSFA_API_APK")
	//openTracingSvr = utils.GetEnv("JAEGER_HOSTURL", "13.250.21.165:5775")
}

// Server ..
func Server(listenAddress string) error {
	sugarLogger := ottologer.GetLogger()

	ottoRouter := OttoRouter{}
	ottoRouter.InitTracing()
	ottoRouter.Routers()
	defer ottoRouter.Close()

	err := ginserver.GinServerUp(listenAddress, ottoRouter.Router)

	if err != nil {
		fmt.Println("Error:", err)
		sugarLogger.Error("Error ", zap.Error(err))
		return err
	}

	fmt.Println("Server UP")
	sugarLogger.Info("Server UP ", zap.String("listenAddress", listenAddress))

	return err
}

// OttoRouter ..
type OttoRouter struct {
	Tracer   opentracing.Tracer
	Reporter jaeger.Reporter
	Closer   io.Closer
	Err      error
	GinFunc  gin.HandlerFunc
	Router   *gin.Engine
}

// Routers ..
func (ottoRouter *OttoRouter) Routers() {
	gin.SetMode(server.DebugMode)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "ottosfa-api-apk API"
	docs.SwaggerInfo.Description = "<ottosfa-api-apk description>"
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.BasePath = "/ottosfa-api-apk"
	docs.SwaggerInfo.Schemes = []string{"http"}
	switch utils.GetEnv("APPS_ENV", "local") {
	case "local":
		docs.SwaggerInfo.Host = utils.GetEnv("SWAGGER_HOST_LOCAL", "localhost:8044")
	case "dev":
		docs.SwaggerInfo.Host = utils.GetEnv("SWAGGER_HOST_DEV", "")
	}

	router := gin.New()
	router.Use(ottoRouter.GinFunc)
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	v22 := router.Group("/v2.2")
	{
		merchants := v22.Group("/merchants")
		{
			merchants.POST("/check_qris", v22merchantsController.CheckQRIS)
			merchants.POST("/reverse_qr", v22merchantsController.ReverseQR)
			merchants.POST("/check_idcard", v22merchantsController.CheckIdCard)
			merchants.POST("/update", v22merchantsController.UpdateMerchant)

			merchants.GET("/business_type/list", v22merchantsController.BusinessTypeList)
			merchants.GET("/business_category/list", v22merchantsController.BusinessTypeList)
			merchants.GET("/acquitision/dropdown/list", v22merchantsController.AcquitisionDropdownList)

			merchants.POST("/indomarco/check_merchant_ottopay", v22merchantsController.IndomarcoCheckMerchantOttopay)

			merchants.POST("/report_by_sales", v22merchantsController.ReportBySales)
			merchants.POST("/history_summary", v22merchantsController.HistorySummary)
			merchants.POST("/history_detail", v22merchantsController.HistoryDetail)
		}
		sales := v22.Group("/sales")
		{
			sales.GET("/profile", v22salesController.Profile)
		}
		sfa := v22.Group("/sfa")
		{
			sfa.POST("/merchant/list", v22sfaController.MerchantList)

			callPlan := sfa.Group("/call_plan")
			{
				callPlan.POST("/list", v22sfaController.CallPlanList)
				callPlan.POST("/merchant/list", v22sfaController.CallPlanMerchantList)
				callPlan.POST("/todolist/list", v22sfaController.CallPlanTodolistList)
				callPlan.POST("/todolist/merchant_not_found", v22sfaController.CallPlanTodolistMerchantNotFound)
				callPlan.GET("/action_merchant/list", v22sfaController.CallPlanActionMerchantList)
				callPlan.GET("/product_merchant/list", v22sfaController.CallPlanProductMerchantList)
				callPlan.POST("/description_merchant", v22sfaController.CallPlanDescriptionMerchant)
				callPlan.POST("/visit/check_merchant_phone", v22sfaController.CallPlanVisitCheckMerchantPhone)
				callPlan.POST("/visit/check_qris", v22sfaController.CallPlanVisitCheckQRIS)
				callPlan.POST("/visit/add", v22sfaController.CallPlanVisitAdd)
				callPlan.POST("/detail", v22sfaController.CallPlanDetail)
				callPlan.POST("/action/check_merchant_phone", v22sfaController.CallPlanActionCheckMerchantPhone)
				callPlan.POST("/action/check_qris", v22sfaController.CallPlanActionCheckQRIS)
				callPlan.POST("/action/update_clock_in", v22sfaController.CallPlanActionUpdateClockInMerchant)
				callPlan.POST("/action/merchant_unknown", v22sfaController.CallPlanActionMerchantUnknown)
				callPlan.POST("/action/detail", v22sfaController.CallPlanActionDetail)
				callPlan.POST("/action/save", v22sfaController.CallPlanActionAddOrEdit)
				callPlan.POST("/action/submit", v22sfaController.CallPlanActionSubmit)
				callPlan.POST("/action/delete", v22sfaController.CallPlanActionDelete)
			}

			auth := sfa.Group("/auth")
			{
				auth.POST("/login", v22sfaController.Login)
				auth.POST("/update_pin_idcard_photo", v22sfaController.ChangePinAndIdCardPhoto)
			}

			todolist := sfa.Group("/todolist")
			{
				todolist.POST("/post", v22sfaController.TodolistPost)
				todolist.POST("/filter/village", v22sfaController.TodolistFilterVillageList)
				todolist.GET("/count", v22sfaController.TodolistCount)

				//todolist rose
				todolist.POST("/taskbysubcategory", v22sfaController.TodolistTaskBySubCategory)
			}

			sfa.POST("/attendance/post", v22sfaController.AttendancePost)
			//sfa.POST("/indomarco/login", v22sfaController.LoginIndomarco)
			sfa.POST("/sales/upload/photo", v22sfaController.UploadPhotoProfilSales)

			sfa.GET("/user_category", v22sfaController.UserCategory)
			sfa.GET("/clear_session", v22sfaController.ClearSession)
			sfa.POST("/push_notif", v22sfaController.PushNotif)
			sfa.POST("/merchant_not_found/list", v22sfaController.MerchantNotFoundList)

			sfa.POST("/check_indomarco_account", v22sfaController.CheckIndomarcoAccount)
		}

		qris := v22.Group("/qris")
		{
			qris.POST("/merchant/detail", v22updateQrisController.MerchantDetailQRIS)
			qris.POST("/scan", v22updateQrisController.ScanAndUpdateQRIS)
		}

		tokoOttopay := v22.Group("/toko-ottopay")
		{
			tokoOttopay.POST("/merchant/detail", v22tokoOttopayController.MerchantDetailTokoOttopay)
		}

		v22.GET("/province/:country_id", v22Controller.Province)
		v22.GET("/city/:province_id", v22Controller.City)
		v22.GET("/district/:city_id", v22Controller.District)
		v22.GET("/village/:district_id", v22Controller.Village)

		v22.POST("/ottosfa/merchants/check_idcardchange_status", v22Controller.ChangeStatusMerchant)
		v22.POST("/acquitisions/merchant/detail", v22Controller.AcquisitionsMerchantDetail)

		v22.POST("/ottosfa/health_check", v22Controller.HealthCheck)
	}

	v23 := router.Group("/v2.3")
	{
		sfa := v23.Group("/sfa")
		{
			callPlan := sfa.Group("/call_plan")
			{
				callPlan.POST("/list", v23sfaController.CallPlanList)
				//callPlan.POST("/action/detail", v23sfaController.CallPlanActionDetail)
				callPlan.POST("/action/scan_qr", v23sfaController.CallPlanActionScanQR)
				callPlan.POST("/action/save", v23sfaController.CallPlanActionAddOrEdit)
				callPlan.POST("/action/merchant_unknown", v23sfaController.CallPlanActionMerchantUnknown)
				callPlan.POST("/action/submit", v23sfaController.CallPlanActionSubmit)
				callPlan.POST("/merchant/list", v23sfaController.CallPlanMerchantList)

				//callplan rose
				callPlan.POST("/visit/check_merchant_phone", v23sfaController.CallPlanVisitCheckMerchantPhone)
				callPlan.POST("/visit/check_qris", v23sfaController.CallPlanVisitCheckQRIS)
				callPlan.POST("/todolist/list", v23sfaController.CallPlanTodolistList)
			}

			todolist := sfa.Group("/todolist")
			{
				//todolist.POST("/post", v22sfaController.TodolistPost)
				todolist.POST("/list", v23sfaController.TodolistList)
				todolist.POST("/detail", v23sfaController.TodolistDetail)
				todolist.POST("/post", v23sfaController.TodolistPost)
				todolist.POST("/merchant_not_found/list", v23sfaController.TodolistMerchantNotFoundList)
			}

			auth := sfa.Group("/auth")
			{
				auth.POST("/login", v23sfaController.Login)
			}

			sfa.POST("/attendance/post", v23sfaController.AttendancePost)
			sfa.POST("/merchant/list", v23sfaController.MerchantList)
		}

		sfaIdm := v23.Group("/sfa-idm")
		{
			todolist := sfaIdm.Group("/todolist")
			{
				todolist.POST("/list", v23sfaIdmController.IDMTodolistList)
				todolist.POST("/detail", v23sfaIdmController.IDMTodolistDetail)
			}
		}

		merchants := v23.Group("/merchants")
		{
			merchants.GET("/business_type/list", v23Controller.BusinessTypeList)
			merchants.GET("/business_category/list", v23Controller.BusinessCategoryList)
			merchants.POST("/business_type/list", v23Controller.BusinessTypeListBySR)
			merchants.POST("/check_idcard", v23merchantsController.CheckIdCard)
		}

		tokoOttopay := v23.Group("/toko-ottopay")
		{
			tokoOttopay.POST("/merchant/detail", v23tokoOttopayController.MerchantDetailTokoOttopay)
		}

	}

	v24 := router.Group("/v2.4")
	{
		sfa := v24.Group("/sfa")
		{
			todolist := sfa.Group("/todolist")
			{
				todolist.POST("/merchant_not_found/list", v24sfaController.TodolistMerchantNotFoundList)
				todolist.POST("/list", v24sfaController.TodolistList)

				//todolist rose
				todolist.POST("/detail", v24sfaController.TodolistDetail)
				todolist.POST("/post", v24sfaController.TodolistPost)
			}

			callPlan := sfa.Group("/call_plan")
			{
				callPlan.POST("/merchant/list", v24sfaController.CallPlanMerchantList)

				//callplan rose
				callPlan.POST("/action/save", v24sfaController.CallPlanActionAddOrEdit)
			}

			auth := sfa.Group("/auth")
			{
				auth.POST("/login", v24sfaController.Login)
			}
		}
	}

	v25 := router.Group("/v2.5")
	{
		sfa := v25.Group("/sfa")
		{
			businessTypeList := sfa.Group("/business_type")
			{
				businessTypeList.POST("/list", v25sfaController.BusinessTypeList)
			}

			aquisition := sfa.Group("/acquisition")
			{
				aquisition.POST("/list", v25sfaController.AcquisitionList)
			}

			callPlan := sfa.Group("/call_plan")
			{
				//callplan rose
				callPlan.POST("/merchant/list", v25sfaController.CallPlanMerchantList)
			}

			todolist := sfa.Group("/todolist")
			{
				//todolist rose
				todolist.POST("/list", v25sfaController.TodolistList)
			}
		}

		indomarco := v25.Group("/sfa-idm")
		{
			updateMerchant := indomarco.Group("/merchant")
			{
				updateMerchant.POST("/update", v25indomarcoController.UpdateMerchantIndomarco)
			}
		}
	}

	v26 := router.Group("/v2.6")
	{
		indomarco := v26.Group("/sfa-idm")
		{
			updateMerchant := indomarco.Group("/merchant")
			{
				updateMerchant.POST("/update", v26indomarcoController.UpdateMerchantIndomarco)
			}
		}

		sfa := v26.Group("/sfa")
		{
			aquisition := sfa.Group("/acquisition")
			{
				aquisition.POST("/list", v26sfaController.AcquisitionList)
			}
		}

	}

	router.GET("/version", ver.Version)

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ottoRouter.Router = router

}

// InitTracing ..
func (ottoRouter *OttoRouter) InitTracing() {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "PROD"
	}

	tracer, reporter, closer, err := ottotracing.InitTracing(fmt.Sprintf("%s::%s", server.ServiceName, hostName), server.OpenTracingHost, ottotracing.WithEnableInfoLog(true))
	if err != nil {
		fmt.Println("Error :", err)
	}
	opentracing.SetGlobalTracer(tracer)

	ottoRouter.Closer = closer
	ottoRouter.Reporter = reporter
	ottoRouter.Tracer = tracer
	ottoRouter.Err = err
	ottoRouter.GinFunc = ottotracing.OpenTracer([]byte("api-request-"))
}

// Close ..
func (ottoRouter *OttoRouter) Close() {
	ottoRouter.Closer.Close()
	ottoRouter.Reporter.Close()
}
