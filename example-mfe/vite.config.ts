import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import federation from "@originjs/vite-plugin-federation";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "example-mfe",
      filename: "remoteEntry.js",
      exposes: {
        "./example-mfe": "./src/Widget",
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
  preview: {
    cors: true,
  },
});
