import { ValidationMiddleware } from "../../shared/middlewares/class-validator.middleware";
import { ProcessMatrixStatsDto } from "./dto/process-matrix-stats.dto";
import { statsController } from "./stats.container";
import { Router } from "express";
const router = Router();

router.post(
  "/calculate",
  ValidationMiddleware(ProcessMatrixStatsDto),
  statsController.handleCalculate.bind(statsController),
);

export default router;
