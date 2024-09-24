/// <reference types="vite/client" />

interface ImportMetaEnv {
	readonly VITE_API_VIEWS_URL: string;
}

interface ImportMeta {
	readonly env: ImportMetaEnv;
}

