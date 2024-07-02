export interface ArchiveEntry {
  name: string;
  type: string;
  lastModified: number;
  offset: number;
  size: number;
}

export interface Plugin {
  (hot: Omit<Hot, "fire" | "listen">): void;
  displayName?: string;
}

export interface FireOptions {
  main?: string;
  buildTarget?: `es20${15 | 16 | 17 | 18 | 19 | 20 | 21 | 22}`;
  swModule?: string;
  swScope?: string;
  swUpdateViaCache?: ServiceWorkerUpdateViaCache;
}

export interface Hot {
  use(...plugins: readonly Plugin[]): this;
  onFetch(handler: (event: FetchEvent) => void): this;
  onFire(handler: (sw: ServiceWorker) => void): this;
  onUpdateFound: () => void;
  waitUntil(...promises: readonly Promise<any>[]): this;
  fire(options?: FireOptions): Promise<ServiceWorker>;
  listen(): void;
}

export const hot: Hot;
export default hot;
