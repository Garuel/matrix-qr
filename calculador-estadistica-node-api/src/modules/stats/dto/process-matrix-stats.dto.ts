import { ArrayMinSize, IsArray, IsNumber } from "class-validator";

export class ProcessMatrixStatsDto {
  @IsArray()
  @ArrayMinSize(1)
  matrixQ!: number[][];

  @IsArray()
  @ArrayMinSize(1)
  // @Validate(IsNumberMatrix, {
  //     message: "matrixR debe ser una matriz de números (number[][])",
  //   })
  matrixR!: number[][];
}
