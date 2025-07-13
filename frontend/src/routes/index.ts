import type { RouteType } from "../types";

export const routes: RouteType[] = [
  {label: "Pacientes", icon: "patient", 
    subroutes: [
      {label: "Listagem", path: "/patients"},
      {label: "Novo", path: "/patients/new"}
    ]
  },
  {label: "Calendário", icon: "calendar", 
    subroutes: [
      {label: "Visualizar", path: "/calendar"},
      {label: "Nova Consulta", path: "/calendar/new"}
    ]
  },
  {label: "Usuários", icon: "user", adminOnly: true, 
    subroutes: [
      {label: "Visualizar", path: "/users"},
      {label: "Novo usuário", path: "/users/new"}
    ]
  }
];
