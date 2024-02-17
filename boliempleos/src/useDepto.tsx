import { Accessor, JSX, createContext, createSignal, useContext } from "solid-js";
import { DeptoOption, ALL_DEPTOS } from "./utils";

const DeptoContext = createContext<{
	depto: Accessor<DeptoOption>;
	setDepto: (d: DeptoOption) => void;
}>();

export function DeptoProvider(props: { children: JSX.Element }) {
	const [depto, setDepto] = createSignal<DeptoOption>(ALL_DEPTOS);

	return (
		<DeptoContext.Provider value={{ depto: depto, setDepto }}>
			{props.children}
		</DeptoContext.Provider>
	);
}

export function useDepto() {
	return useContext(DeptoContext)!;
};
