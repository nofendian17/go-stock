/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'candlestick-green': '#26a69a',
        'candlestick-red': '#ef5350',
      },
    },
  },
  plugins: [],
}
