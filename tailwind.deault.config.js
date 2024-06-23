const config = {
	content: ['./theme/default/**/*.html'],
	darkMode: "class",
	theme: {
		extend: {
			typography: () => ({
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
    ]
}
export default config;