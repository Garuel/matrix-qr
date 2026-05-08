import {
  ValidatorConstraint,
  ValidatorConstraintInterface,
} from "class-validator";

@ValidatorConstraint({ name: "isNumberMatrix", async: false })
export class IsNumberMatrix implements ValidatorConstraintInterface {
  validate(matrix: any[][]) {
    if (!Array.isArray(matrix)) return false;
    return matrix.every(
      (row) =>
        Array.isArray(row) &&
        row.length > 0 &&
        row.every((item) => typeof item === "number"),
    );
  }
}
