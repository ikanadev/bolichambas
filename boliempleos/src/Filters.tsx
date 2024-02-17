import { For } from "solid-js";
import { useAppState } from "./useAppState";
import { Depto, ALL_DEPTOS, DeptoOption } from "./utils";

export default function Filters() {
	const { depto, setDepto } = useAppState();
	return (
		<div class="navbar justify-between my-2 rounded-md">
			<label class="form-control items-start">
				<div class="label pb-0">
					<span class="label-text">Departamento</span>
				</div>
				<select
					class="select select-bordered select-sm"
					value={depto()}
					onChange={(e) => setDepto(e.currentTarget.value as DeptoOption)}
				>
					<option value={ALL_DEPTOS}>Todos</option>
					<For each={Object.values(Depto)}>
						{(d) => <option value={d}>{d}</option>}
					</For>
				</select>
			</label>
			<label class="input input-bordered flex items-center gap-2">
				<input type="text" class="grow" placeholder="Search" />
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 16 16"
					fill="currentColor"
					class="w-4 h-4 opacity-70"
				>
					<path
						fill-rule="evenodd"
						d="M9.965 11.026a5 5 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06l-2.755-2.754ZM10.5 7a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0Z"
						clip-rule="evenodd"
					/>
				</svg>
			</label>
		</div>
	)
}
