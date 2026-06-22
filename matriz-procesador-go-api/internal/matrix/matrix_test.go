package matrix

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// matrizMultiplicar multiplica dos matrices para verificar que Q*R == original.
func matrizMultiplicar(a, b [][]float64) [][]float64 {
	filasA := len(a)
	columnasB := len(b[0])
	filasBColumnasA := len(b)
	result := make([][]float64, filasA)
	for i := 0; i < filasA; i++ {
		result[i] = make([]float64, columnasB)
		for j := 0; j < columnasB; j++ {
			for p := 0; p < filasBColumnasA; p++ {
				result[i][j] += a[i][p] * b[p][j]
			}
		}
	}
	return result
}

func TestFactorizeQR(t *testing.T) {
	t.Parallel()

	t.Run("matriz identidad 2x2", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{1, 0},
			{0, 1},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err, "no se esperaba error con matriz identidad")

		assert.Len(t, res.MatrixQ, 2, "Q debe tener 2 filas")
		assert.Len(t, res.MatrixR, 2, "R debe tener 2 filas")
		assert.Len(t, res.MatrixQ[0], 2, "Q debe tener 2 columnas")
		assert.Len(t, res.MatrixR[0], 2, "R debe tener 2 columnas")
	})

	t.Run("matriz 3x3 produce factorización válida", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{12, -51, 4},
			{6, 167, -68},
			{-4, 24, -41},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		assert.Len(t, res.MatrixQ, 3, "Q debe tener 3 filas")
		assert.Len(t, res.MatrixR, 3, "R debe tener 3 filas")

		producto := matrizMultiplicar(res.MatrixQ, res.MatrixR)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				assert.True(t, almostEqual(producto[i][j], input[i][j], 1e-9),
					"Q*R[%d][%d] = %f, esperado %f", i, j, producto[i][j], input[i][j])
			}
		}
	})

	t.Run("R es triangular superior", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 10},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		for i := 0; i < len(res.MatrixR); i++ {
			for j := 0; j < i; j++ {
				assert.True(t, almostEqual(res.MatrixR[i][j], 0, 1e-9),
					"R[%d][%d] = %f, debería ser ~0 (triangular superior)", i, j, res.MatrixR[i][j])
			}
		}
	})

	t.Run("Q es ortogonal (Q^T * Q ≈ I)", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{2, -1, 0},
			{-1, 2, -1},
			{0, -1, 2},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		n := len(res.MatrixQ)

		// Trasponer Q
		qt := make([][]float64, n)
		for i := 0; i < n; i++ {
			qt[i] = make([]float64, n)
			for j := 0; j < n; j++ {
				qt[i][j] = res.MatrixQ[j][i]
			}
		}

		// Q^T * Q debe ser ≈ I
		resultado := matrizMultiplicar(qt, res.MatrixQ)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				expected := 0.0
				if i == j {
					expected = 1.0
				}
				assert.True(t, almostEqual(resultado[i][j], expected, 1e-9),
					"(Q^T * Q)[%d][%d] = %f, esperado %f", i, j, resultado[i][j], expected)
			}
		}
	})
}

func TestFactorizeQR_Validaciones(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nombre        string
		input         [][]float64
		errorEsperado string
	}{
		{
			nombre:        "matriz vacía",
			input:         [][]float64{},
			errorEsperado: "La matriz de entrada no puede estar vacía",
		},
		{
			nombre:        "fila vacía",
			input:         [][]float64{{}},
			errorEsperado: "La matriz de entrada no puede estar vacía",
		},
		{
			nombre: "filas de distinta longitud",
			input: [][]float64{
				{1, 2, 3},
				{4, 5},
			},
			errorEsperado: "Todas las filas tienen que tener la misma longitud",
		},
		{
			nombre: "matriz no cuadrada (más columnas que filas)",
			input: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
			},
			errorEsperado: "La matriz debe ser cuadrada",
		},
		{
			nombre: "matriz no cuadrada (más filas que columnas)",
			input: [][]float64{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			errorEsperado: "La matriz debe ser cuadrada",
		},
	}

	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			t.Parallel()
			s := NewService()

			res, err := s.FactorizeQR(tc.input)

			require.Error(t, err, "se esperaba error para: %s", tc.nombre)
			assert.Equal(t, tc.errorEsperado, err.Error())
			assert.Empty(t, res.MatrixQ, "MatrixQ debe estar vacía en caso de error")
			assert.Empty(t, res.MatrixR, "MatrixR debe estar vacía en caso de error")
		})
	}
}

func TestFactorizeQR_CasosEspeciales(t *testing.T) {
	t.Parallel()

	t.Run("matriz 1x1", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{{5}}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		assert.Len(t, res.MatrixQ, 1)
		assert.Len(t, res.MatrixR, 1)
		assert.Len(t, res.MatrixQ[0], 1)
		assert.Len(t, res.MatrixR[0], 1)

		// Q * R debe dar la original
		producto := res.MatrixQ[0][0] * res.MatrixR[0][0]
		assert.True(t, almostEqual(producto, 5, 1e-9),
			"Q[0][0] * R[0][0] = %f, esperado 5", producto)
	})

	t.Run("matriz con valores negativos", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{-3, -1},
			{-2, -4},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		producto := matrizMultiplicar(res.MatrixQ, res.MatrixR)
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				assert.True(t, almostEqual(producto[i][j], input[i][j], 1e-9),
					"Q*R[%d][%d] = %f, esperado %f", i, j, producto[i][j], input[i][j])
			}
		}
	})

	t.Run("matriz con valores decimales", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{1.5, 2.7},
			{3.2, 0.8},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		producto := matrizMultiplicar(res.MatrixQ, res.MatrixR)
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				assert.True(t, almostEqual(producto[i][j], input[i][j], 1e-9),
					"Q*R[%d][%d] = %f, esperado %f", i, j, producto[i][j], input[i][j])
			}
		}
	})

	t.Run("matriz con ceros", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{0, 1},
			{1, 0},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		producto := matrizMultiplicar(res.MatrixQ, res.MatrixR)
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				assert.True(t, almostEqual(producto[i][j], input[i][j], 1e-9),
					"Q*R[%d][%d] = %f, esperado %f", i, j, producto[i][j], input[i][j])
			}
		}
	})
}

func TestMatrixToSlice(t *testing.T) {
	t.Parallel()

	t.Run("preserva dimensiones correctamente", func(t *testing.T) {
		t.Parallel()
		s := NewService()

		input := [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		res, err := s.FactorizeQR(input)
		require.NoError(t, err)

		// Q y R deben tener las mismas dimensiones que la entrada
		assert.Len(t, res.MatrixQ, 3, "Q debe tener 3 filas")
		assert.Len(t, res.MatrixR, 3, "R debe tener 3 filas")
		for i := 0; i < 3; i++ {
			assert.Len(t, res.MatrixQ[i], 3, "Q fila %d debe tener 3 columnas", i)
			assert.Len(t, res.MatrixR[i], 3, "R fila %d debe tener 3 columnas", i)
		}
	})
}
