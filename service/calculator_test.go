package service_test

import (
	"hello-api/service"
	"testing"

	"github.com/stretchr/testify/mock"
)

type stubbedService struct{}

func (s *stubbedService) Sum(x, y int) int {
	return x + y
}

// สร้าง Mock Object
type MockService struct {
	mock.Mock
}

func (m *MockService) Sum(x, y int) int {
	args := m.Called(x, y)
	return args.Int(0)
}

// func TestCalculatorService(t *testing.T) {

// 	data := []struct {
// 		name     string
// 		x, y     int
// 		expected int
// 	}{
// 		{"positive numbers", 2, 3, 5},
// 		{"negative numbers", -1, -5, -6},
// 		{"mixed numbers", -1, 10, 9},
// 		{"zero values", 0, 0, 0},
// 	}

// 	// 2. Setup Handler โดยส่ง Stub เข้าไป (Dependency Injection)
// 	// ตรงนี้แหละครับที่เรียกว่า "Mock ได้" เพราะ NewCalculatorHandler รับ Interface
// 	// stub := &stubbedService{}
// 	// h := handler.NewCalculatorHandler(stub)

// 	for _, tt := range data {
// 		t.Run(tt.name, func(t *testing.T) {
// 			handler := service.NewCalculatorService()
// 			// 3. สร้าง HTTP Request จำลอง
// 			url := fmt.Sprintf("/sum?x=%d&y=%d", tt.x, tt.y)
// 			req, err := http.NewRequest("GET", url, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			// 4. สร้างตัวบันทึก Response (Recorder)
// 			rr := httptest.NewRecorder()

// 			// 5. รัน Handler
// 			handlerFunc := http.HandlerFunc(handler.Sum())
// 			handlerFunc.ServeHTTP(rr, req)

// 			// 6. Assert Status Code
// 			if rr.Code != http.StatusOK {
// 				t.Errorf("expected status 200, got %d", rr.Code)
// 			}

// 			// 7. Assert Body Result
// 			if rr.Body.String() != fmt.Sprint(tt.expected) {
// 				t.Errorf("handler result wrong: got %s, want %d", rr.Body.String(), tt.expected)
// 			}
// 		})
// 	}

// }

func TestCalculatorServiceSum(t *testing.T) {
	s := service.NewCalculatorService()

	data := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -5, -6},
		{"mixed numbers", -1, 10, 9},
		{"zero values", 0, 0, 0},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := s.Sum(d.x, d.y)
			if result != d.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", d.x, d.y, result, d.expected)
			}
		})
	}
}
