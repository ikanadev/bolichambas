/** @type {import('tailwindcss').Config} */
export default {
	content: [
		"./index.html",
		"./src/**/*.{js,ts,jsx,tsx}",
	],
	theme: {
		extend: {},
	},
	plugins: [require("daisyui")],
	daisyui: {
		themes: [
			{
				dark: {
					...require("daisyui/src/theming/themes")["sunset"],
					primary: "#2F4A6F",
					secondary: "#929a8F",
					accent: "#3D5A80",
					"--rounded-box": "0.5rem",
					"--rounded-btn": "0.25rem",
				},
			},
		],
	}
}

