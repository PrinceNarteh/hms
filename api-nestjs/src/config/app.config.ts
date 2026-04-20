import { Value } from "@itgorillaz/configify";
import { IsIn, IsPort } from "class-validator";

export class AppConfig {
  @IsPort()
  @Value("APP_PORT", { parse: parseInt })
  port: number;

  @IsIn(["development", "production"])
  @Value("NODE_ENV")
  env: string;
}
