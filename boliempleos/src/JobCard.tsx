import { Calendar, Location } from "./icons";
import { JobItem } from "./utils";

export default function JobCard(
	props: { job: JobItem, onSelect: (j: JobItem) => void },
) {
	return (
		<div
			class="card card-compact flex-row items-center bg-base-200 hover:bg-base-300 cursor-pointer shadow shadow-gray-800"
			onClick={() => props.onSelect(props.job)}
		>
			<div class="pl-3 w-20 min-w-0 md:w-28">
				<img src={props.job.companyLogo} alt={props.job.company} class="w-full" />
			</div>
			<div class="card-body flex-1">
				<h2 class="card-title leading-none">{props.job.title}</h2>
				<p class="flex items-center" classList={{ italic: !props.job.depto }}>
					<Location class="text-lg mr-1" />
					{props.job.depto || "No especificado"}
				</p>
				<p class="flex items-center" classList={{ italic: !props.job.publishDate }}>
					<Calendar class="text-lg mr-1" />
					{props.job.publishDate ? 'Publicado el ' + new Date(props.job.publishDate).toLocaleDateString() : "No especificado"}
				</p>
			</div>
		</div>
	);
}
