import { JSX } from "solid-js";

export default function ConfusedFace(props: JSX.IntrinsicElements['svg']) {
	return (
		<svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 36 36" {...props}>
			<path fill="#FFCC4D" d="M36 18c0 9.941-8.059 18-18 18c-9.94 0-18-8.059-18-18C0 8.06 8.06 0 18 0c9.941 0 18 8.06 18 18">
			</path>
			<ellipse cx="11.5" cy="16.5" fill="#664500" rx="2.5" ry="3.5">
			</ellipse>
			<ellipse cx="24.5" cy="16.5" fill="#664500" rx="2.5" ry="3.5">
			</ellipse>
			<path fill="#664500" d="M12 28c2-5 13-5 13-3c0 1-8-1-13 3">
			</path>
		</svg>
	)
}
