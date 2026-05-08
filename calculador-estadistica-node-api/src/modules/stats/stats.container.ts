import dotenv from "dotenv";
import { StatsService } from "./stats.service";
import { StatsController } from "./stats.controller";

dotenv.config();

const statsService = new StatsService();

// para inyectar el servicio en el controlador
const statsController = new StatsController(statsService);

export { statsController };
