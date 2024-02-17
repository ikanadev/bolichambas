import { For } from "solid-js";
import { useAppState } from "./useAppState";
import { Depto, ALL_DEPTOS, DeptoOption } from "./utils";

export default function Filters() {
	const { depto, setDepto } = useAppState();
	return (
		<div class="navbar justify-end pb-0 rounded-md">
			<label class="form-control items-start">
				<div class="label pb-1">
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
		</div>
	)
}
