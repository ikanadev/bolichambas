import { Match, Show, Switch, createResource } from "solid-js";
import Filters from "./Filters"
import Header from "./Header"
import getData from "./data";
import Items from "./Items";

function App() {
	const [jobsRes] = createResource(getData);

	return (
		<>
			<Header />
			<Filters />
			<Show when={jobsRes.loading}>
				<div class="flex flex-col items-center justify-center py-16">
					<span class="loading loading-ring loading-lg"></span>
					<p class="text-sm text-secondary">Cargando...</p>
				</div>
			</Show>
			<Switch>
				<Match when={jobsRes.error}>
					<div role="alert" class="alert alert-error my-4">
						<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
							/>
						</svg>
						<span>Error cargando datos: {jobsRes.error}</span>
					</div>
				</Match>
				<Match when={jobsRes()}>
					{(companies) => <Items companies={companies()} />}
				</Match>
			</Switch>
		</>
	)
}

export default App
