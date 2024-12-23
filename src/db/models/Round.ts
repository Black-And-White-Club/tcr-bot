import {
  pgTable,
  serial,
  varchar,
  json,
  boolean,
  date,
  time,
} from "drizzle-orm/pg-core";
import { timestamps } from "./timetamps.helpers";

export const RoundModel = pgTable("rounds", {
  roundID: serial().primaryKey(),
  title: varchar().notNull(),
  location: varchar().notNull(),
  eventType: varchar(),
  date: date().notNull(),
  time: time().notNull(),
  participants: json().notNull(),
  scores: json().notNull(),
  finalized: boolean().default(false),
  creatorID: varchar().notNull(),
  state: varchar().notNull(),
  ...timestamps,
});
