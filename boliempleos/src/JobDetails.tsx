import { Show, JSX } from "solid-js";
import { JobItem } from "./utils";
import { Calendar, DegreeHat, Location, OpenInNew } from "./icons";

function ItemInfo(props: { label: string, value: string, icon: JSX.Element }) {
	return (
		<div class="flex items-center gap-2">
			{props.icon}
			<div>
				<p class="text-sm leading-tight">{props.label}</p>
				<p class="leading-tight" classList={{ italic: !props.value }}>
					{props.value || "Sin especificar"}
				</p>
			</div>
		</div>
	);
}

export default function JobDetails(
	props: { job: JobItem | null, onClose: () => void },
) {
	return (
		<dialog open={!!props.job} class="modal modal-bottom sm:modal-middle">
			<div class="modal-box max-h-full w-full bg-base-300 shadow-xl mx-auto shadow-gray-800 relative" style="max-width: 48rem !important;">
				<div class="sticky top-0 flex justify-end">
					<button
						class="btn btn-sm btn-circle btn-ghost text-xl"
						onClick={props.onClose}
					>
						✕
					</button>
				</div>
				<Show when={props.job}>
					{(job) => (
						<div class="flex flex-col items-stretch">
							<div>
								<img src={job().companyLogo} alt={job().company} class="w-40 mx-auto" />
							</div>
							<h1 class="card-title text-center self-center mt-4">{job().title}</h1>
							<h2 class="text-lg text-center self-center">{job().company}</h2>
							<div class="divider" />
							<div class="grid grid-cols-1 sm:grid-cols-2 gap-x-2 gap-y-2 sm:gap-y-3">
								<ItemInfo
									label="Departamento"
									value={job().depto}
									icon={<Location class="text-2xl" />}
								/>
								<ItemInfo
									label="Fecha de publicación"
									value={!!job().publishDate ? new Date(job().publishDate).toLocaleDateString() : ""}
									icon={<Calendar class="text-2xl" />}
								/>
								<ItemInfo
									label="Área"
									value={job().area}
									icon={<DegreeHat class="text-2xl" />}
								/>
								<ItemInfo
									label="Límite de postulación"
									value={job().dueDate}
									icon={<Calendar class="text-2xl" />}
								/>
							</div>
							<div class="divider" />
							<div class="text-base" innerHTML={job().content}></div>
							<a
								href={job().url}
								class="text-lg btn btn-primary mt-4 sticky bottom-0"
								target="_blank"
								rel="noopener noreferrer"
							>
								Ver oferta
								<OpenInNew />
							</a>
						</div>
					)}
				</Show>
			</div>
		</dialog>
	);
}
