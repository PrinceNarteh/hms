import { ConfigifyModule } from "@itgorillaz/configify";
import { Module } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { MongoConfig } from "./config/mongo.config";

@Module({
  imports: [
    ConfigifyModule.forRootAsync({
      expandConfig: true,
    }),
    MongooseModule.forRootAsync({
      imports: [ConfigifyModule],
      inject: [MongoConfig],
      useFactory: (config: MongoConfig) => ({
        uri: config.dbUri,
      }),
    }),
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
