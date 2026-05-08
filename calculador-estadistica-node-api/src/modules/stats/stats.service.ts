import { logger } from "../../shared/infrastructure/utils/logger.util";
import { ProcessMatrixStatsDto } from "./dto/process-matrix-stats.dto";
import { ProcessMatrixStatsResponse } from "./interface/process-matrix-stats.interface";

export class StatsService {
  public calculateStats(
    data: ProcessMatrixStatsDto,
  ): ProcessMatrixStatsResponse {
    logger.verbose("calculateStats");
    const { matrixQ, matrixR } = data;
    const allValues = [...matrixQ.flat(), ...matrixR.flat()];
    const sum = allValues.reduce((acc, val) => acc + val, 0);
    const avg = sum / allValues.length;
    const min = Math.min(...allValues);
    const max = Math.max(...allValues);

    return {
      max,
      min,
      average: avg,
      sum,
      isDiagonal:
        this.checkIfDiagonal(matrixQ) || this.checkIfDiagonal(matrixR),
    };
  }

  private checkIfDiagonal(matrix: number[][]): boolean {
    logger.info(
      "Revisando si matriz es diagonal: para todo i != j, el valor es 0",
    );
    for (let i = 0; i < matrix.length; i++) {
      for (let j = 0; j < matrix[i].length; j++) {
        if (i !== j && Math.abs(matrix[i][j]) > 1e-10) {
          return false;
        }
      }
    }
    return true;
  }
}
