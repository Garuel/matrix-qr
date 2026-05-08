package matrix

import (
	"errors"
	"log"
	"matriz-procesador-go-api/internal/domain/models"

	"gonum.org/v1/gonum/mat"
)


type Service interface {
	FactorizeQR(input [][]float64) (models.MatrixResponse, error)
}


type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) FactorizeQR(input [][]float64) (models.MatrixResponse, error) {

    if len(input) == 0 || len(input[0]) == 0 {
		return models.MatrixResponse{}, errors.New("input matrix cannot be empty")
	}

    log.Println("Validando la matriz...")

	
	rows := len(input)
	cols := len(input[0])


    for _, row := range input {
		if len(row) != cols {
			return models.MatrixResponse{}, errors.New("all rows must have the same length")
		}
	}

    log.Println("Matriz validada exitosamente!")
    
    data := make([]float64, 0, rows*cols)

    for _, row := range input {
        data = append(data, row...)
    }

	log.Printf("matriz convertida a dense")
    
    dense := mat.NewDense(rows, cols, data)
    var qr mat.QR
	log.Printf("Factorizando matriz...")
    qr.Factorize(dense)


    var q, r mat.Dense

    qr.QTo(&q)
    qr.RTo(&r)

	log.Printf("matriz factorizada exitosamente!")

    return models.MatrixResponse{
        MatrixQ: matrixToSlice(&q),
        MatrixR: matrixToSlice(&r),
    }, nil
}

func matrixToSlice(m mat.Matrix) [][]float64 {
    r, c := m.Dims()
    matrix := make([][]float64, r)
    for i := 0; i < r; i++ {
        matrix[i] = make([]float64, c)
        for j := 0; j < c; j++ {
            matrix[i][j] = m.At(i, j)
        }
    }
    return matrix
}