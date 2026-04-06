import express, { type Application, type Request, type Response } from "express";
import dotenv from "dotenv";
import cors from "cors";
import helmet from "helmet";
import cookieParser from "cookie-parser";
import morgan from "morgan";

// load environment variables from .env file
dotenv.config();

// initialize Express application
const app: Application = express();
const PORT = process.env.PORT || 4000;

// Middleware
app.use(
  cors({
    origin: process.env.FRONTEND_URL || "http://localhost:5173",
  }),
);

// Configure Helment to allow cross-origin resource sharing
app.use(
  helmet({
    crossOriginResourcePolicy: { policy: "cross-origin" },
  }),
);

// Use cookie parser middleware to parse cookies in incoming requests
app.use(cookieParser());

// Body parsers
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Logging middleware (only in development mode)
if (process.env.NODE_ENV == "development") {
  app.use(morgan("dev"));
}

// Basic route for testing
app.get("/", (_req: Request, res: Response) => {
  res.send("Hello from Hospital Management System Backend");
});

// Global Error Handler
app.use((err: any, _req: Request, res: Response, _next: any) => {
  const statusCode = res.statusCode === 200 ? 500 : res.statusCode;
  res.status(statusCode);
  res.json({
    message: err.message,
    stack: process.env.NODE_ENV === "production" ? null : err.stack,
  });
});

// start the server
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
