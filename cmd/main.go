package main

import (
	"hello-api/rest"
	"hello-api/translation"
	"log"
	"net/http"
)

func main() {
	// 1. สร้าง Concrete Service (ตัวทำงานจริง)
	// ในอนาคตถ้าต้องการเปลี่ยนจาก Static เป็น Database หรือ External API
	// ก็แค่เปลี่ยนบรรทัดนี้โดยที่โค้ดส่วนอื่นไม่ต้องแก้ไขเลย
	translationService := translation.NewStaticService()

	// 2. ฉีด Service เข้าไปใน Handler (Dependency Injection)
	translateHandler := rest.NewTranslateHandler(translationService)

	// 3. กำหนด Routing
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)

	// 4. เริ่มการทำงานของ Server
	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
