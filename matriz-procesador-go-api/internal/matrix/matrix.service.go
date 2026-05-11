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
		return models.MatrixResponse{}, errors.New("La matriz de entrada no puede estar vacía")
	}

    log.Println("Validando la matriz...")

	
	rows := len(input)
	cols := len(input[0])


    for _, row := range input {
		if len(row) != cols {
			return models.MatrixResponse{}, errors.New("Todas las filas tienen que tener la misma longitud")
		}
	}

    if rows != cols {
        return models.MatrixResponse{}, errors.New("La matriz debe ser cuadrada")
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
    rows, cols := m.Dims()
    matrix := make([][]float64, rows)
    for i := 0; i < rows; i++ {
        matrix[i] = make([]float64, cols)
        for j := 0; j < cols; j++ {
            matrix[i][j] = m.At(i, j)
        }
    }
    return matrix
}