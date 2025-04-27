import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import federation from "@originjs/vite-plugin-federation";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "host-app",
      filename: "remoteEntry.js",
      remotes: {
        remote_app_1: "http://localhost:4173/assets/remoteEntry.js",
      },
      shared: {
        react: {
          requiredVersion: "^18.0.0",
        },
        "react-dom": {
          requiredVersion: "^18.0.0",
        },
      },
    }),
  ],
  build: {
    minify: false,
    target: "esnext",
  },
});
