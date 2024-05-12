/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./theme/glass/**/*.html'],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        'eerie': {
          DEFAULT: '#1c2522',
          100: '#060807',
          200: '#0b0f0e',
          300: '#111715',
          400: '#171e1c',
          500: '#1c2522',
          600: '#435851',
          700: '#698b80',
          800: '#9ab3aa',
          900: '#ccd9d5'
        },
        'hooker': {
          DEFAULT: '#60746a',
          100: '#131715',
          200: '#272f2b',
          300: '#3a4640',
          400: '#4e5d56',
          500: '#60746a',
          600: '#7e9389',
          700: '#9eaea6',
          800: '#bfc9c4',
          900: '#dfe4e1'
        },
        'cambridge': {
          DEFAULT: '#a4c3b2',
          100: '#1c2b23',
          200: '#385646',
          300: '#558269',
          400: '#77a68c',
          500: '#a4c3b2',
          600: '#b5cfc1',
          700: '#c7dbd0',
          800: '#dae7e0',
          900: '#ecf3ef'
        },
        'mint': {
          DEFAULT: '#cce3de',
          100: '#1f3832',
          200: '#3e7065',
          300: '#5ea697',
          400: '#96c5bb',
          500: '#cce3de',
          600: '#d8e9e5',
          700: '#e1efec',
          800: '#ebf4f2',
          900: '#f5faf9'
        },
        'azure': {
          DEFAULT: '#d8eae6',
          100: '#1f3a35',
          200: '#3f7569',
          300: '#63ab9b',
          400: '#9dcac0',
          500: '#d8eae6',
          600: '#dfeeeb',
          700: '#e7f2f0',
          800: '#eff6f5',
          900: '#f7fbfa'
        }
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            ul: {
              listStyle: 'disc',
            },
          },
        },
        invert: {
          css: {
            ul: {
              listStyle: 'revert',
            },
          },
        },
      }),
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
