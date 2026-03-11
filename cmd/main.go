package main

import (
	"hello-api/handler"
	"hello-api/rest"
	"hello-api/service"
	"hello-api/translation"
	"log"
	"net/http"
)

func main() {
	// 1. สร้าง Concrete Service (ตัวทำงานจริง)
	// ในอนาคตถ้าต้องการเปลี่ยนจาก Static เป็น Database หรือ External API
	// ก็แค่เปลี่ยนบรรทัดนี้โดยที่โค้ดส่วนอื่นไม่ต้องแก้ไขเลย
	translationService := translation.NewStaticService()

	calculatorService := service.NewCalculatorService()
	calculatorHandler := handler.NewCalculatorHandler(*calculatorService)

	// 2. ฉีด Service เข้าไปใน Handler (Dependency Injection)
	translateHandler := rest.NewTranslateHandler(translationService)

	// 3. กำหนด Routing
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/calculator", calculatorHandler.Sum)

	// 4. เริ่มการทำงานของ Server
	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
