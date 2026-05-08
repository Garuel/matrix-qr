import "reflect-metadata";

import app from "./server";
import { logger } from "./shared/infrastructure/utils/logger.util";

const PORT = process.env.PORT || 6687;

async function startServer() {
  try {
    app.listen(PORT, () => {
      logger.info(`Servidor corriendo en http://localhost:${PORT}`);
    });
  } catch (error) {
    logger.error("Error fatal al iniciar el servidor:", error);
    process.exit(1);
  }
}

startServer();
