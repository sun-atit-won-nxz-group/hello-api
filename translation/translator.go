package translation

// Translator คือ Interface สำหรับบริการแปลภาษา (Dependency Inversion)
type Translator interface {
	Translate(word string, language string) string
}

// StaticService คือ Service จริงที่ทำงานตาม Interface
type StaticService struct{}

func NewStaticService() *StaticService {
	return &StaticService{}
}

func (s *StaticService) Translate(word string, language string) string {
	if language == "german" {
		return "hallo"
	}
	return "hello"
}
