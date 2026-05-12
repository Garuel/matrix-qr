export interface MatrixRequest {
  matrix: number[][];
}

export interface MatrixResponse {
  data: {
    average: number;
    isDiagonal: boolean;
    max: number;
    min: number;
    sum: number;
  };
}
