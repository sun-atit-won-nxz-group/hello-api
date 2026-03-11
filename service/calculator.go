package service

type Calculator interface {
	Sum(x, y int) int
}

type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) Sum(x, y int) int {
	return x + y
}
