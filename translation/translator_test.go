package translation_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// สร้าง Mock สำหรับ Translator
type MockTranslator struct {
	mock.Mock
}

func (m *MockTranslator) Translate(word, language string) string {
	args := m.Called(word, language)
	return args.String(0)
}

// สร้าง Test Suite
type TranslatorTestSuite struct {
	suite.Suite
	mockService *MockTranslator
}

func (suite *TranslatorTestSuite) SetupTest() {
	suite.mockService = new(MockTranslator)
}

func (suite *TranslatorTestSuite) TestTranslateWithMock() {
	// stub
	// Arrange: กำหนดว่าถ้าเรียก Translate ด้วย "hello", "thai" ให้คืนค่า "สวัสดี"
	suite.mockService.On("Translate", "hello", "thai").Return("สวัสดี")

	// Act
	res := suite.mockService.Translate("hello", "thai")

	// Assert
	suite.Equal("สวัสดี", res)
	suite.mockService.AssertExpectations(suite.T()) // ตรวจสอบว่าถูกเรียกจริงตามเงื่อนไข
}

func TestTranslatorSuite(t *testing.T) {
	suite.Run(t, new(TranslatorTestSuite))
}
