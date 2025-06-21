import type { ResponseType } from "../types";
import { config } from "../config";

const api_url = config.api_url;

export class api {
  public static async get<T>(
    path: string,
    access_token?: string
  ): Promise<ResponseType<T>> {
    const res = await fetch(`${api_url}/${path}`, {
      method: "GET",
      headers: {
        authorization: access_token || "",
      },
    });

    return {
      data: res.ok ? await res.json() : null,
      error: res.ok ? false : true,
    };
  }

  public static async post<T>(
    path: string,
    data: T,
    access_token?: string
  ): Promise<ResponseType<T>> {
    const res = await fetch(`${api_url}/${path}`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        authorization: access_token || "",
      },
    });

    return {
      data: res.ok ? await res.json() : null,
      error: res.ok ? false : true,
    };
  }

  public static async patch<T>(
    path: string,
    data: T,
    access_token?: string
  ): Promise<ResponseType<T>> {
    const res = await fetch(`${api_url}/${path}`, {
      method: "PATCH",
      body: JSON.stringify(data),
      headers: {
        authorization: access_token || "",
      },
    });

    return {
      data: res.ok ? await res.json() : null,
      error: res.ok ? false : true,
    };
  }

  public static async delete<T>(
    path: string,
    access_token?: string
  ): Promise<ResponseType<T>> {
    const res = await fetch(`${api_url}/${path}`, {
      method: "DELETE",
      headers: {
        authorization: access_token || "",
      },
    });

    return {
      data: res.ok ? await res.json() : null,
      error: res.ok ? false : true,
    };
  }
}
