import type { Subroute } from "./subroute.type";

export interface RouteType {
  label: string;
  icon: string;
  subroutes: Subroute[];
  adminOnly?: boolean
}
