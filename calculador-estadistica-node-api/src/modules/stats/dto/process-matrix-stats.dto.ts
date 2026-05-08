import { ArrayMinSize, IsArray, Validate } from "class-validator";
import { IsNumberMatrix } from "../../../shared/infrastructure/validators/matrix.validator";

export class ProcessMatrixStatsDto {
  @IsArray()
  @ArrayMinSize(1)
  @Validate(IsNumberMatrix, {
    message: "matrixQ debe ser una matriz de números (number[][])",
  })
  matrixQ!: number[][];

  @IsArray()
  @ArrayMinSize(1)
  @Validate(IsNumberMatrix, {
    message: "matrixR debe ser una matriz de números (number[][])",
  })
  matrixR!: number[][];
}
