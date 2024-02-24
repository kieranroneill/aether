declare namespace NodeJS {
  export interface ProcessEnv {
    // private
    readonly ENVIRONMENT: string;
    readonly NAME: string;
    readonly PORT: string;

    // public
    readonly NEXT_PUBLIC_DESCRIPTION: string;
    readonly NEXT_PUBLIC_LOG_LEVEL: string;
    readonly NEXT_PUBLIC_TAGLINE: string;
    readonly NEXT_PUBLIC_TITLE: string;
    readonly NEXT_PUBLIC_VERSION: string;
  }
}
