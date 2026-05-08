import express, { Application } from "express";
import cors from "cors";
import { globalErrorHandler } from "./shared/middlewares/error.middleware";
import dotenv from "dotenv";
import statsRouter from "./modules/stats/stats.route";

dotenv.config();

const app: Application = express();

app.use(cors());
app.use(express.json());

app.use(`/stats`, statsRouter);

app.use(globalErrorHandler);

export default app;
