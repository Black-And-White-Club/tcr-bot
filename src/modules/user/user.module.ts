// src/modules/user/user.module.ts
import { Module } from "@nestjs/common";
import { UserResolver } from "./user.resolver";
import { UserService } from "./user.service";
import { DatabaseModule } from "src/db/database.module";
import * as schema from "src/schema";

@Module({
  imports: [DatabaseModule.forFeature(schema, "USER_DATABASE_CONNECTION")],
  providers: [UserResolver, UserService],
  exports: [UserService],
})
export class UserModule {}
