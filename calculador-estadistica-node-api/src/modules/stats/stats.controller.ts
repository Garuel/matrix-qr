import { NextFunction, Request, Response } from "express";
import { StatsService } from "./stats.service";
import { logger } from "../../shared/infrastructure/utils/logger.util";
import { HttpStatusCode } from "../../shared/domain/enum/custom-error.enum";
import { TipoRespuestaEnum } from "../../shared/domain/enum/tipo-alerta.enum";

export class StatsController {
  constructor(private statsService: StatsService) {}

  public handleCalculate(req: Request, res: Response, next: NextFunction) {
    try {
      logger.info("Recibiendo matrices para cálculo de estadísticas...");
      const result = this.statsService.calculateStats(req.body);
      logger.info("Estadísticas calculadas con éxito.");
      res.status(HttpStatusCode.OK).json({
        message: "Estadística calculada",
        tipoRespuesta: TipoRespuestaEnum.Success,
        title: "HandleCalculate",
        data: result,
      });
    } catch (error) {
      next(error);
    }
  }
}
