import { For, Show, createMemo, createSignal } from "solid-js";
import { ALL_DEPTOS, Company, JobItem, MIN_CHARS_FOR_SEARCH } from "./utils";
import JobCard from "./JobCard";
import JobDetails from "./JobDetails";
import { useAppState } from "./useAppState";
import ConfusedFace from "./icons/ConfusedFace";

export default function Items(props: { companies: Company[] }) {
	const [selected, setSelected] = createSignal<JobItem | null>(null);

	const { depto } = useAppState();

	function closeModal() {
		setSelected(null);
	}

	const jobs = createMemo(() => {
		const items: Array<JobItem> = [];
		props.companies.forEach((c) => {
			c.jobs.forEach((j) => {
				if (j.depto === depto() || depto() === ALL_DEPTOS) {
					items.push({
						company: c.name,
						companyLogo: c.logoUrl,
						...j,
					});
				}
			});
		});
		items.sort((a, b) => {
			return new Date(b.publishDate).getTime() - new Date(a.publishDate).getTime();
		});
		return items;
	});

	return (
		<>
			<div class="flex flex-col gap-3 py-0 sm:py-2">
				<p class="text-sm">
					Mostrando <strong>{jobs().length}</strong> {jobs().length === 1 ? "trabajo " : "trabajos "}
					<Show when={depto() !== ALL_DEPTOS}>
						en <strong>{depto()} </strong>
					</Show>
				</p>
				<For each={jobs()}>
					{(job) => <JobCard job={job} onSelect={setSelected} />}
				</For>
				<Show when={jobs().length === 0}>
					<div class="flex flex-col justify-center items-center py-6">
						<ConfusedFace width={60} height={60} />
						<p class="italic text-secondary">No hay resultados</p>
					</div>
				</Show>
			</div>
			<JobDetails job={selected()} onClose={closeModal} />
		</>
	);
}
