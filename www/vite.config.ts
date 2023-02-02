import { defineConfig } from "vite";
import relay from "vite-plugin-relay-lite";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      "/graphql": "http://localhost:8080",
    },
  },
  plugins: [
    react(),
    relay({
      codegen: false,
    }),
  ],
});
