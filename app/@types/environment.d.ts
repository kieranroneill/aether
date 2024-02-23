declare namespace NodeJS {
  export interface ProcessEnv {
    readonly API_URL: string;
    readonly NEXT_PUBLIC_LOG_LEVEL: string;
    readonly NEXT_PUBLIC_RAPYD_SCRIPT_SRC: string;
  }
}
