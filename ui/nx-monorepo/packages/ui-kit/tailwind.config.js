/* eslint-disable */

const { fontFamily } = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} \*/
module.exports = {
  mode: 'jit',
  darkMode: 'class',
  content: {
    files: [
      './src/**/*.{js,ts,jsx,tsx}',
    ],
    options: {
      safelist: ['dark'], // specific classes
    },
  },
  theme: {
    fontFamily: {
      display: ['Roboto Mono', 'Menlo', 'monospace'],
      body: ['Roboto Mono', 'Menlo', 'monospace'],
    },
    extend: {
      typography: (theme) => ({
        dark: {
          css: {
            color: 'white',
          },
        },
        fontFamily: {
          sans: ['var(--font-inter)', ...fontFamily.sans],
        },
      }),
    },
  },
  variants: {
    typography: ['dark'],
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
  ],
}
