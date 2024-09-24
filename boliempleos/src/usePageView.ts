import { getCurrentBrowserFingerPrint } from "@rajesh896/broprint.js";
import { onMount, createSignal } from "solid-js";

export function usePageView(visitedUrl: string) {
	const [browserId, setBrowserId] = createSignal<string | null>(null);

	async function registerPageView() {
		if (!browserId()) {
			setBrowserId(await getCurrentBrowserFingerPrint());
		}
		const url = import.meta.env.VITE_API_COMMON_URL + "/page_view";
		await fetch(url, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				userId: browserId()?.toString(),
				app: "boliempleos",
				url: visitedUrl,
			}),
		});
	}

	onMount(() => {
		registerPageView();
	});
}
