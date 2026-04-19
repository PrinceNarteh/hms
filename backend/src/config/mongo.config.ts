import { Configuration, Value } from "@itgorillaz/configify";
import { IsString } from "class-validator";

@Configuration()
export class MongoConfig {
  @IsString()
  @Value("DB_NAME")
  dbName: string;

  @IsString()
  @Value("DB_USER")
  dbUser: string;

  @IsString()
  @Value("DB_PASSWORD")
  dbPassword: string;

  @IsString()
  @Value("DB_URI")
  dbUri: string;
}
