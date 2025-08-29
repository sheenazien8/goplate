import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
  plugins: [vue()],
  publicDir: false, // Disable public directory since we're outputting to public
  build: {
    outDir: 'public',
    rollupOptions: {
      input: path.resolve(__dirname, "templates/js/app.js")
    },
    manifest: true,
    assetsDir: 'assets'
  },
  server: {
    host: '0.0.0.0',
    port: 5173,
    hmr: {
      host: 'localhost'
    },
    cors: true
  }
})
