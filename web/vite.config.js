import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "alpha-uhmi",
      manifest: true,
      remotes: {
        esm_remote: {
          type: "module",
          name: "esm_remote",
          entry: "https://[...]/remoteEntry.js",
        },
        var_remote: "var_remote@https://[...]/remoteEntry.js",
      },
      exposes: {
        "./button": "./src/components/button",
      },
      shared: {
        react: {
          singleton: true,
        },
        "react/": {
          singleton: true,
        },
      },
    }),
  ],
});
