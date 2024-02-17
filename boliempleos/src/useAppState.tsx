import { Accessor, JSX, createContext, createSignal, useContext } from "solid-js";
import { DeptoOption, ALL_DEPTOS } from "./utils";

const AppStateContext = createContext<{
	depto: Accessor<DeptoOption>;
	setDepto: (d: DeptoOption) => void;
}>();

export function AppStateProvider(props: { children: JSX.Element }) {
	const [depto, setDepto] = createSignal<DeptoOption>(ALL_DEPTOS);

	const appState = {
		depto,
		setDepto
	};

	return (
		<AppStateContext.Provider value={appState}>
			{props.children}
		</AppStateContext.Provider>
	);
}

export function useAppState() {
	return useContext(AppStateContext)!;
};
