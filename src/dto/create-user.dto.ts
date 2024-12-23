// src/dto/create-user.dto.ts
import {
  IsString,
  IsInt,
  IsOptional,
  IsEnum,
  IsNotEmpty,
} from "class-validator";
import { UserRole } from "../enums/user-role.enum"; // Adjust the path according to your structure

export class CreateUserDto {
  @IsNotEmpty({ message: "Name should not be empty" })
  @IsString()
  name!: string;

  @IsNotEmpty({ message: "DiscordID should not be empty" })
  @IsString()
  discordID!: string;

  @IsInt()
  @IsOptional()
  tagNumber: number | undefined;

  @IsEnum(UserRole)
  role!: UserRole;
}
