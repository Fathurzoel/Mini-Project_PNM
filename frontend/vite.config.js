// vite.config.js
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
 resolve: {
  alias: {
    '@components': path.resolve(__dirname, 'src/components'),
    // '@pages': path.resolve(import.meta.url, 'src/pages'),
    // Pastikan tidak ada tanda koma ekstra di sini
  },
},

});
// eslint-disable-next-line no-undef