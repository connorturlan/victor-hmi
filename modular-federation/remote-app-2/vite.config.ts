import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import federation from "@originjs/vite-plugin-federation";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "remote-app",
      filename: "remoteEntry.js",
      exposes: {
        "./Asset": "./src/components/Asset",
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
    target: "esnext",
  },
});
