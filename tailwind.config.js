/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./theme/**/*.html'],
  darkMode: "class",
  theme: {
    extend: {
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
