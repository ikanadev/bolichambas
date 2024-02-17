import { For, createMemo, createSignal } from "solid-js";
import { ALL_DEPTOS, Company, JobItem } from "./utils";
import JobCard from "./JobCard";
import JobDetails from "./JobDetails";
import { useAppState } from "./useAppState";

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
			<div class="flex flex-col gap-3 py-4">
				<For each={jobs()}>
					{(job) => <JobCard job={job} onSelect={setSelected} />}
				</For>
			</div>
			<JobDetails job={selected()} onClose={closeModal} />
		</>
	);
}
