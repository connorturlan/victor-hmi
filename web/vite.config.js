import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { federation } from "@module-federation/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "alpha-uhmi",
      remotes: {
        // module federation is being handled dynamically. see app.tsx
        // add dummy.js to prevent vite from throwing an error
        dummy: "dummy.js",
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
  build: {
    target: "esnext",
  },
});
