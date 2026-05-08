package models

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixResponse struct {
	MatrixQ [][]float64 `json:"matrixQ"`
	MatrixR [][]float64 `json:"matrixR"`
}