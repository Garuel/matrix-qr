package matrix

import (
	"testing"
)

func TestFactorizeQR(t *testing.T) {
    s := NewService()

	//Matriz valida
    input := [][]float64{
        {1, 0},
        {0, 1},
    }

    res, err := s.FactorizeQR(input)
    if err != nil {
        t.Errorf("No se esperaba error, pero se obtuvo: %v", err)
    }

    if len(res.MatrixQ) != 2 || len(res.MatrixR) != 2 {
        t.Errorf("El tamaño de las matrices resultantes es incorrecto")
    }

    // Matriz invalida
    invalidInput := [][]float64{
        {1, 2, 3},
        {4, 5}, 
    }

    _, err = s.FactorizeQR(invalidInput)
    if err == nil {
        t.Error("Se esperaba un error por filas de distinta longitud, pero no se obtuvo")
    }
}